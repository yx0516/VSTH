package tag

import (
	"reflect"
	"strings"

	"gpc/util/ref"
)

//-----------------------------------------------------------------------------------------------------------//

// 注解 检查
type FieldCheck struct {
	TagChk     string
	TagDef     string
	TagNull    string
	TagNotNull string
	TagVal     string
}

func NewFieldCheck() *FieldCheck {
	return &FieldCheck{
		TagChk:     CST_TAG_CHK,
		TagNull:    CST_TAG_NULL,
		TagNotNull: CST_TAG_NOT_NULL,
		TagDef:     CST_TAG_DEF,
		TagVal:     CST_TAG_VAL,
	}
}

//-----------------------------------------------------------------------------------------------------------//

func (self *FieldCheck) GetTags() []string {
	return []string{self.TagChk, self.TagDef, self.TagNotNull, self.TagNull, self.TagVal}
}

func (self *FieldCheck) UpdateStructFieldVal(sf *ref.StructField) {
	if sf != nil && sf.Tags != nil && sf.Value.CanInterface() {
		if valField, ok := sf.Tags[self.TagVal]; ok && valField != "" {
			sf.Value = reflect.ValueOf(self.GetStructFieldValByValue(sf.Value, valField))
		}
	}
}

func (self *FieldCheck) GetStructFieldVal(sf *ref.StructField) interface{} {
	if sf != nil {
		if sf.Tags != nil {
			if valField, ok := sf.Tags[self.TagVal]; ok && valField != "" {
				return self.GetStructFieldValByValue(sf.Value, valField)
			}
		}
		if sf.Value.CanInterface() {
			return sf.Value.Interface()
		}
	}
	return nil
}

// 获取 结构体 自定义 字段的值
func (self *FieldCheck) GetStructFieldValByValue(value reflect.Value, filedName string) (val interface{}) {
	arr := strings.Split(filedName, ";")
	field := strings.TrimSpace(arr[0])
	valType := ""
	if len(arr) > 1 {
		valType = strings.TrimSpace(arr[1])
	}

	isOk := false
	defer func() {
		if !isOk {
			switch valType {
			case CST_VAL_TYPE_NUM:
				val = 0
			case CST_VAL_TYPE_STR:
				val = ""
			}
		}
	}()

	if value.IsValid() && value.CanInterface() {
		if val = value.Interface(); val != nil {
			if value = reflect.Indirect(reflect.ValueOf(val)); value.Kind() == reflect.Struct {
				if value = value.FieldByName(field); value.IsValid() && value.CanInterface() {
					isOk = true
					return value.Interface()
				}
			}
		}
	}
	return
}

//-----------------------------------------------------------------------------------------------------------//
