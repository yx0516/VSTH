package log

//-----------------------------------------------------------------------------------------------------------//

// 基本 日志 接口
type BaseLog interface {
	Critical(format string, v ...interface{})      // Level = 2
	Error(format string, v ...interface{})         // Level = 3
	Warning(format string, v ...interface{})       // Level = 4
	Informational(format string, v ...interface{}) // Level = 6
	Debug(format string, v ...interface{})         // Level = 7
}

// 日志 接口【beego 日志】
type BeegoLog interface {
	BaseLog
	Emergency(format string, v ...interface{}) // Level = 0
	Alert(format string, v ...interface{})     // Level = 1
	Notice(format string, v ...interface{})    // Level = 5
}

//-----------------------------------------------------------------------------------------------------------//
