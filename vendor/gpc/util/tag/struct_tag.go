package tag

import (
	"errors"
	"reflect"

	"gpc/util/ref"
)

//-----------------------------------------------------------------------------------------------------------//

// 结构体 注解
type StructTag struct {
	Struct  *ref.Struct
	Fields  map[string]*ref.StructField
	convert ref.NameConvert
	Fc      *FieldCheck
	tags    []string
}

func NewStructTag(convert ref.NameConvert, tags ...string) *StructTag {
	st := &StructTag{
		Fields:  make(map[string]*ref.StructField),
		convert: convert,
		Fc:      NewFieldCheck(),
	}
	st.tags = append(tags, st.Fc.GetTags()...)
	return st
}

//-----------------------------------------------------------------------------------------------------------//

// 反射获取 结构体信息
func (self *StructTag) Reflect(v interface{}) (err error) {
	if self.Struct, err = ref.RefStructInfo(v, self.tags...); err == nil {
		for _, sf := range self.Struct.Fields {
			self.Fc.UpdateStructFieldVal(sf)
		}
		self.Convert()
	}
	return
}

func (self *StructTag) Convert() {
	if self.convert != nil {
		self.Fields = self.convert.Convert(self.Struct.Fields)
	} else {
		self.Fields = self.Struct.Fields
	}
}

func (self *StructTag) Update(v interface{}) (err error) {
	var name string
	if name, err = ref.RefStructTypeFullName(v); err != nil {
		return
	}
	if self.Struct == nil || self.Fields == nil || self.Struct.FullName != name {
		return self.Reflect(v)
	} else {
		value := reflect.Indirect(reflect.ValueOf(v))
		for field, sf := range self.Struct.Fields {
			if val := value.FieldByName(field); val.IsValid() {
				sf.Value = val
				self.Fc.UpdateStructFieldVal(sf)
			}
		}
	}
	return
}

//-----------------------------------------------------------------------------------------------------------//

func (self *StructTag) ExistField(field string) (ok bool) {
	if len(self.Fields) > 0 {
		_, ok = self.Fields[field]
	}
	return
}

func (self *StructTag) GetFieldInfo(field string) *ref.StructField {
	if len(self.Fields) > 0 {
		if sf, ok := self.Fields[field]; ok {
			return sf
		} else {
			for _, sf = range self.Fields {
				if sf.Name == field {
					return sf
				}
			}
		}
	}
	return nil
}

func (self *StructTag) FindConvertName(field string) (string, error) {
	if len(self.Fields) > 0 {
		if sf, ok := self.Fields[field]; ok {
			return field, nil
		} else {
			var name string
			for name, sf = range self.Fields {
				if sf.Name == field {
					return name, nil
				}
			}
		}
	}
	return "", errors.New("convert name not found.")
}

func (self *StructTag) GetFieldNames() (fields []string) {
	for field, _ := range self.Fields {
		fields = append(fields, field)
	}
	return
}

func (self *StructTag) GetStrucFieldNames() (fields []string) {
	if self.Struct != nil {
		for _, sf := range self.Fields {
			fields = append(fields, sf.Name)
		}
	}
	return
}

func (self *StructTag) GetFieldValues(fields ...string) (fv map[string]interface{}) {
	fv = make(map[string]interface{})
	if len(fields) > 0 && len(self.Fields) > 0 {
		for _, field := range fields {
			if sf := self.GetFieldInfo(field); sf != nil && sf.Value.CanInterface() {
				fv[field] = sf.Value.Interface()
			}
		}
	} else {
		for field, sf := range self.Fields {
			if sf.Value.CanInterface() {
				fv[field] = sf.Value.Interface()
			}
		}
	}
	return
}

//-----------------------------------------------------------------------------------------------------------//

func (self *StructTag) GetExistTagFieldNames(tags ...string) (fields []string) {
	for field, sf := range self.Fields {
		for _, tag := range tags {
			if _, ok := sf.Tags[tag]; ok {
				fields = append(fields, field)
				break
			}
		}
	}
	return
}

func (self *StructTag) GetNotExistTagFieldNames(tags ...string) (fields []string) {
	var isExist bool
	for field, sf := range self.Fields {
		isExist = false
		for _, tag := range tags {
			if _, ok := sf.Tags[tag]; ok {
				isExist = true
				break
			}
		}
		if !isExist {
			fields = append(fields, field)
		}
	}
	return
}

