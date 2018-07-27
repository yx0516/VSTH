package xorm

import (
	"errors"
	"strings"

	"gpc/web/cst"
)

//-----------------------------------------------------------------------------------------------------------//

// 检查 字段
type ChkField struct{}

func (self *ChkField) Id(id int) (err error) {
	if id <= 0 {
		err = errors.New(obj.pack.IdInvalid(id))
	}
	return
}

func (self *ChkField) Name(name string) (err error) {
	if name = strings.TrimSpace(name); name == "" {
		err = errors.New(obj.pack.NameIsNull())
	}
	return
}

func (self *ChkField) Status(status int) (err error) {
	if status > 2 || status <= 0 { // 状态值(0预留):1为正常,2为禁用
		return errors.New(obj.pack.FieldInvalid(cst.FIELD_STATUS, "1|2"))
	}
	return
}

//-----------------------------------------------------------------------------------------------------------//
