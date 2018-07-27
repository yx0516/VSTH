package xorm

import (
	"errors"
	"reflect"
	"strings"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"gpc/util/ref"
	"gpc/util/tag"

	"gpc/db/mongodb/conn/gsession"
	"gpc/db/mongodb/orm/structure"
	"gpc/db/mongodb/utils"

	mdl "gpc/db/mongodb/xorm/model"
	plib "gpc/publib"
)

//-----------------------------------------------------------------------------------------------------------//

// 结构体 公用 类
type ModelClass struct {
	object                 mdl.Model                                    `json:"-"`
	FuncCheckInsert        func() error                                 `json:"-"`
	FuncMakeMatch          func() (bson.M, error)                       `json:"-"`
	FuncUpdateVerifyFields func(fields plib.StringArray) (MapSI, error) `json:"-"`
	FuncUpdateByFields     func(fvs MapSI, vSelf interface{}) error     `json:"-"`
}

func NewModelClass(object mdl.Model) *ModelClass {
	return &ModelClass{
		object: object,
	}
}

// 注册 子类
func (self *ModelClass) init(object mdl.Model) {
	self.object = object
}

//-----------------------------------------------------------------------------------------------------------//

// 检查 子类是否为 nil
func (self *ModelClass) CheckObject() error {
	if self.object == nil {
		return errors.New("object is nil.")
	} else {
		return nil
	}
}

// 复制 子类
func (self *ModelClass) Copy() (mdl.Model, error) {
	if err := self.CheckObject(); err != nil {
		return nil, err
	}
	return self.object.Copy()
}

// 获取 集合 名称
func (self *ModelClass) GetCollName() (string, error) {
	return GetCollName(self.object)
}

// 忽略错误
func (self *ModelClass) CollName() (name string) {
	name, _ = self.GetCollName()
	return
}

// 生成一个 子类指针 模型 实例
func (self *ModelClass) MakeModelPtr() (object mdl.Model, err error) {
	if subType, err := self.GetStructType(); err == nil {
		object = reflect.New(subType).Interface().(mdl.Model)
		object.Init()
		return object, nil
	} else {
		return nil, err
	}
}

// 生成一个 子类Slice指针 实例 *[]Model
func (self *ModelClass) MakeStructSlicePtr() (interface{}, error) {
	if subType, err := self.GetStructType(); err == nil {
		return reflect.New(reflect.SliceOf(subType)).Interface(), nil
	} else {
		return nil, err
	}
}

// 获取 子类类型 Model 【非指针】
func (self *ModelClass) GetStructType() (subType reflect.Type, err error) {
	if err = self.CheckObject(); err == nil {
		return reflect.Indirect(reflect.ValueOf(self.object)).Type(), nil
	}
	return
}

// 获取 子类 全名【不区分指针】
func (self *ModelClass) GetStructFullName() (string, error) {
	return ref.RefStructTypeFullName(self.object)
}

//-----------------------------------------------------------------------------------------------------------//

// 获取 结构体 字段
func (self *ModelClass) GetStructFields() (*tag.StructField, error) {
	return globalStructField.GetCopy(self.object)
}

// 获取 参数 sts 里的第一个非 nil 值 或 重新生成
func (self *ModelClass) GetStructTag(sts ...*tag.StructTag) (*tag.StructTag, error) {
	if len(sts) > 0 && sts[0] != nil {
		return sts[0], nil
	}
	if err := self.CheckObject(); err != nil {
		return nil, err
	}
	return globalStructCache.Add(self.object)
}

// 字符串字段删除头尾空格
func (self *ModelClass) StringTrim(sts ...*tag.StructTag) (*tag.StructTag, error) {
	if st, err := self.GetStructTag(sts...); err != nil {
		return nil, err
	} else {
		return st, st.StringTrim()
	}
}

// 设置空默认值
func (self *ModelClass) SetNullDefValue(sts ...*tag.StructTag) (*tag.StructTag, error) {
	if st, err := self.GetStructTag(sts...); err != nil {
		return nil, err
	} else {
		return st, st.SetNullDefValue()
	}
}

