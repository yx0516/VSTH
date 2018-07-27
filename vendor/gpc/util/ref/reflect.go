package ref

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	plib "gpc/publib"
)

//-----------------------------------------------------------------------------------------------------------//

// 反射 接口的值
func RefValue(object interface{}) (reflect.Value, error) {
	if object == nil {
		return reflect.Value{}, errors.New("object is nil.")
	}
	if val := reflect.ValueOf(object); val.Kind() == reflect.Ptr && val.IsNil() {
		return reflect.Value{}, errors.New("object ptr value is nil.")
	} else {
		return val, nil
	}
}

// 反射 接口的值
func RefAbsValue(object interface{}) (reflect.Value, error) {
	if object == nil {
		return reflect.Value{}, errors.New("object is nil.")
	}

	val := reflect.ValueOf(object)
	for {
		if val.Kind() == reflect.Ptr {
			if val.IsNil() {
				return reflect.Value{}, errors.New("object ptr value is nil.")
			} else {
				val = val.Elem()
			}
		} else {
			break
		}
	}
	return val, nil
}

// 设置 字段 的值
func SetFieldValue(object interface{}, fieldName string, SetVal func(val reflect.Value)) error {
	if val, err := RefValue(object); err != nil {
		return err
	} else if val.Kind() == reflect.Ptr {
		if val = reflect.Indirect(val).FieldByName(fieldName); !val.IsValid() {
			return fmt.Errorf("object can not find %s field.", fieldName)
		} else if !val.CanSet() {
			return fmt.Errorf("object can not find %s field or can not set.", fieldName)
		} else {
			SetVal(val)
			return nil
		}
	} else {
		return errors.New("object must be pointer type.")
	}
}

//-----------------------------------------------------------------------------------------------------------//

// 反射结构体的信息
func RefStructInfo(object interface{}, tags ...string) (si *Struct, err error) {
	if object == nil {
		err = errors.New("object is nil.")
		return
	}

	si = &Struct{
		Fields: make(map[string]*StructField),
	}

	value := reflect.ValueOf(object)
	kind := value.Kind()
	if kind == reflect.Ptr {
		kind = value.Elem().Kind()
		si.HasPtr = true
	}
	if kind != reflect.Struct {
		err = errors.New("object is not a struct.")
		return
	}

	si.FullName = strings.TrimPrefix(value.Type().String(), "*")
	if idx := strings.LastIndex(si.FullName, "."); idx >= 0 {
		si.Name = si.FullName[idx+1:]
	}

	var realTags plib.StringArray
	for _, tag := range tags {
		tag := strings.TrimSpace(tag)
		if tag != "" {
			realTags.Add(tag)
		}
	}
	realTags.RemoveDuplicate()

	value = reflect.Indirect(value)
	iType := value.Type()
	field := ""
	count := value.NumField()

	// 开始遍历结构体的 Field
	for i := 0; i < count; i++ {
		field = iType.Field(i).Name
		if !plib.CharPrefixIsUpper(field) {
			continue
		}
		sfi := NewStructField(field, value.Field(i))
		for _, tag := range realTags {
			comment := strings.TrimSpace(HandleStructTag(iType.Field(i).Tag).Get(tag))
			if comment != "" {
				sfi.Tags[tag] = comment
			}
		}
		si.Fields[field] = sfi
	}
	return
}

//-----------------------------------------------------------------------------------------------------------//

// 处理 注解里
func HandleStructTag(tag reflect.StructTag) reflect.StructTag {
	return reflect.StructTag(strings.Join(strings.Fields(string(tag)), " "))
}

// 获取 结构体 的全名
func RefStructFullName(object interface{}) (name string, err error) {
	if object == nil {
		err = errors.New("object is nil.")
		return
	}
	value := reflect.ValueOf(object)
	kind := value.Kind()
	if kind == reflect.Ptr {
		kind = value.Elem().Kind()
	}
	if kind != reflect.Struct {
		err = errors.New("object is not a struct.")
		return
	}
	name = value.Type().String()
	return
}

// 获取 结构体 类型 全名
func RefStructTypeFullName(object interface{}) (name string, err error) {
	if name, err = RefStructFullName(object); err == nil {
		name = strings.TrimPrefix(name, "*")
	}
	return
}

// 获取 对象 类型，处理指针，直到类型
func GetAbsType(object interface{}) (reflect.Type, error) {
	if object == nil {
		return nil, errors.New("object is nil.")
	}
	t := reflect.TypeOf(object)
	for ; t.Kind() == reflect.Ptr; t = t.Elem() {
	}
	return t, nil
}

//-----------------------------------------------------------------------------------------------------------//

// 获取 结构体 类型 名称
func GetStructName(object interface{}) (name string, err error) {
	if name, err = RefStructTypeFullName(object); err == nil {
		if idx := strings.LastIndex(name, "."); idx >= 0 {
			name = name[idx+1:]
		}
	}
	return
}

