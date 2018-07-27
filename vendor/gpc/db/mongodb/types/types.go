package types

import (
	"gpc/db/mongodb/cst"
)

//-----------------------------------------------------------------------------------------------------------//

// 状态
type Status int

func (self *Status) SetOn() {
	*self = Status(cst.Status_On)
}

func (self *Status) SetOff() {
	*self = Status(cst.Status_Off)
}

func (self Status) IsOn() bool {
	return int(self) == cst.Status_On
}

func (self Status) IsOff() bool {
	return int(self) == cst.Status_Off
}

func (self Status) IsNotSet() bool {
	return int(self) == cst.Status_NotSet
}

//-----------------------------------------------------------------------------------------------------------//
