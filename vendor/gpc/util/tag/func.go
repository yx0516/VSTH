package tag

import (
	"errors"
	"strings"

	"gpc/util/ref"
	"gpc/web/cst"

	plib "gpc/publib"
)

//-----------------------------------------------------------------------------------------------------------//

// 返回 不同值的字段名列表
func DiffValueFields(st1, st2 *StructTag, fields ...string) (diffFields plib.StringArray, err error) {
	if st1 == nil || st1 == nil || st1.Struct == nil || st2.Struct == nil || st1.Struct.FullName != st2.Struct.FullName {
		return nil, errors.New("struct is nil OR diff struct.")
	}
	diff := func(sf1, sf2 *ref.StructField) bool {
		return sf1.Value.CanInterface() && sf2.Value.CanInterface() && sf1.Value.Interface() != sf2.Value.Interface()
	}
	if len(fields) > 0 {
		plib.StrArrayRemoveDuplicate(&fields)
		for _, field := range fields {
			if sf1 := st1.GetFieldInfo(field); sf1 != nil {
				if sf2 := st2.GetFieldInfo(field); sf2 != nil {
					if diff(sf1, sf2) {
						diffFields.Add(field)
					}
				}
			}
		}
	} else {
		for field, sf1 := range st1.Fields {
			if sf2, ok := st2.Fields[field]; ok {
				if diff(sf1, sf2) {
					diffFields.Add(field)
				}
			}
		}
	}
	return
}

// 获取 Beego orm 的外键字段名
func GetBeegoOrmFKfields(st *StructTag) (fields []string, err error) {
	if st == nil || st.Fields == nil {
		return nil, errors.New("st is nil.")
	}
	for _, sf := range st.Fields {
		items := strings.Split(sf.GetTagBody(cst.TAG_BEEGO_ORM), ";")
		for _, item := range items {
			if strings.Join(strings.Fields(item), "") == "rel(fk)" {
				fields = append(fields, sf.Name)
				break
			}
		}
	}
	return
}

//-----------------------------------------------------------------------------------------------------------//

// 获取 结构体 的名称，如果有指定的方法名，则优先调用方法名获取
func GetStructName(v interface{}, methodName string, convert func(name string) string) (string, error) {
	if v == nil {
		return "", errors.New("v is nil.")
	}

	name := ""

	if methodName = strings.TrimSpace(methodName); methodName != "" {
		if vals, err := ref.CallStructMethod(v, methodName); err != nil {
			return "", err
		} else if len(vals) > 0 && vals[0].IsValid() && vals[0].CanInterface() {
			ok := false
			if name, ok = vals[0].Interface().(string); !ok {
				return "", errors.New("method ret val can not convert string.")
			} else if name = strings.TrimSpace(name); name == "" {
				return "", errors.New("method ret val string is null.")
			}
		} else {
			return "", errors.New("method ret val null OR invalid.")
		}
	} else {
		var err error
		if name, err = ref.GetStructName(v); err != nil {
			return "", err
		} else if convert != nil {
			if name = strings.TrimSpace(convert(name)); name == "" {
				return "", errors.New("name convert is null.")
			}
		}
	}

	return name, nil
}

//-----------------------------------------------------------------------------------------------------------//

// 按 ; 分隔符 分割 tag 里的内容
func splitTag(tag string) (exps plib.StringArray, isOr bool) {
	tag = strings.Join(strings.Fields(tag), "")
	switch {
	case strings.HasPrefix(tag, CST_LOGIC_AND):
		tag = strings.TrimPrefix(tag, CST_LOGIC_AND)
	case strings.HasPrefix(tag, CST_LOGIC_OR):
		tag = strings.TrimPrefix(tag, CST_LOGIC_OR)
		isOr = true
	}
	arr := strings.Split(tag, ";") // 按 ; 分割
	for _, exp := range arr {
		exp = strings.TrimSpace(exp)
		if exp != "" {
			exps.Add(exp)
		}
	}
	exps.RemoveDuplicate()
	return
}

// 是否 不执行 trim 操作
func isNoTrim(sf *ref.StructField) (ok bool) {
	if body := sf.GetTagBody(CST_TAG_CFG); body != "" {
		exps, _ := splitTag(body)
		ok = exps.Contains(CST_CFG_NOT_TRIM)
	}
	return
}

//-----------------------------------------------------------------------------------------------------------//
