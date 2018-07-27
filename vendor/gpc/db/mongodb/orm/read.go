package orm

import (
	"errors"
	"reflect"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"gpc/db/mongodb/conn/gsession"
	"gpc/db/mongodb/orm/structure"
)

//-----------------------------------------------------------------------------------------------------------//

// 读
type Read struct {
	Result interface{}
	match  struct { // 匹配 ，二选一，优先 M，未指定，默认会自动使用 id 字段
		M      interface{}
		Fields []string
	}
}

func NewRead(result interface{}) *Read {
	return &Read{
		Result: result,
	}
}

// 设置 过滤 匹配 M
func (self *Read) SetMatchM(m interface{}) *Read {
	self.match.M = m
	return self
}

// 设置 过滤 匹配 字段
func (self *Read) SetMatchFields(fields ...string) *Read {
	self.match.Fields = append(self.match.Fields, fields...)
	return self
}

// 查询 读取数据【 fields 指定只读取哪些字段的数据返回回来】
func (self *Read) Query(fields ...string) error {
	if self.Result == nil {
		return errors.New("result is nil.")
	}
	value := reflect.ValueOf(self.Result)
	if value.Kind() != reflect.Ptr {
		return errors.New("result must be ptr.")
	}
	value = reflect.Indirect(value)
	isSlice := false
	switch value.Kind() {
	case reflect.Struct:
	case reflect.Slice:
		iType := value.Type().Elem()
		if iType.Kind() == reflect.Ptr {
			iType = iType.Elem()
		}
		value = reflect.New(iType)
		isSlice = true
	default:
		return errors.New("result must be struct or slice struct.")
	}

	count := len(fields)
	fieldM := make(bson.M)
	for i := 0; i < count; i++ {
		fieldM[fields[i]] = true
	}

	si, err := structure.RefValue(value)
	if err != nil {
		return err
	}

	if (!isSlice) && self.match.M == nil {
		if len(self.match.Fields) == 0 {
			self.match.Fields = []string{structure.CST_ID_NAME}
		}
		if self.match.M, err = si.FilterFields(true, self.match.Fields...); err != nil {
			return err
		}
	}

	gsession.WithDB(func(db *mgo.Database) {
		query := db.C(si.CollName).Find(self.match.M).Select(fieldM)
		if isSlice {
			err = query.All(self.Result)
		} else {
			err = query.One(self.Result)
		}
	})
	return err
}

//-----------------------------------------------------------------------------------------------------------//

// 根据 id 来读取
func ReadId(result interface{}, fields ...string) error {
	return NewRead(result).SetMatchFields(structure.CST_ID_NAME).Query(fields...)
}

//-----------------------------------------------------------------------------------------------------------//
