package tag

import (
	"reflect"

	"gpc/util/ref"

	plib "gpc/publib"
)

//-----------------------------------------------------------------------------------------------------------//

// 如果 字段是 空值，则设置默认值
func (self *FieldCheck) SetNullDefValue(sf *ref.StructField, tags ...string) (ret *RetStatus) {
	if ret = self.IsNull(sf, tags...); !ret.IsOK() {
		return
	}
	defer func() {
		if !ret.IsOK() {
			ret.Msg = sf.Name + " field " + ret.Msg
		}
	}()

	defValStr, ok := sf.Tags[CST_TAG_DEF]
	if !ok {
		return ret.SetErr("not def tag.")
	}
	if !sf.Value.CanSet() {
		return ret.SetErr("value can not set.")
	}
	kind := sf.Value.Kind()
	switch kind {
	case reflect.String:
		sf.Value.SetString(defValStr)
		return
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int, reflect.Int64:
		if x, tErr := plib.StrToInt64(defValStr); tErr != nil {
			return ret.SetErr(tErr.Error())
		} else {
			sf.Value.SetInt(x)
		}
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint, reflect.Uint64:
		if x, tErr := plib.StrToUInt64(defValStr); tErr != nil {
			return ret.SetErr(tErr.Error())
		} else {
			sf.Value.SetUint(x)
		}
	case reflect.Float32, reflect.Float64:
		if x, tErr := plib.StrToFloat(defValStr); tErr != nil {
			return ret.SetErr(tErr.Error())
		} else {
			sf.Value.SetFloat(x)
		}
	default:
		return ret.SetErr("unsupported type:" + kind.String())
	}
	return
}

//-----------------------------------------------------------------------------------------------------------//