// 获取非空内容的字段
func (self *ModelClass) GetNotNullFields(sts ...*tag.StructTag) (plib.StringArray, error) {
	if st, err := self.GetStructTag(sts...); err != nil {
		return nil, err
	} else {
		return st.NotNullFields()
	}
}

// 获取 可更新的非空字段
func (self *ModelClass) GetNotNullUpdateFields() (fields plib.StringArray, err error) {
	var st *tag.StructTag
	if st, err = self.StringTrim(); err != nil {
		return
	}
	var tagSf *tag.StructField
	if tagSf, err = self.GetStructFields(); err != nil {
		return
	}
	for field, sf := range st.Fields {
		if !tagSf.White.Contains(field) {
			continue
		}
		realTag := sf.FilterTags(st.Fc.TagNotNull)
		if realTag.Count() <= 0 {
			continue // 不执行无注解的字段
		}
		if ret := st.Fc.IsNotNull(sf, realTag...); ret.IsOK() {
			fields.Add(field)
		} else if ret.IsErr() {
			err = errors.New(ret.Msg)
			return
		}
	}
	return
}

// 检查字段内容是否合法
func (self *ModelClass) VerifyFields(fields plib.StringArray, sts ...*tag.StructTag) (MapSI, error) {
	if st, err := self.GetStructTag(sts...); err != nil {
		return nil, err
	} else {
		if len(fields) == 0 {
			if fields, err = st.NotNullFields(); err != nil {
				return nil, err
			}
		}
		return st.Check(fields...)
	}
}

// 字符串 删除 头尾空格 并且设置 默认值
func (self *ModelClass) TrimSetDef(fields ...string) error {
	if st, err := self.StringTrim(); err == nil {
		return st.SetNullDefValue(fields...)
	} else {
		return err
	}
}

// 字符串 删除 头尾空格 并且设置 默认值 | 并且获取 非空字段
func (self *ModelClass) TrimSetDefGetNotNull(fields ...string) (plib.StringArray, error) {
	if st, err := self.StringTrim(); err == nil {
		if err = st.SetNullDefValue(fields...); err != nil {
			return nil, err
		} else {
			return st.NotNullFields()
		}
	} else {
		return nil, err
	}
}

