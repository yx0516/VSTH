package openbabel

/*
#include "convert.h"
#cgo LDFLAGS: -lopenbabel
*/
import "C"

import (
	"encoding/json"
	"unsafe"
)

//-----------------------------------------------------------------------------------------------------------//

// 状态
type State struct {
	Status int
	Msg    string
	Data   string
}

// Status == 0 为 True
func (self *State) OK() bool {
	return self.Status == 0
}

// 格式化 JSON 字符串
func (self *State) StringJson() string {
	if bytes, err := json.MarshalIndent(self, "", "    "); err != nil {
		return ""
	} else {
		return string(bytes)
	}
}

func NewState() *State {
	return &State{}
}

//-----------------------------------------------------------------------------------------------------------//

// mol 格式的分子 转换 2D -> 3D【命令 obabel 1.mol -O 2.mol --gen3D】
func ConvertMol(strMol string) *State {
	inMol := C.CString(strMol)
	defer C.free(unsafe.Pointer(inMol)) // #include <stdlib.h>

	ret := &C.struct_state{}
	C.ConvertMol(inMol, ret)

	return &State{
		Status: int(ret.status),
		Msg:    C.GoString(ret.msg),
		Data:   C.GoString(ret.data),
	}
}

//-----------------------------------------------------------------------------------------------------------//