// 遍历 结构体 类型 Field
func GetStructFieldsByType(rType reflect.Type, isExport bool, isRecursive bool) (fields []string) {
	if rType == nil {
		return
	}
	for ; rType.Kind() == reflect.Ptr; rType = rType.Elem() {
	}
	if rType.Kind() != reflect.Struct {
		return
	}
	count := rType.NumField()
	for i := 0; i < count; i++ {
		sf := rType.Field(i)
		field := sf.Name
		if isExport { // 是否需要大写可导出
			if !plib.CharPrefixIsUpper(field) {
				continue
			}
		}

		// 如果是嵌套字段，并且是递归
		if sf.Anonymous && isRecursive {
			st := sf.Type
			// 解指针
			for ; st.Kind() == reflect.Ptr; st = st.Elem() {
			}
			if st.Kind() == reflect.Struct {
				if arr := GetStructFieldsByType(st, isExport, isRecursive); len(arr) != 0 {
					fields = append(fields, arr...)
				}
				continue
			}
		}
		fields = append(fields, field)
	}
	return
}

// 反射 获取 结构体的 Field，不递归嵌套的Struct
func GetStructFields(object interface{}, isExport bool, isRecursive ...bool) ([]string, error) {
	if t, err := GetAbsType(object); err != nil {
		return nil, err
	} else if t.Kind() == reflect.Struct {
		hasRecursive := false
		if len(isRecursive) > 0 {
			hasRecursive = isRecursive[0]
		}
		return GetStructFieldsByType(t, isExport, hasRecursive), nil
	} else {
		return nil, errors.New("object is not struct.")
	}
}

// 获取 结构体 指定字段的值
func GetStructFieldValueIntf(object interface{}, field string) (val interface{}, err error) {
	if object == nil {
		err = errors.New("object is nil.")
		return
	}
	if field = strings.TrimSpace(field); field == "" {
		err = errors.New("field name is null.")
		return
	}
	value := reflect.Indirect(reflect.ValueOf(object))
	if value.Kind() != reflect.Struct {
		err = errors.New("object is not struct.")
		return
	}
	valValue := value.FieldByName(field)
	if !valValue.IsValid() {
		err = errors.New(field + " field not found.")
		return
	} else if !valValue.CanInterface() {
		err = errors.New(field + " field can not interface.")
		return
	}
	return valValue.Interface(), nil
}

// 过滤结构体字段
func FilterFields(object interface{}, fileds ...string) (interface{}, error) {
	if object == nil {
		return nil, errors.New("object is nil.")
	}
	count := len(fileds)
	if count > 0 {
		plib.StrArrayRemoveDuplicate(&fileds)
		plib.StrArrayDeleteAllSpaceTrim(&fileds)
	}
	if count = len(fileds); count == 0 {
		return object, nil
	}
	filter := func(val reflect.Value) (interface{}, error) {
		if val = reflect.Indirect(val); val.Kind() != reflect.Struct {
			return nil, errors.New("object is not struct.")
		}
		fieldValues := make(map[string]interface{})
		for _, name := range fileds {
			if fVal := val.FieldByName(name); fVal.IsValid() && fVal.CanInterface() {
				fieldValues[name] = fVal.Interface()
			}
		}
		return fieldValues, nil
	}
	value := reflect.ValueOf(object)
loop:
	switch value.Kind() {
	case reflect.Ptr:
		value = value.Elem()
		goto loop
	case reflect.Slice:
		var arr []interface{}
		count := value.Len()
		for i := 0; i < count; i++ {
			if fvs, err := filter(value.Index(i)); err != nil {
				return nil, err
			} else {
				arr = append(arr, fvs)
			}
		}
		return arr, nil
	case reflect.Struct:
		return filter(value)
	}
	return nil, errors.New("unsupported type:" + value.Kind().String())
}

//-----------------------------------------------------------------------------------------------------------//

// 获取 结构体 指定方法
func GetStructMethod(object interface{}, methodName string) (methodValue reflect.Value, err error) {
	if methodName = strings.TrimSpace(methodName); methodName == "" {
		err = errors.New("method name is null.")
		return
	}
	if object == nil {
		err = errors.New("object is nil.")
		return
	}
	value := reflect.ValueOf(object)
	kind := value.Kind()
	isPtr := false
	if kind == reflect.Ptr {
		kind = value.Elem().Kind()
		isPtr = true
	}
	if kind != reflect.Struct {
		err = errors.New("object is not a struct.")
		return
	}
	if isPtr {
		methodValue = value.MethodByName(methodName)
	} else if methodValue = value.MethodByName(methodName); !methodValue.IsValid() {
		ptr := reflect.New(reflect.TypeOf(object))
		ptr.Elem().Set(value)
		methodValue = ptr.MethodByName(methodName)
	}
	if !methodValue.IsValid() {
		err = errors.New("method name not found.(" + methodName + ")")
	}
	return
}

// 调用 结构体 的方法
func CallStructMethod(object interface{}, methodName string, args ...interface{}) (retValues []reflect.Value, err error) {
	var methodValue reflect.Value
	if methodValue, err = GetStructMethod(object, methodName); err != nil {
		return
	}
	defer func() {
		if rErr := recover(); rErr != nil {
			if msg, ok := rErr.(string); ok {
				err = errors.New(msg + "(panic).")
			} else {
				err = errors.New("call func input arguments error(panic).")
			}
		}
	}()
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	retValues = methodValue.Call(inputs)

	return
}

//-----------------------------------------------------------------------------------------------------------//
