package xorm

import (
	"errors"
	"reflect"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"gpc/db/mongodb/utils"
	"gpc/util/ref"

	mdl "gpc/db/mongodb/xorm/model"
	plib "gpc/publib"
	deep "gpc/util/deep/json"
)

//-----------------------------------------------------------------------------------------------------------//

// 模型 公用
type ModelPublic struct {
	Model ModelClass `bson:"-" json:"-"`
	// 是否 禁止 这些功能【Init函数调用设置】
	disableInsert error `bson:"-" json:"-"`
	disableRead   error `bson:"-" json:"-"`
	disableUpdate error `bson:"-" json:"-"`
	disableDelete error `bson:"-" json:"-"`
}

func NewModelPublic(model mdl.Model) *ModelPublic {
	return &ModelPublic{
		Model: *NewModelClass(model),
	}
}

// 注册
func (self *ModelPublic) Init(model mdl.Model) {
	self.Model.init(model)
}

//-----------------------------------------------------------------------------------------------------------//

// 禁止 添加
func (self *ModelPublic) DisableInsert(err ...error) *ModelPublic {
	if len(err) > 0 {
		self.disableInsert = err[0]
	} else {
		self.disableInsert = NewFuncNotImplemented()
	}
	return self
}

// 禁止 读取
func (self *ModelPublic) DisableRead(err ...error) *ModelPublic {
	if len(err) > 0 {
		self.disableRead = err[0]
	} else {
		self.disableRead = NewFuncNotImplemented()
	}
	return self
}

// 禁止 更新
func (self *ModelPublic) DisableUpdate(err ...error) *ModelPublic {
	if len(err) > 0 {
		self.disableUpdate = err[0]
	} else {
		self.disableUpdate = NewFuncNotImplemented()
	}
	return self
}

// 禁止 删除
func (self *ModelPublic) DisableDelete(err ...error) *ModelPublic {
	if len(err) > 0 {
		self.disableDelete = err[0]
	} else {
		self.disableDelete = NewFuncNotImplemented()
	}
	return self
}

//-----------------------------------------------------------------------------------------------------------//

// 包装 检查 是否 禁止 Add
func (self *ModelPublic) WithInsert(fun func() (int, error)) (int, error) {
	if self.disableInsert == nil {
		return fun()
	} else {
		return 0, self.disableInsert
	}
}

// 包装 检查 是否 禁止 Read
func (self *ModelPublic) WithRead(fun func(fields ...string) error, fields ...string) error {
	if self.disableRead == nil {
		return fun(fields...)
	} else {
		return self.disableRead
	}
}

// 包装 检查 是否 禁止 Update
func (self *ModelPublic) WithUpdate(fun func() error) error {
	if self.disableUpdate == nil {
		return fun()
	} else {
		return self.disableUpdate
	}
}

// 包装 检查 是否 禁止 Delete
func (self *ModelPublic) WithDelete(fun func() error) error {
	if self.disableDelete == nil {
		return fun()
	} else {
		return self.disableDelete
	}
}

//-----------------------------------------------------------------------------------------------------------//

// 检查 ID 是否存在
func (self *ModelPublic) ExistId() (ok bool, err error) {
	var id int
	if id, err = self.Model.VerifyId(); err == nil {
		return self.Model.Exist(bson.M{"id": id})
	}
	return
}

//-----------------------------------------------------------------------------------------------------------//

// 插入数据
func (self *ModelPublic) Insert() (int, error) {
	return self.WithInsert(func() (id int, err error) {
		var intf mdl.Model
		if intf, err = self.Copy(); err != nil {
			return
		}

		if err = intf.CheckInsert(); err == nil { // 检查参数【子类实现】
			self.Model.WithColl(func(c *mgo.Collection) {
				id, err = utils.GenNextId(c.Database.C(CST_COLL_SYS_ID_COUNTER), intf.GetCollName())
				if err == nil {
					if err = self.Model.SetFieldIdValue(intf, id); err == nil {
						if err = c.Insert(intf); err == nil {
							err = intf.CopyTo(self.Model.object)
						}
					}
				}
			})
		}
		return
	})
}

//  执行 Insert 时，调用检查参数的函数【子类 重写】
func (self *ModelPublic) CheckInsert() (err error) {
	if _, err = self.Model.TrimAllSetDefVerify(); err != nil {
		return
	}

	if self.Model.FuncCheckInsert != nil {
		return self.Model.FuncCheckInsert()
	}
	return
}

// 查询条件 id 或 X , 优先使用 id【子类 重写】
func (self *ModelPublic) MakeMatch() (bson.M, error) {
	id := self.GetFieldIdValue()
	if err := obj.chk.Id(id); err == nil {
		return bson.M{"id": id}, nil
	} else if self.Model.FuncMakeMatch != nil {
		return self.Model.FuncMakeMatch()
	} else {
		return nil, err
	}
}

//-----------------------------------------------------------------------------------------------------------//

// 复制 读取 数据
func (self *ModelPublic) CopyRead(fields ...string) (intf mdl.Model, err error) {
	err = self.WithRead(func(fields ...string) error {
		intf, err = self.Model.CopyRead(fields...)
		return err
	}, fields...)

	return
}

// 读取 数据
func (self *ModelPublic) Read(fields ...string) error {
	return self.WithRead(func(fields ...string) error {
		if err := self.Model.CheckObject(); err == nil {
			return self.Model.ReadOne(self.Model.object, fields...)
		} else {
			return err
		}
	}, fields...)
}

// 根据 ID 读取数据
func (self *ModelPublic) ReadById(fields ...string) error {
	return self.WithRead(func(fields ...string) error {
		if id, err := self.Model.VerifyId(); err == nil {
			return self.Model.ReadById(id, fields...)
		} else {
			return err
		}
	}, fields...)
}

