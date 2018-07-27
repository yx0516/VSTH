package db

import (
	"gpc/db/mongodb/xorm"
	"gpc/intf/log"
)

var (
	obj = struct { // 内部 对象
		pack PackString // 包装 内容
		chk  ChkField   // 检查 字段
		doc  CollDoc    // 操作 文档
		err  *ErrType   // 错误 对象
	}{
		err: NewErrType(xorm.Log),
	}
)

//-----------------------------------------------------------------------------------------------------------//

// 错误 类型
type ErrType struct {
	xorm.ErrType
}

func NewErrType(log log.BeegoLog) *ErrType {
	return &ErrType{
		ErrType: *xorm.NewErrType(log),
	}
}

// 包装 内容
type PackString struct {
	xorm.PackString
	SharePackString
}

// 检查 字段
type ChkField struct {
	xorm.ChkField
	ShareChkField
}

// 操作 文档
type CollDoc struct {
	ShareCollDoc
}

//-----------------------------------------------------------------------------------------------------------//
