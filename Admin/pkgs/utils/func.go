package utils

import (
	plib "gpc/publib"
)

//-----------------------------------------------------------------------------------------------------------//

// 字母开头 + 数字
func CheckName(name string, lenMin, lenMax int) bool {
	if !plib.IsCharHeadNum(name) {
		return false
	}
	if lenMin > 0 && len(name) < lenMin {
		return false
	}
	if lenMax > 0 && len(name) > lenMax {
		return false
	}
	return true
}

// 字母开头 + 数字，长度[4,100]
func CheckStdName(name string) bool {
	return CheckName(name, 4, 100)
}

//-----------------------------------------------------------------------------------------------------------//
