package tag

import (
	"errors"
	"reflect"
	"strings"

	"gpc/util/ref"
)

//-----------------------------------------------------------------------------------------------------------//

// 如果是字符串类型，则删除头尾空格
func (self *FieldCheck) StringTrim(sf *ref.StructField) error {
	if sf == nil {
		return errors.New("sf is nil.")
	}

	if sf.Value.Kind() != reflect.String {
		return nil
	}

	if isNoTrim(sf) {
		return nil
	}

	if !sf.Value.CanSet() {
		return errors.New(sf.Name + " field value can not set.")
	}

	if s, ok := sf.Value.Interface().(string); ok {
		sf.Value.SetString(strings.TrimSpace(s))
	}
	return nil
}

//-----------------------------------------------------------------------------------------------------------//
