package publib

import (
	"fmt"
	"os"
)

//-----------------------------------------------------------------------------------------------------------//

// 实现类似 try...catch..的异常处理机制
// Try(func() {
//	 panic("this is panic")
// }, func(e interface{}) {
//	 fmt.Println(e)
// })
func Try(fun func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
	}()
	fun()
}

// 结束程序退出【defer 函数都不会被执行到】
func OsExit(code int, errMsg string) {
	os.Stderr.WriteString(errMsg + "\n")
	os.Exit(code)
}

// 结束程序退出【defer 函数都不会被执行到】
func OsExitPrint(code int, msg ...interface{}) {
	fmt.Println(msg...)
	os.Exit(code)
}

// 执行函数 错误，则 终止程序
func FuncErrExit(fun func() error) {
	if err := fun(); err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
}

//-----------------------------------------------------------------------------------------------------------//