//-----------------------------------------------------------------------------------------------------------//

func (self *StructTag) NullFields(tags ...string) (fields []string, err error) {
	return self.nullFields(true, tags...)
}

func (self *StructTag) NotNullFields(tags ...string) (fields []string, err error) {
	return self.nullFields(false, tags...)
}

func (self *StructTag) nullFields(isNull bool, tags ...string) (fields []string, err error) {
	var fv map[string]interface{}
	if fv, err = self.fieldValues(isNull, tags...); err != nil {
		return
	}
	for field, _ := range fv {
		fields = append(fields, field)
	}
	return
}

func (self *StructTag) NullFieldValues(tags ...string) (fieldValues map[string]interface{}, err error) {
	return self.fieldValues(true, tags...)
}

func (self *StructTag) NotNullFieldValues(tags ...string) (fieldValues map[string]interface{}, err error) {
	return self.fieldValues(false, tags...)
}

func (self *StructTag) fieldValues(isNull bool, tags ...string) (fieldValues map[string]interface{}, err error) {
	ret := NewRetStatus()
	fieldValues = make(map[string]interface{})
	defer func() {
		if err != nil {
			fieldValues = make(map[string]interface{})
		}
	}()
	var fun func(sf *ref.StructField, tags ...string) (ret *RetStatus)
	if isNull {
		fun = self.Fc.IsNull
		if len(tags) == 0 {
			tags = append(tags, self.Fc.TagNull)
		}
	} else {
		fun = self.Fc.IsNotNull
		if len(tags) == 0 {
			tags = append(tags, self.Fc.TagNotNull)
		}
	}
	for field, sf := range self.Fields {
		realTag := sf.FilterTags(tags...)
		if realTag.Count() <= 0 {
			continue
		}
		if ret = fun(sf, realTag...); ret.IsOK() && sf.Value.CanInterface() {
			fieldValues[field] = sf.Value.Interface()
		} else if ret.IsErr() {
			err = errors.New(ret.Msg)
			return
		}
	}
	return
}

//-----------------------------------------------------------------------------------------------------------//

func (self *StructTag) Check(fields ...string) (fieldValues map[string]interface{}, err error) {
	fieldValues = make(map[string]interface{})
	defer func() {
		if err != nil {
			fieldValues = make(map[string]interface{})
		}
	}()
	if len(fields) > 0 && len(self.Fields) > 0 {
		for _, field := range fields {
			if sf := self.GetFieldInfo(field); sf != nil {
				if sf.ExistTag(self.Fc.TagChk) {
					if ret := self.Fc.Check(sf); !ret.IsOK() {
						err = errors.New(ret.Msg)
						return
					}
					fieldValues[field] = sf.Value.Interface()
				}
			} else {
				err = errors.New(field + " field does not exist.")
				return
			}
		}
	} else {
		for field, sf := range self.Fields {
			if sf.ExistTag(self.Fc.TagChk) {
				if ret := self.Fc.Check(sf); !ret.IsOK() {
					err = errors.New(ret.Msg)
					return
				}
				fieldValues[field] = sf.Value.Interface()
			}
		}
	}
	return
}

func (self *StructTag) SetNullDefValue(fields ...string) error {
	fun := func(sf *ref.StructField) error {
		if _, ok := sf.Tags[self.Fc.TagDef]; ok {
			if ret := self.Fc.SetNullDefValue(sf); ret.IsErr() {
				return errors.New(ret.Msg)
			}
		}
		return nil
	}

	if len(fields) > 0 && len(self.Fields) > 0 {
		for _, field := range fields {
			if sf := self.GetFieldInfo(field); sf != nil {
				if err := fun(sf); err != nil {
					return err
				}
			}
		}
	} else {
		for _, sf := range self.Fields {
			if err := fun(sf); err != nil {
				return err
			}
		}
	}

	return nil
}

//-----------------------------------------------------------------------------------------------------------//

func (self *StructTag) StringTrim(fields ...string) (err error) {
	if len(fields) > 0 && len(self.Fields) > 0 {
		for _, field := range fields {
			if sf := self.GetFieldInfo(field); sf != nil {
				if err = self.Fc.StringTrim(sf); err != nil {
					return
				}
			}
		}
	} else {
		for _, sf := range self.Fields {
			if err = self.Fc.StringTrim(sf); err != nil {
				return
			}
		}
	}
	return
}

//-----------------------------------------------------------------------------------------------------------//
