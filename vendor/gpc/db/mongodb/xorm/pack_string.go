package xorm

import (
	"fmt"
	"strings"

	"gpc/web/cst"

	plib "gpc/publib"
)

//-----------------------------------------------------------------------------------------------------------//

// 包装 内容
type PackString struct{}

// 包装 字符串
func (self *PackString) Pack(msg interface{}, tags ...interface{}) string {
	message := fmt.Sprintf("%v", msg)
	if count := len(tags); count > 0 {
		arr := make([]string, count)
		for i := 0; i < count; i++ {
			arr[i] = fmt.Sprintf("%v", tags[i])
		}
		return message + "(" + strings.Join(arr, "|") + ")"
	} else {
		return message
	}
}

//-----------------------------------------------------------------------------------------------------------//

func (self *PackString) Id(id interface{}) string {
	return fmt.Sprintf("%s=%v", cst.FIELD_ID, id)
}

func (self *PackString) Pid(pid uint) string {
	return fmt.Sprintf("%s=%v", cst.FIELD_PID, pid)
}

func (self *PackString) Name(name string) string {
	return cst.FIELD_NAME + "=" + name
}

func (self *PackString) Map(m map[string]interface{}) string {
	return plib.JsonMarshalToString(m)
}

//-----------------------------------------------------------------------------------------------------------//

func (self *PackString) FieldInvalid(field string, tags ...interface{}) string {
	return self.Pack(field+" field invalid.", tags...)
}

func (self *PackString) FieldIsZero(field string) string {
	return field + " field == 0"
}

func (self *PackString) FieldIsNull(field string, tags ...interface{}) string {
	return self.Pack(field+" field null", tags...)
}

func (self *PackString) FieldUnkown(field string) string {
	return field + " field unkown"
}

func (self *PackString) ForeignKeyRef(keyVal, tableName string) string {
	return keyVal + " have foreign key references.(" + tableName + ")"
}

//-----------------------------------------------------------------------------------------------------------//

func (self *PackString) AlreadyExist(name string, tags ...interface{}) string {
	return self.Pack(name+" already exist.", tags...)
}

func (self *PackString) AlreadyExistId(tags ...interface{}) string {
	return self.Pack(cst.FIELD_ID+" already exist.", tags...)
}

func (self *PackString) AlreadyExistName(tags ...interface{}) string {
	return self.Pack(cst.FIELD_NAME+" already exist.", tags...)
}

func (self *PackString) NotExist(name string, tags ...interface{}) string {
	return self.Pack(name+" not exist.", tags...)
}

func (self *PackString) NotExistId(tags ...interface{}) string {
	return self.Pack(cst.FIELD_ID+" not exist.", tags...)
}

func (self *PackString) NotExistName(tags ...interface{}) string {
	return self.Pack(cst.FIELD_NAME+" not exist.", tags...)
}

//-----------------------------------------------------------------------------------------------------------//

func (self *PackString) IdInvalid(id int) string {
	return self.FieldInvalid(cst.FIELD_ID, self.Id(id))
}

func (self *PackString) NameIsNull() string {
	return self.FieldIsNull(cst.FIELD_NAME)
}

//-----------------------------------------------------------------------------------------------------------//
