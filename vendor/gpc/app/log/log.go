package log

import (
	"strings"

	"github.com/astaxie/beego/logs"

	"gpc/app/cst"
	"gpc/os/path"
	"gpc/web/beego/log"

	plib "gpc/publib"
)

//-----------------------------------------------------------------------------------------------------------//

// 扩展 beego 的日志
type Log struct {
	*logs.BeeLogger
}

// 错误，则退出【errMsg + err.Error()】
func (self *Log) ErrExit(err error, errMsg ...string) {
	if err != nil {
		msg := err.Error()
		if len(errMsg) > 0 {
			msg = errMsg[0] + msg
		}
		self.Emergency(msg)
		plib.TimeSleep(1 * 1000)
		plib.OsExit(1, msg)
	}
}

// 输出到控制台
func (self *Log) EnableConsole() error {
	return self.SetLogger("console", "")
}

//-----------------------------------------------------------------------------------------------------------//

// 创建 日志【错误会终止程序】
func NewLog(fileName string) *Log {
	return &Log{NewBeeLog(fileName)}
}

// 创建 BeeLogger 日志【错误会终止程序】
func NewBeeLog(fileName string) *logs.BeeLogger {
	if fileName = strings.TrimSpace(fileName); fileName == "" {
		fileName = "app.log"
	}
	return log.CreateBeeLoggerByFile(path.NoAbsJoinExeDir(cst.CST_DIR_LOG + "/" + fileName))
}

//-----------------------------------------------------------------------------------------------------------//
