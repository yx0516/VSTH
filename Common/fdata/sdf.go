package fdata

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

//-----------------------------------------------------------------------------------------------------------//

// SDF 格式数据 mol + attribute
type SDF struct {
	MOL
	Attrs Attributes
}

func NewSDF(name string) *SDF {
	sdf := &SDF{}
	sdf.Name = name
	return sdf
}

func (self *SDF) AddNewAttr(name, value string) *SDF {
	self.Attrs = append(self.Attrs, NewAttribute(name, value))
	return self
}

func (self *SDF) AddAttr(attrs ...*Attribute) *SDF {
	self.Attrs = append(self.Attrs, attrs...)
	return self
}

// 转成 sdf 字符串格式
func (self *SDF) ToString() (string, error) {
	buf := new(bytes.Buffer)
	if err := self.WriteTo(buf); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (self *SDF) WriteTo(w io.Writer) (err error) {
	_, err = fmt.Fprintf(w, "%s\n%s\n\n%s\n",
		strings.TrimSpace(self.Name),
		strings.TrimSpace(self.Comment),
		strings.TrimRight(self.Data, "\n"),
	)
	if err != nil {
		return
	}
	return self.Attrs.WriteTo(w)
}

//-----------------------------------------------------------------------------------------------------------//

type SDFs []*SDF

// 转成 sdf 字符串格式
func (self *SDFs) ToString() (string, error) {
	buf := new(bytes.Buffer)
	if err := self.WriteTo(buf); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (self SDFs) WriteTo(w io.Writer) (err error) {
	for _, sdf := range self {
		if err = sdf.WriteTo(w); err != nil {
			return
		}
	}
	return
}

//-----------------------------------------------------------------------------------------------------------//
