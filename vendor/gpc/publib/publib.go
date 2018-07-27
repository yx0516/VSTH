package publib

import (
	"fmt"
)

// 调试函数输出信息
func Debug(v ...interface{}) (n int, err error) {
	return fmt.Println(v...)
}
