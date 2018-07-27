package structure

import (
	"errors"
	"reflect"
	"regexp"
	"strings"

	"gopkg.in/mgo.v2/bson"

	plib "gpc/publib"
	"gpc/util/ref"
)

//-----------------------------------------------------------------------------------------------------------//

// 只能是 字母 和 _【并且不能是空】
func CheckFieldName(s string) (ok bool) {
	ok, _ = regexp.MatchString(`^[a-zA-Z\_]+$`, s)
	return
}

//-----------------------------------------------------------------------------------------------------------//

// 获取 结构体 的信息
func Ref(objetc interface{}) (si *StructInfo, err error) {
	if objetc == nil {
		err = errors.New("objetc is nil.")
		return
	}
	value := reflect.ValueOf(objetc)
	return refStructValueInfo(value)
}

func RefValue(value reflect.Value) (si *StructInfo, err error) {
	if !value.IsValid() {
		return nil, errors.New("value is invalid.")
	}
	return refStructValueInfo(value)
}

// 获取 结构体 值 的信息
func refStructValueInfo(value reflect.Value) (*StructInfo, error) {
	kind := value.Kind()
	isPtr := false
	if kind == reflect.Ptr {
		if value.IsNil() {
			return nil, errors.New("value is nil ptr.")
		}

		kind = value.Elem().Kind()
		isPtr = true
	}
	if kind != reflect.Struct {
		return nil, errors.New("value is not a struct.")
	}

	si := &StructInfo{
		FullName: strings.TrimPrefix(value.Type().String(), "*"),
	}
	idx := strings.LastIndex(si.FullName, ".")
	if idx >= 0 {
		si.Name = si.FullName[idx+1:]
	}
	if si.Name == "" {
		return nil, errors.New("value struct name invalid.")
	}

	var methodValue reflect.Value

	if isPtr {
		methodValue = value.MethodByName(CST_METHOD_COLL_NAME)
	} else if methodValue = value.MethodByName(CST_METHOD_COLL_NAME); !methodValue.IsValid() {
		ptr := reflect.New(value.Type())
		ptr.Elem().Set(value)
		methodValue = ptr.MethodByName(CST_METHOD_COLL_NAME)
	}
	if !methodValue.IsValid() {
		goto SetDefCollName
	} else {
		vs := methodValue.Call(nil)
		if len(vs) != 0 && vs[0].CanInterface() {
			if name, ok := vs[0].Interface().(string); ok {
				name := strings.TrimSpace(name)
				if name == "" {
					goto SetDefCollName
				} else {
					si.CollName = name
					goto GetFieldInfo
				}
			} else {
				goto SetDefCollName
			}
		} else {
			goto SetDefCollName
		}
	}
SetDefCollName:
	si.CollName = strings.ToLower(si.Name)
GetFieldInfo:
	value = reflect.Indirect(value)
	valueType := value.Type()
	count := value.NumField()
	for i := 0; i < count; i++ {
		field := valueType.Field(i).Name
		if !plib.CharPrefixIsUpper(field) {
			continue
		}
		fi := &FieldInfo{
			Name:  field,
			Value: value.Field(i),
		}

		comment := strings.TrimSpace(ref.HandleStructTag(valueType.Field(i).Tag).Get(CST_TAG_NAME))
		if comment == "-" {
			continue
		}

		if CheckFieldName(comment) {
			fi.FieldName = comment
		} else {
			fi.FieldName = strings.ToLower(fi.Name)
		}
		si.FieldInfos = append(si.FieldInfos, fi)
	}

	return si, nil
}

//-----------------------------------------------------------------------------------------------------------//

// 获取 结构体 的信息 和 匹配条件
func GetMatch(v interface{}, match interface{}, filterFields ...string) (si *StructInfo, matchM interface{},
	fieldM bson.M, err error) {
	if si, err = Ref(v); err != nil {
		return
	}
	if fieldM, err = si.FilterFields(true, filterFields...); err != nil {
		return
	}
	if match == nil {
		if mId, err := si.GetId(); err != nil {
			return nil, nil, nil, err
		} else {
			matchM = mId
		}
	} else {
		matchM = match
	}
	return
}

//-----------------------------------------------------------------------------------------------------------//
