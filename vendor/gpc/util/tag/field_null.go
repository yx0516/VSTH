package tag

import (
	"gpc/util/ref"
)

//-----------------------------------------------------------------------------------------------------------//

// 字段是否为 空字段
func (self *FieldCheck) IsNull(sf *ref.StructField, tags ...string) (ret *RetStatus) {
	if len(tags) == 0 {
		tags = append(tags, self.TagNull)
	}
	return self.Check(sf, tags...)
}

// 字段是否为 非空字段
func (self *FieldCheck) IsNotNull(sf *ref.StructField, tags ...string) (ret *RetStatus) {
	if len(tags) == 0 {
		tags = append(tags, self.TagNotNull)
	}
	return self.Check(sf, tags...)
}

//-----------------------------------------------------------------------------------------------------------//
