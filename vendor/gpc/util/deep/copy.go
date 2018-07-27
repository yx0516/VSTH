package deep

import (
	"reflect"
)

//-----------------------------------------------------------------------------------------------------------//

func Copy(old interface{}, copyFun func(src, dst interface{}) error) (interface{}, error) {
	if old == nil {
		return nil, nil
	}
	oldValue := reflect.ValueOf(old)
	if oldValue.Kind() == reflect.Ptr && oldValue.IsNil() {
		return oldValue.Interface(), nil
	}
	newValue := reflect.New(oldValue.Type())
	if err := copyFun(old, newValue.Interface()); err != nil {
		return nil, err
	} else {
		return newValue.Elem().Interface(), nil
	}
}

//-----------------------------------------------------------------------------------------------------------//
