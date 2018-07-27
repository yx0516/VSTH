package log

import (
	"fmt"
	"os"
	"strings"

	"github.com/astaxie/beego/logs"

	plib "gpc/publib"
)

const (
	CST_DEF_LOG_LINES int64 = 10000 // 默认 日志 分割 行数
)

//-----------------------------------------------------------------------------------------------------------//

// 日志
type BeeLogger struct {
	FileName string
	Lines    int64
	Console  bool
	ErrExit  bool
}

func NewBeeLogger(fileName string, console, errExit bool) *BeeLogger {
	return &BeeLogger{
		FileName: fileName,
		Lines:    CST_DEF_LOG_LINES,
		Console:  console,
		ErrExit:  errExit,
	}
}

func (self *BeeLogger) InitArgs() *BeeLogger {
	if self.FileName = strings.TrimSpace(self.FileName); self.FileName == "" {
		self.FileName = "logs/app.log"
	}
	if self.Lines <= 0 {
		self.Lines = CST_DEF_LOG_LINES
	}
	return self
}

func (self *BeeLogger) New() (*logs.BeeLogger, error) {
	return self.Make(self.FileName)
}

func (self *BeeLogger) Make(logFilePath string) (log *logs.BeeLogger, err error) {
	self.InitArgs()

	if logFilePath = strings.TrimSpace(logFilePath); logFilePath == "" {
		logFilePath = self.FileName
	}

	defer func() {
		if err != nil && self.ErrExit {
			os.Stderr.WriteString(err.Error() + "\n")
			os.Exit(1)
		}
	}()

	if err = plib.CreateParentPath(logFilePath); err != nil {
		return nil, fmt.Errorf("create log file(%s) error:%s", logFilePath, err.Error())
	}

	log = logs.NewLogger(self.Lines)

	if err := log.SetLogger("file", `{"filename":"`+logFilePath+`"}`); err != nil {
		return nil, fmt.Errorf("init log error:" + err.Error())
	}

	if self.Console {
		log.SetLogger("console", "")
	} else {
		log.DelLogger("console")
	}
	return log, nil
}

//-----------------------------------------------------------------------------------------------------------//

// 初始化日志对象返回
func CreateBeeLogger(logFilePath string, isConsole, isErrExit bool) (log *logs.BeeLogger, err error) {
	return NewBeeLogger(logFilePath, isConsole, isErrExit).New()
}

// 创建 默认 日志【错误会终止程序】
func CreateDefBeeLogger() (log *logs.BeeLogger) {
	log, _ = NewBeeLogger("logs/app.log", false, true).New()
	return
}

// 创建 默认 日志【错误会终止程序】
func CreateBeeLoggerByFile(filePath string) (log *logs.BeeLogger) {
	log, _ = NewBeeLogger(filePath, false, true).New()
	return
}

// 创建 输出到 控制台的 日志对象
func CreateConsoleBeeLogger() (log *logs.BeeLogger) {
	log = logs.NewLogger(10000)
	log.SetLogger("console", "")
	return
}

//-----------------------------------------------------------------------------------------------------------//
