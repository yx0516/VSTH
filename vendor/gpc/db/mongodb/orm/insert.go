package orm

import (
	"errors"
	"reflect"

	"gopkg.in/mgo.v2"

	"gpc/db/mongodb/conn/gsession"
	"gpc/db/mongodb/orm/structure"
	"gpc/util/ref"
)

//-----------------------------------------------------------------------------------------------------------//

// 插入 结构体【 objetc 可以是 结构体 也可以是 结构体Slice】
func Insert(objetc interface{}) error {
	value, err := ref.RefAbsValue(objetc)
	if err != nil {
		return err
	}

	var docs []interface{}
	switch value.Kind() {
	case reflect.Struct:
		docs = append(docs, objetc)
	case reflect.Slice:
		count := value.Len()
		for i := 0; i < count; i++ {
			docs = append(docs, value.Index(i).Interface())
		}

		iType := value.Type().Elem()
		if iType.Kind() == reflect.Ptr {
			iType = iType.Elem()
		}
		value = reflect.New(iType)
	default:
		return errors.New("objetc type must be struct or slice struct.")
	}

	si, err := structure.RefValue(value)
	if err != nil {
		return err
	}

	gsession.WithDB(func(db *mgo.Database) {
		err = db.C(si.CollName).Insert(docs...)
	})
	return err
}

//-----------------------------------------------------------------------------------------------------------//