// 字符串 删除 头尾空格 | 设置 默认值 | 检验数据
func (self *ModelClass) TrimSetDefVerify(fields ...string) (map[string]interface{}, error) {
	if st, err := self.StringTrim(); err == nil {
		if err = st.SetNullDefValue(fields...); err == nil {
			return st.Check(fields...)
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}

// 所有字段
func (self *ModelClass) TrimAllSetDefVerify() (map[string]interface{}, error) {
	if sf, err := self.GetStructFields(); err != nil {
		return nil, err
	} else {
		return self.TrimSetDefVerify(sf.All...)
	}
}

//-----------------------------------------------------------------------------------------------------------//

// 检查更新白名单字段
func (self *ModelClass) CheckUpdateWhiteFields(fields plib.StringArray) error {
	if sf, err := self.GetStructFields(); err != nil {
		return err
	} else {
		for _, field := range fields {
			if !sf.White.Contains(field) {
				return errors.New(field + " field can not be updated.")
			}
		}
		return nil
	}
}

// 删除更新黑名单里的字段
func (self *ModelClass) CheckUpdateBlackFields(fields *plib.StringArray) error {
	if fields == nil {
		return nil
	}
	if sf, err := self.GetStructFields(); err != nil {
		return err
	} else {
		if len(sf.Black) > 0 { // 是否有黑名单
			for _, field := range sf.Black {
				if fields.Contains(field) {
					fields.Delete(field, true)
				}
			}
		}
		return nil
	}
}

// 检查 过滤 更新 字段
func (self *ModelClass) CheckUpdateFliterFields(filterOpt uint8, fields *plib.StringArray) (err error) {
	if fields == nil {
		return errors.New("fields is nil.")
	} else {
		switch filterOpt {
		case CST_FILTER_OPT_WHITE: // 白名单过滤,有 error 返回
			return self.CheckUpdateWhiteFields(*fields)
		case CST_FILTER_OPT_BLACK: // 黑名单过滤,无 error 返回
			return self.CheckUpdateBlackFields(fields)
		}
	}
	return
}

// 返回不同值的可更新的字段名
func (self *ModelClass) DiffValueUpdateFields(object interface{}, filterOpt uint8, fields ...string) (plib.StringArray, error) {
	var st1, st2 *tag.StructTag
	var err error
	var diffFields plib.StringArray

	if st1, err = globalStructCache.Add(object); err != nil {
		return nil, err
	}
	if st2, err = globalStructCache.Add(self.object); err != nil {
		return nil, err
	}

	if diffFields, err = tag.DiffValueFields(st1, st2, fields...); err != nil {
		return nil, err
	}
	diffFields.RemoveDuplicate()
	err = self.CheckUpdateFliterFields(filterOpt, &diffFields)
	return diffFields, err
}

//-----------------------------------------------------------------------------------------------------------//

// 获取 字段 键值对
func (self *ModelClass) GetFieldValues(fields ...string) (map[string]interface{}, error) {
	if st, err := globalStructCache.Add(self.object); err != nil {
		return nil, err
	} else if err = st.StringTrim(fields...); err != nil {
		return nil, err
	}

	if info, err := structure.Ref(self.object); err != nil {
		return nil, err
	} else {
		return info.FilterFields(true, fields...)
	}
}

// 获取 Id 字段值
func (self *ModelClass) GetFieldIdValue(object interface{}) (id int, err error) {
	var vId interface{}
	if vId, err = ref.GetStructFieldValueIntf(object, "Id"); err == nil {
		var ok bool
		if id, ok = vId.(int); !ok {
			err = errors.New("id field interface convert fail.")
		} else {
			err = self.CheckId(id)
		}
	}
	return
}

// 检查 Id 是否有效
func (self *ModelClass) CheckId(id int) error {
	return obj.chk.Id(id)
}

// 检查 ID 是否有效
func (self *ModelClass) VerifyId() (id int, err error) {
	if id, err = self.GetFieldIdValue(self.object); err == nil {
		err = obj.chk.Id(id)
	}
	return
}

// 获取 Name 字段值
func (self *ModelClass) GetFieldNameValue(object interface{}) (name string, err error) {
	var val interface{}
	if val, err = ref.GetStructFieldValueIntf(object, "Name"); err == nil {
		var ok bool
		if name, ok = val.(string); !ok {
			err = errors.New("name field interface convert fail.")
			return
		}
	}
	return
}

// 设置 字段 Id 的值
func (self *ModelClass) SetFieldIdValue(object mdl.Model, id int) error {
	err := ref.SetFieldValue(object, "Id", func(val reflect.Value) { val.SetInt(int64(id)) })
	if err == nil {
		object.Init()
	}
	return err
}

// 设置 字段 Name 的值
func (self *ModelClass) SetFieldNameValue(object mdl.Model, name string) error {
	err := ref.SetFieldValue(object, "Name", func(val reflect.Value) { val.SetString(name) })
	if err == nil {
		object.Init()
	}
	return err
}

//-----------------------------------------------------------------------------------------------------------//

// 记得 c.Database.Session.Close() 关闭掉
func (self *ModelClass) NewColl() *mgo.Collection {
	return gsession.NewDB().C(self.CollName())
}

// 切换到集合
func (self *ModelClass) WithColl(fun func(c *mgo.Collection)) {
	db := gsession.NewDB()
	defer db.Session.Close()
	fun(db.C(self.CollName()))
}

// 检查 集合 查询表达式 是否存在
func (self *ModelClass) Exist(match interface{}) (ok bool, err error) {
	if match == nil {
		return false, errors.New("match is nil.")
	}
	db := gsession.NewDB()
	defer db.Session.Close()
	return utils.Exist(db.C(self.CollName()), match)
}

// 检查 集合 里的 【属性】 【值】 是否存在
func (self *ModelClass) ExistValue(field string, value interface{}) (ok bool, err error) {
	field = strings.ToLower(field) // 转成小写
	db := gsession.NewDB()
	defer db.Session.Close()
	return utils.ExistValue(db.C(self.CollName()), field, value)
}

// 是否 存在 指定字段匹配的记录
func (self *ModelClass) ExistFieldValue(fields ...string) (ok bool, err error) {
	cond := utils.NewCond(nil)
	if cond.Select, err = self.GetFieldValues(fields...); err == nil {
		self.WithColl(func(c *mgo.Collection) {
			cond.Coll = c
			ok, err = cond.QueryExist()
		})
	}
	return
}

// 可指定 检查 函数
func (self *ModelClass) ExistName(name string, chk func(name string) error) (bool, error) {
	name = strings.TrimSpace(name)
	if chk != nil {
		if err := chk(name); err != nil {
			return false, err
		}
	}

	return self.ExistValue("name", name)
}

func (self *ModelClass) ExistId() (bool, error) {
	if id, err := self.VerifyId(); err == nil {
		return self.ExistValue("id", id)
	} else {
		return false, err
	}
}

//-----------------------------------------------------------------------------------------------------------//

// 复制 一个 实例 后读取数据【fields 指定只读取哪些字段的数据返回回来】
func (self *ModelClass) CopyRead(fields ...string) (object mdl.Model, err error) {
	if object, err = self.Copy(); err == nil {
		err = self.ReadOne(object, fields...)
	}
	return
}

// 根据 makeWhere 条件 读取一条数据【fields 指定只读取哪些字段的数据返回回来】
func (self *ModelClass) ReadOne(object mdl.Model, fields ...string) (err error) {
	if object == nil {
		return errors.New("object is nil.")
	}
	var match bson.M
	if match, err = object.MakeMatch(); err != nil {
		return
	}

	count := len(fields)
	fieldM := make(bson.M)
	for i := 0; i < count; i++ {
		fieldM[strings.ToLower(fields[i])] = true
	}

	init := object.Init
	msg := self.CollName() + ".readOne"
	self.WithColl(func(c *mgo.Collection) {
		query := c.Find(match)
		if len(fieldM) > 0 {
			query = query.Select(fieldM)
		}

		if err = query.One(object); err != nil {
			err = obj.err.AssRead(err, msg)
		} else {
			init()
		}
	})
	return
}

// 根据 ID 读取数据
func (self *ModelClass) ReadById(id int, fields ...string) (err error) {
	if err = self.CheckId(id); err != nil {
		return
	}
	return self.Read(bson.M{"id": id}, fields...)
}

// 根据 名称 读取数据【指定 检查 函数】【fields 指定只读取哪些字段的数据返回回来】
func (self *ModelClass) ReadByName(name string, chk func(name string) error, fields ...string) (err error) {
	name = strings.TrimSpace(name)
	if chk != nil {
		if err = chk(name); err != nil {
			return
		}
	}
	return self.Read(bson.M{"name": name}, fields...)
}

// 根据 字段 读取数据
func (self *ModelClass) ReadByFields(fields ...string) error {
	if fieldValues, err := self.GetFieldValues(fields...); err == nil {
		return self.Read(fieldValues)
	} else {
		return err
	}
}

// 读取 一条 数据【fields 指定只读取哪些字段的数据返回回来】
func (self *ModelClass) Read(m bson.M, fields ...string) (err error) {
	// 添加 过滤 字段
	count := len(fields)
	fieldM := make(bson.M)
	for i := 0; i < count; i++ {
		fieldM[strings.ToLower(fields[i])] = true
	}

	msg := self.CollName() + ".read"
	init := self.object.Init
	self.WithColl(func(c *mgo.Collection) {
		query := c.Find(m)
		if len(fieldM) > 0 {
			query = query.Select(fieldM)
		}
		if err = query.One(self.object); err != nil {
			err = obj.err.AssRead(err, msg)
		} else {
			init()
		}
	})
	return
}

// 查询 获取 数据
func (self *ModelClass) Query(collName string, match bson.M, result interface{}, fields ...string) (err error) {
	if collName = strings.TrimSpace(collName); collName == "" {
		return errors.New("coll name is null.")
	}
	if result == nil {
		return errors.New("result is nil.")
	}
	value := reflect.ValueOf(result)
	if value.Kind() != reflect.Ptr {
		return errors.New("result must be ptr.")
	}
	isSlice := reflect.Indirect(value).Kind() == reflect.Slice

	// 添加 过滤 字段
	count := len(fields)
	fieldM := make(bson.M)
	for i := 0; i < count; i++ {
		fieldM[strings.ToLower(fields[i])] = true
	}

	db := gsession.NewDB()
	defer db.Session.Close()
	query := db.C(collName).Find(match)

	if len(fieldM) > 0 {
		query = query.Select(fieldM)
	}
	if isSlice {
		err = query.All(result)
	} else {
		err = query.One(result)
	}
	err = obj.err.AssRead(err, collName+".query")

	return
}

//-----------------------------------------------------------------------------------------------------------//

// 通过 ID 来更新 指定字段
func (self *ModelClass) UpdateFieldsById(id int, fields ...string) error {
	return self.UpdateFields(bson.M{"id": id}, fields...)
}

// 更新指定字段
func (self *ModelClass) UpdateFields(match bson.M, fields ...string) (err error) {
	if fieldValues, err := self.GetFieldValues(fields...); err == nil {
		return self.Update(match, fieldValues)
	} else {
		return err
	}
}

// 通过 ID 来更新
func (self *ModelClass) UpdateById(id int, fieldValues bson.M, tags ...string) (err error) {
	return self.Update(bson.M{"id": id}, fieldValues, tags...)
}

// 更新
func (self *ModelClass) Update(match, fieldValues bson.M, tags ...string) (err error) {
	if len(match) == 0 {
		return errors.New("match is null.")
	}
	if len(fieldValues) == 0 {
		return errors.New("fieldValues is null.")
	}
	msg := self.CollName() + ".update"
	self.WithColl(func(c *mgo.Collection) {
		if err = c.Update(match, bson.M{"$set": fieldValues}); err != nil {
			err = obj.err.AssUpdate(err, msg, tags...)
		}
	})
	return
}

// 更新 指定 字段【勾入处理函数，可选】【白名单限制】
func (self *ModelClass) UpdateByFields(fields plib.StringArray, handle HookUpdate, tags ...string) (err error) {
	if len(fields) == 0 {
		return nil // 取消空字段更新时返回错误
	}
	fields.RemoveDuplicate()
	if err = self.CheckUpdateWhiteFields(fields); err != nil {
		return
	}
	var st *tag.StructTag
	if st, err = self.StringTrim(); err != nil {
		return
	}
	var object mdl.Model
	if object, err = self.CopyRead(); err != nil {
		return
	}
	if fields, err = self.DiffValueUpdateFields(object, CST_FILTER_OPT_NONE, fields...); err != nil {
		return
	}

	if len(fields) == 0 {
		return
	}

	var fieldVales MapSI
	if self.FuncUpdateVerifyFields != nil {
		if fieldVales, err = self.FuncUpdateVerifyFields(fields); err != nil {
			return
		}
	} else if fieldVales, err = self.VerifyFields(fields, st); err != nil {
		return
	}

	if sf, err := self.GetStructFields(); err == nil && sf != nil {
		var val interface{}
		var ok bool

		if sf.ZeroNil.Count() > 0 {
			var id int64
			for _, field := range sf.ZeroNil {
				if val, ok = fieldVales[field]; ok && val != nil {
					switch val.(type) {
					case time.Time:
						if val.(time.Time).IsZero() {
							fieldVales[field] = nil
						}
					default:
						if id, err = plib.ToInt64(val); err == nil && id == 0 {
							fieldVales[field] = nil
						}
					}
				}
			}
		}

		if sf.NilDel.Count() > 0 {
			for _, field := range sf.NilDel {
				if val, ok = fieldVales[field]; ok && val == nil {
					fieldVales.Delete(field)
				}
			}
		}
	}

	if handle != nil {
		if err = handle(fieldVales, object); err != nil {
			return
		}
	}

	if len(fieldVales) == 0 {
		return
	}

	var id int
	if id, err = self.GetFieldIdValue(object); err == nil {
		return self.Update(bson.M{"id": id}, bson.M(fieldVales), tags...)
	}
	return
}

//-----------------------------------------------------------------------------------------------------------//

// 根据 ID 删除数据
func (self *ModelClass) DeleteById(id int, tag ...string) (err error) {
	if err = self.CheckId(id); err != nil {
		return
	}
	self.WithColl(func(c *mgo.Collection) {
		if err = c.Remove(bson.M{"id": id}); err != nil {
			if err == mgo.ErrNotFound {
				err = nil
			} else {
				err = obj.err.AssDelete(err, self.CollName()+".deleteById("+obj.pack.Id(id)+")", tag...)
			}
		}
	})
	return
}

func (self *ModelClass) DeleteByName(name string, chk func(name string) error, tag ...string) (int, error) {
	name = strings.TrimSpace(name)
	if chk != nil {
		if err := chk(name); err != nil {
			return 0, err
		}
	}
	if err := self.CheckObject(); err != nil {
		return 0, err
	}
	return self.Delete(bson.M{"name": name})
}

// 根据 指定 字段 匹配 删除数据
func (self *ModelClass) DeleteByFields(fields ...string) (int, error) {
	if err := self.CheckObject(); err != nil {
		return 0, err
	} else if match, err := self.GetFieldValues(fields...); err != nil {
		return 0, err
	} else {
		return self.Delete(match)
	}
}

// 根据 指定 匹配 删除数据
func (self *ModelClass) Delete(match bson.M) (num int, err error) {
	self.WithColl(func(c *mgo.Collection) {
		var info *mgo.ChangeInfo
		if info, err = c.RemoveAll(match); err != nil {
			if err == mgo.ErrNotFound {
				err = nil
			} else {
				err = obj.err.AssDelete(err, self.CollName()+".delete", obj.pack.Map(match))
			}
		} else {
			num = info.Removed
		}
	})
	return
}

//-----------------------------------------------------------------------------------------------------------//

// 通过多个 ID 来删除
func (self *ModelClass) DeleteInIds(collName string, ids ...int) (int, error) {
	if collName = strings.TrimSpace(collName); collName == "" {
		return 0, errors.New("coll name is null.")
	}
	if len(ids) == 0 {
		return 0, errors.New("ids is null.")
	}
	db := gsession.NewDB()
	defer db.Session.Close()

	if info, err := db.C(collName).RemoveAll(bson.M{"id": bson.M{"$in": ids}}); err != nil {
		err = obj.err.AssDelete(err, collName+".deleteInIds")
		return 0, err
	} else {
		return info.Removed, nil
	}
}

// 通过多个 ID 来读取
func (self *ModelClass) ReadInIds(collName string, result interface{}, ids ...int) error {
	if collName = strings.TrimSpace(collName); collName == "" {
		return errors.New("coll name is null.")
	}
	if len(ids) == 0 {
		return errors.New("ids is null.")
	}

	if result == nil {
		return errors.New("result is nil.")
	}
	value := reflect.ValueOf(result)
	if value.Kind() != reflect.Ptr {
		return errors.New("result must be ptr.")
	}
	if reflect.Indirect(value).Kind() != reflect.Slice {
		return errors.New("result must be ptr slice.")
	}

	db := gsession.NewDB()
	defer db.Session.Close()

	if err := db.C(collName).Find(bson.M{"id": bson.M{"$in": ids}}).All(result); err != nil {
		return obj.err.AssDelete(err, collName+".readInIds")
	} else {
		return nil
	}
}

//-----------------------------------------------------------------------------------------------------------//

func (self *ModelClass) GobDecode([]byte) error {
	return nil
}

func (self *ModelClass) GobEncode() ([]byte, error) {
	return nil, nil
}

//-----------------------------------------------------------------------------------------------------------//
