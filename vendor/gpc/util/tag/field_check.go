package tag

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"gpc/util/ref"
)

//-----------------------------------------------------------------------------------------------------------//

// 检查
func (self *FieldCheck) Check(sf *ref.StructField, tags ...string) (ret *RetStatus) {
	ret = NewRetStatus()
	if sf == nil {
		return ret.SetErr("sf is nil.")
	}
	defer func() {
		if !ret.IsOK() {
			ret.Msg = sf.Name + " field " + ret.Msg
		}
	}()

	if !sf.Value.CanInterface() {
		return ret.SetErr("value can not interface.")
	}

	if len(tags) == 0 {
		tags = append(tags, self.TagChk)
	}
	realTags := sf.FilterTags(tags...)
	if len(realTags) == 0 {
		return ret.SetErr("tags not set:" + strings.Join(tags, ","))
	}

	tags = []string{}
	for _, tag := range realTags {
		if exp, ok := sf.Tags[tag]; ok {
			if exp != CST_EXP_IGNORE {
				tags = append(tags, tag)
			}
		}
	}

	if len(tags) == 0 {
		return
	}

	kind := sf.Value.Kind()
	switch kind {
	case reflect.String:
		ret = self.checkString(sf, tags...)
	case reflect.Ptr:
		ret = self.checkPtr(sf, tags...)
	case reflect.Map:
		ret = self.checkMap(sf, tags...)
	case reflect.Interface, reflect.Struct:
		ret.SetErr("unsupported type:" + kind.String())
	default:
		ret = self.checkNum(sf, tags...)
	}
	return
}

// 检查 Map 类型
func (self *FieldCheck) checkMap(sf *ref.StructField, tags ...string) (ret *RetStatus) {
	ret = NewRetStatus()
	for _, tag := range tags {
		exps, isOr := splitTag(sf.Tags[tag])
		for _, exp := range exps {
			switch exp {
			case "!nil":
				if sf.Value.IsNil() {
					if !isOr {
						return ret.SetFail("is nil.")
					}
				} else if isOr {
					return
				}
			case "nil":
				if !sf.Value.IsNil() {
					if !isOr {
						return ret.SetFail("is not nil.")
					}
				} else if isOr {
					return
				}
			default:
				if strings.HasPrefix(exp, "len") {
					if oRet := self.Operator(int64(sf.Value.Len()), exp); !oRet.IsOK() {
						if !isOr {
							return oRet
						}
					} else if isOr {
						return
					}
				} else {
					return ret.SetErr("unsupported expression:" + exp)
				}
			}
		}
		if isOr {
			return ret.SetFail("no match:" + sf.Tags[tag])
		}
	}
	return
}

// 检查 指针 类型
func (self *FieldCheck) checkPtr(sf *ref.StructField, tags ...string) (ret *RetStatus) {
	ret = NewRetStatus()
	for _, tag := range tags {
		exps, isOr := splitTag(sf.Tags[tag])
		for _, exp := range exps {
			switch exp {
			case "!nil":
				if sf.Value.IsNil() {
					if !isOr {
						return ret.SetFail("is nil.")
					}
				} else if isOr {
					return
				}
			case "nil":
				if !sf.Value.IsNil() {
					if !isOr {
						return ret.SetFail("is not nil.")
					}
				} else if isOr {
					return
				}
			default:
				return ret.SetErr("unsupported expression:" + exp)
			}
		}
		if isOr {
			return ret.SetFail("no match:" + sf.Tags[tag])
		}
	}
	return
}

// 检查 字符串 类型
func (self *FieldCheck) checkString(sf *ref.StructField, tags ...string) (ret *RetStatus) {
	ret = NewRetStatus()
	val, ok := sf.Value.Interface().(string)
	if !ok {
		return ret.SetErr("value failed to convert a string type.")
	}

	if !isNoTrim(sf) {
		val = strings.TrimSpace(val)
	}

	for _, tag := range tags {
		exps, isOr := splitTag(sf.Tags[tag])
		for _, exp := range exps {
			switch exp {
			case "!null":
				if val == "" {
					if !isOr {
						return ret.SetFail("is null.")
					}
				} else if isOr {
					return
				}
			case "null":
				if val != "" {
					if !isOr {
						return ret.SetFail("is not null.")
					}
				} else if isOr {
					return
				}
			default:
				oRet := NewRetStatus()
				if strings.HasPrefix(exp, "len") {
					oRet = self.Operator(int64(len(val)), exp)
				} else {
					oRet = self.Operator(val, exp)
				}
				if !oRet.IsOK() {
					if !isOr {
						return oRet
					}
				} else if isOr {
					return
				}
			}
		}
		if isOr {
			return ret.SetFail("no match:" + sf.Tags[tag])
		}
	}
	return
}

// 检查 数字 类型
func (self *FieldCheck) checkNum(sf *ref.StructField, tags ...string) (ret *RetStatus) {
	ret = NewRetStatus()
	var num int64

	val := sf.Value.Interface()
	str := fmt.Sprintf("%v", val)

	kind := sf.Value.Kind()
	switch kind {
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int, reflect.Int64:
		if vv, err := strconv.ParseInt(str, 10, 64); err != nil {
			return ret.SetErr(fmt.Sprintf("int(%v) unsupported type:%s", val, kind.String()))
		} else {
			num = int64(vv)
		}
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint, reflect.Uint64:
		if vv, err := strconv.ParseUint(str, 10, 64); err != nil {
			return ret.SetErr(fmt.Sprintf("uint(%v) unsupported type:%s", val, kind.String()))
		} else {
			num = int64(vv)
		}
	case reflect.Float32, reflect.Float64:
		if vv, err := strconv.ParseFloat(str, 64); err != nil {
			return ret.SetErr(fmt.Sprintf("float(%v) unsupported type:%s", val, kind.String()))
		} else {
			num = int64(vv)
		}
	default:
		return ret.SetErr("unsupported type:" + kind.String())
	}

	for _, tag := range tags {
		exps, isOr := splitTag(sf.Tags[tag])
		for _, exp := range exps {
			if oRet := self.Operator(num, exp); !oRet.IsOK() {
				if !isOr {
					return oRet
				}
			} else if isOr {
				return
			}
		}
		if isOr {
			return ret.SetFail("no match:" + sf.Tags[tag])
		}
	}
	return
}

//-----------------------------------------------------------------------------------------------------------//
