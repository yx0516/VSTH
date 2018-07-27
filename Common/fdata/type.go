package fdata

import (
	"fmt"
	"io"
	"strings"
)

//-----------------------------------------------------------------------------------------------------------//

type InChi struct {
	Value string
	Key   string
}

//-----------------------------------------------------------------------------------------------------------//

// mol 格式 数据
type MOL struct {
	Name    string // 名称
	Comment string // 备注
	Data    string // 不包含前面3行的头文件，从 V2000 | V3000 开始
}

func NewMOL(name string) *MOL {
	return &MOL{
		Name: name,
	}
}

// 尾部包含 \n
func (self *MOL) ToString() string {
	return fmt.Sprintf("%s\n%s\n\n%s\n",
		strings.TrimSpace(self.Name),
		strings.TrimSpace(self.Comment),
		strings.TrimRight(self.Data, "\n"),
	)
}

//-----------------------------------------------------------------------------------------------------------//

// 属性
type Attribute struct {
	Name  string
	Value string // 包含\n分割多行
}

func NewAttribute(name, value string) *Attribute {
	return &Attribute{
		Name:  name,
		Value: value,
	}
}

//-----------------------------------------------------------------------------------------------------------//

// 属性 数组
type Attributes []*Attribute

func (self Attributes) Len() int { return len(self) }

// 唯一 名称
func (self *Attributes) Unique() *Attributes {
	m := make(map[string]bool)
	var array []*Attribute
	for i := self.Len() - 1; i >= 0; i-- {
		if _, ok := m[(*self)[i].Name]; !ok {
			m[(*self)[i].Name] = true
			array = append(array, (*self)[i])
		}
	}
	*self = array

	return self
}

// 写入属性 并且补上 $$$$\n
func (self Attributes) WriteTo(w io.Writer) (err error) {
	for _, attr := range self {
		_, err = fmt.Fprintf(w, fmt.Sprintf("> <%s>\n%s\n\n",
			strings.TrimSpace(attr.Name),
			strings.TrimRight(attr.Value, "\n"),
		))
		if err != nil {
			return
		}
	}
	_, err = w.Write([]byte("$$$$\n"))
	return
}

//-----------------------------------------------------------------------------------------------------------//
