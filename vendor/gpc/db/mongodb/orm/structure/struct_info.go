package structure

import (
	"errors"
	"reflect"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

const (
	CST_METHOD_COLL_NAME string = "CollName"
	CST_TAG_NAME         string = "bson"
	CST_ID_NAME          string = "id"
)

//-----------------------------------------------------------------------------------------------------------//

// 结构体 属性 信息
type FieldInfo struct {
	Name      string
	FieldName string
	Value     reflect.Value
}

// 结构体 信息
type StructInfo struct {
	FullName   string
	Name       string
	CollName   string
	FieldInfos []*FieldInfo
}

// 集合的属性名 Map
func (self *StructInfo) FilterFields(isCheck bool, fields ...string) (m bson.M, err error) {
	m = make(bson.M)
	count := len(fields)
	if count > 0 {
		for i := 0; i < count; i++ {
			fields[i] = strings.ToLower(fields[i])
		}

		for _, fi := range self.FieldInfos {
			for i := 0; i < count; i++ {
				if fields[i] == fi.FieldName {
					m[fi.FieldName] = fi.Value.Interface()
					break
				}
			}
		}
		if isCheck && len(m) != count {
			err = errors.New("fields invalid:" + strings.Join(fields, "|"))
			return
		}
	} else {
		for _, fi := range self.FieldInfos {
			m[fi.FieldName] = fi.Value.Interface()
		}
	}
	return
}

// 获取 ID 的 bson.M
func (self *StructInfo) GetId() (m bson.M, err error) {
	m = make(bson.M)
	for _, fi := range self.FieldInfos {
		if fi.FieldName == CST_ID_NAME {
			m[fi.FieldName] = fi.Value.Interface()
			return
		}
	}
	if len(m) == 0 {
		err = errors.New(CST_ID_NAME + " field not found.")
	}
	return
}

//-----------------------------------------------------------------------------------------------------------//
