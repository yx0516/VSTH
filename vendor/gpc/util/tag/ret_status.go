package tag

import (
	plib "gpc/publib"
)

//-----------------------------------------------------------------------------------------------------------//

// 返回状态
type RetStatus struct {
	Status uint8
	Msg    string
}

func NewRetStatus() *RetStatus {
	return &RetStatus{}
}

func (self *RetStatus) String() string {
	return plib.JsonMarshalToString(self)
}

func (self *RetStatus) StringPretty() string {
	return plib.JsonMarshalPrettyToString(self)
}

//-----------------------------------------------------------------------------------------------------------//

func (self *RetStatus) IsOK() bool {
	return self.Status == CST_STATUS_OK
}

func (self *RetStatus) IsFail() bool {
	return self.Status == CST_STATUS_FAIL
}

func (self *RetStatus) IsErr() bool {
	return self.Status == CST_STATUS_ERR
}

//-----------------------------------------------------------------------------------------------------------//

// 设置值
func (self *RetStatus) Set(status uint8, msg string) *RetStatus {
	self.Status = status
	self.Msg = msg
	return self
}

func (self *RetStatus) SetOk(msg string) *RetStatus {
	self.Status = CST_STATUS_OK
	self.Msg = msg
	return self
}

func (self *RetStatus) SetFail(msg string) *RetStatus {
	self.Status = CST_STATUS_FAIL
	self.Msg = msg
	return self
}

func (self *RetStatus) SetErr(msg string) *RetStatus {
	self.Status = CST_STATUS_ERR
	self.Msg = msg
	return self
}

//-----------------------------------------------------------------------------------------------------------//
