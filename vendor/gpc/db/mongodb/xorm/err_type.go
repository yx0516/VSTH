package xorm

import (
	"errors"
	"fmt"
	"strings"

	"gopkg.in/mgo.v2"

	"gpc/db/mongodb/cst"
	"gpc/intf/log"
)

//-----------------------------------------------------------------------------------------------------------//

// 断言 CRUD 错误记录日志，并重写 error
type ErrType struct {
	Log log.BeegoLog
}

func NewErrType(log log.BeegoLog) *ErrType {
	return &ErrType{
		Log: log,
	}
}

//-----------------------------------------------------------------------------------------------------------//

// 包装字符串
func (self *ErrType) Pack(msg string, tags ...string) string {
	if len(tags) > 0 {
		return msg + "(" + strings.Join(tags, "|") + ")"
	} else {
		return msg
	}
}

func (self *ErrType) Make(msg string, tags ...string) error {
	return errors.New(self.Pack(msg, tags...))
}

//-----------------------------------------------------------------------------------------------------------//

// 断言 CRUD 数据 错误【重写 error】
func (self *ErrType) assertErrCRUD(msg string, err error, tags ...string) error {
	if err != nil {
		self.Log.Critical(self.Pack(msg+" | "+err.Error(), tags...)) // 错误消息 和 tags 数据 只写到日志里
		return errors.New(msg)
	} else {
		return nil
	}
}

// 断言 插入数据 错误
func (self *ErrType) AssInsert(err error, msg string, tags ...string) error {
	return self.assertErrCRUD("insert error:"+msg, err, tags...)
}

// 断言 读取数据 错误
func (self *ErrType) AssRead(err error, msg string, tags ...string) error {
	if err == nil {
		return nil
	}
	switch err {
	case mgo.ErrNotFound: // 没有找到记录
		return errors.New("query no found:" + msg)
	default:
		return self.assertErrCRUD("read error:"+msg, err, tags...) // 需要记录日志
	}
}

// 断言 更新数据 错误
func (self *ErrType) AssUpdate(err error, msg string, tags ...string) error {
	return self.assertErrCRUD("update error:"+msg, err, tags...)
}

// 断言 删除数据 错误
func (self *ErrType) AssDelete(err error, msg string, tags ...string) error {
	return self.assertErrCRUD("delete error:"+msg, err, tags...)
}

//-----------------------------------------------------------------------------------------------------------//

func (self *ErrType) FieldsIsNull() error {
	return errors.New("empty the contents of the fields set.")
}

func (self *ErrType) StatusOFF(name string) error {
	return fmt.Errorf("%s has been banned.(%s=%v)", name, cst.Field_Status, cst.Status_Off)
}

//-----------------------------------------------------------------------------------------------------------//