//-----------------------------------------------------------------------------------------------------------//

// 更新 匹配条件 指定字段【fields 未过滤的】
func (self *ModelPublic) UpdateFields(match bson.M, fields ...string) error {
	return self.WithUpdate(func() error {
		return self.Model.UpdateFields(match, fields...)
	})
}

// 更新 匹配条件 字段值
func (self *ModelPublic) UpdateFieldValues(match, fieldValues bson.M) error {
	return self.WithUpdate(func() error {
		return self.Model.Update(match, fieldValues)
	})
}

// 自动 更新 非空字段
func (self *ModelPublic) Update() error {
	return self.WithUpdate(func() error {
		if fields, err := self.Model.GetNotNullUpdateFields(); err != nil {
			return err
		} else {
			return self.Model.object.UpdateByFields(fields) // 调 子类的 UpdateByFields 方法
		}
	})
}

// 更新 指定 字段【转换后的字段】【子类重写】
func (self *ModelPublic) UpdateByFields(fields plib.StringArray) error {
	/*
		fun := func(fvs xorm.MapSI, vSelf interface{}) (err error) {
			v, ok := vSelf.(*Template) // vSelf 是 Copy 读取数据后的对象
			if !ok {
				return errors.New("vSelf interface convert fail.")
			}

			if fvs.Len() > 0 {
				fvs[cst.Field_UpdateTime] = time.Now()
			}

			return
		}
	*/
	return self.WithUpdate(func() error {
		return self.Model.UpdateByFields(fields, self.Model.FuncUpdateByFields)
	})
}

//-----------------------------------------------------------------------------------------------------------//

// 根据 读取的 Id 来删除
func (self *ModelPublic) Delete() error {
	return self.WithDelete(func() error {
		if intf, err := self.Model.CopyRead(); err != nil {
			return err
		} else {
			return intf.DeleteById() // 调用 子类实现的删除
		}
	})
}

// 通过 Id 来删除
func (self *ModelPublic) DeleteById() error {
	return self.WithDelete(func() error {
		id := self.GetFieldIdValue()
		return self.Model.DeleteById(id)
	})
}

//-----------------------------------------------------------------------------------------------------------//

// 通过多个 ID 来读取
func (self *ModelPublic) ReadInIds(result interface{}, ids ...int) error {
	if self.disableRead == nil {
		if name, err := self.Model.GetCollName(); err != nil {
			return err
		} else {
			return self.Model.ReadInIds(name, result, ids...)
		}
	} else {
		return self.disableRead
	}
}

// 通过多个 ID 来删除
func (self *ModelPublic) DeleteInIds(ids ...int) (int, error) {
	if self.disableRead == nil {
		if name, err := self.Model.GetCollName(); err != nil {
			return 0, err
		} else {
			return self.Model.DeleteInIds(name, ids...)
		}
	} else {
		return 0, self.disableRead
	}
}

//-----------------------------------------------------------------------------------------------------------//

func (self *ModelPublic) StringPretty() string {
	return plib.JsonMarshalPrettyToString(self.Model.object)
}

// 获取 集合
func (self *ModelPublic) GetCollName() (name string) {
	name, _ = self.Model.GetCollName()
	return
}

// 获取 结构体 类型 名称
func (self *ModelPublic) GetModelName() (name string) {
	name, _ = ref.GetStructName(self.Model.object)
	return
}

// 获取 结构体 类型 全名【不区分指针】
func (self *ModelPublic) GetModelFullName() (name string) {
	name, _ = ref.RefStructTypeFullName(self.Model.object)
	return
}

// 获取 字段 Id 的值
func (self *ModelPublic) GetFieldIdValue() (id int) {
	id, _ = self.Model.GetFieldIdValue(self.Model.object)
	return
}

// 设置 字段 Id 的值
func (self *ModelPublic) SetFieldIdValue(id int) error {
	return self.Model.SetFieldIdValue(self.Model.object, id)
}

// 获取 字段 Name 的值
func (self *ModelPublic) GetFieldNameValue() (name string) {
	name, _ = self.Model.GetFieldNameValue(self.Model.object)
	return
}

// 设置 字段 Name 的值
func (self *ModelPublic) SetFieldNameValue(name string) error {
	return self.Model.SetFieldNameValue(self.Model.object, name)
}

//-----------------------------------------------------------------------------------------------------------//

// 生成一个 结构体指针 模型 实例 *Model
func (self *ModelPublic) MakeModelPtr() mdl.Model {
	model, _ := self.Model.MakeModelPtr()
	return model
}

// 生成一个 结构体Slice指针 实例 *[]Model
func (self *ModelPublic) MakeModelSlicePtr() interface{} {
	modelSlicePtr, _ := self.Model.MakeStructSlicePtr()
	return modelSlicePtr
}

// 获取 结构体 类型 Model 【非指针】
func (self *ModelPublic) GetModelType() reflect.Type {
	vType, _ := self.Model.GetStructType()
	return vType
}

//-----------------------------------------------------------------------------------------------------------//

// 复制 对象【返回的是指针】
func (self *ModelPublic) Copy() (mdl.Model, error) {
	if v, err := deep.CopyValue(self.Model.object); err != nil {
		return nil, err
	} else if m, ok := v.(mdl.Model); ok {
		m.Init()
		return m, nil
	} else {
		return nil, errors.New("model interface convert fail.")
	}
}

// 复制对象到
func (self *ModelPublic) CopyTo(model mdl.Model) (err error) {
	if err = deep.Copy(self.Model.object, model); err == nil {
		model.Init()
	}
	return
}

//-----------------------------------------------------------------------------------------------------------//
