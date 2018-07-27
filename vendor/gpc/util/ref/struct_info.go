package ref

import (
	"reflect"

	plib "gpc/publib"
)

//-----------------------------------------------------------------------------------------------------------//

// 结构体 信息
type Struct struct {
	FullName string
	Name     string
	Fields   map[string]*StructField
	HasPtr   bool
}

//-----------------------------------------------------------------------------------------------------------//

// 结构体 属性 信息
type StructField struct {
	Name  string
	Value reflect.Value
	Tags  map[string]string
}

func NewStructField(name string, value reflect.Value) *StructField {
	return &StructField{
		Name:  name,
		Value: value,
		Tags:  make(map[string]string),
	}
}

//-----------------------------------------------------------------------------------------------------------//

func (self *StructField) FilterTags(filterTags ...string) (tags plib.StringArray) {
	if len(self.Tags) > 0 {
		for _, tag := range filterTags {
			if _, ok := self.Tags[tag]; ok {
				tags.Add(tag)
			}
		}
		tags.RemoveDuplicate()
	}
	return
}

func (self *StructField) GetTags() (tags []string) {
	for k, _ := range self.Tags {
		tags = append(tags, k)
	}
	return
}

func (self *StructField) GetTagBody(tag string) (body string) {
	if len(self.Tags) > 0 {
		body, _ = self.Tags[tag]
	}
	return
}

func (self *StructField) ExistTag(tag string) (ok bool) {
	if len(self.Tags) > 0 {
		_, ok = self.Tags[tag]
	}
	return
}

//-----------------------------------------------------------------------------------------------------------//
