// 时间相关
package publib

import (
	"fmt"
	"os"
	"time"
)

const (
	// 默认格式化为 Y-M-D h:m:s 的模板,死变态老把时间模板设置到这个时间点格式
	DEF_TIME_FORMAT     = "2006-01-02 15:04:05"
	TIME_FORMAT_YMD     = "20060102"   // 时间格式：20140725
	TIME_FORMAT_YYYMMDD = "2006-01-02" // 时间格式：2014-07-25
)

// 把时间 time.Time 格式化【格式错误返回的信息也是不对版的】
func TimeFormat(t time.Time, fmt string) string {
	// Time类型有这几个函数可以获取对于的值：t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second()
	// 另外 t.Format() 函数可以按照模板格式化
	// 源码：go/src/pkg/time/format.go 有定义其它的格式化模板
	return t.Format(fmt)
}

// 把时间 time.Time 格式化为:Y-M-D h:m:s
func TimeFormatDef(t time.Time) string {
	return t.Format(DEF_TIME_FORMAT)
}

/*
   把 Y-M-D h:m:s 格式的字符串时间转换为 time.Time 类型
   如果 err != nil 则转换错误， 返回值的 t 值是不对的
   所以需要自己检查 返回的 err 值
*/
func StrTimeToTime(ft string) (time.Time, error) {
	// 指定时区为本地时区的时间
	return time.ParseInLocation(DEF_TIME_FORMAT, ft, time.Local)
}

// 把时间戳 转换为 time.Time 类型
func TimestampToTime(u int64) time.Time {
	return time.Unix(u, 0)
}

// 把时间戳纳秒 转换为 time.Time 类型
func TimestampNToTime(nsec int64) time.Time {
	return time.Unix(0, nsec)
}

// 把 Y-M-D h:m:s 格式的字符串时间转换为 时间戳
func StrTimeToTimestamp(ft string) (u int64, err error) {
	var t time.Time
	t, err = time.ParseInLocation(DEF_TIME_FORMAT, ft, time.Local)
	if err != nil {
		u = 0
	} else {
		u = t.Unix()
	}
	return
}

// 把时间戳 转换为字符串格式的时间 Y-M-D h:m:s
func TimestampToStrTime(u int64) string {
	return TimeFormatDef(time.Unix(u, 0))
}

// 获取当前时间格式化为:Y-M-D h:m:s
func NowTime() string {
	// time.Now() 获取当前时间,返回Time类型
	return TimeFormatDef(time.Now())
}

// 获取当前时间格式化(路径连接的年月日)为: Linux: Y/M/D | Windows: Y\M\D
func NowTimeFmtPathSeparator() string {
	sep := string(os.PathSeparator)
	return time.Now().Format("2006" + sep + "01" + sep + "02")
}

// 获取当前时间格式化(路径连接的年月日)为:Y/M/D【不区分平台】
func NowTimeFmtPath() string {
	return time.Now().Format("2006/01/02")
}

// 获取当前时间的时间戳
func NowTimestamp() int64 {
	return time.Now().Unix()
}

// 阻塞休眠,单位：毫秒【传递负值不休眠】
func TimeSleep(interval int64) {
	time.Sleep(time.Millisecond * time.Duration(interval))
}

// 秒时间格式化为可视字符串显示【110 Days 10:20:30】
func SecTimeHuman(sec uint64) (human string) {
	var year, day, hour, minute, second uint64
	if sec >= 31536000 {
		year = sec / 31536000
		sec = sec % 31536000
		human = fmt.Sprintf("%v years ", year)
	}

	if sec >= 86400 {
		day = sec / 86400
		sec = sec % 86400
		human += fmt.Sprintf("%v days ", day)
	}
	if sec >= 3600 {
		hour = sec / 3600
		sec = sec % 3600
	}
	if sec >= 60 {
		minute = sec / 60
		second = sec % 60
	}
	if sec < 60 {
		second = sec
	}
	human += fmt.Sprintf("%02d:%02d:%02d", hour, minute, second)
	return
}

//-----------------------------------------------------------------------------------------------------------//

// 等待 超时，单位：毫秒
func WaitTimeout(timeout int64, fun func()) error {
	if timeout > 0 {
		chDone := make(chan interface{}, 1) // 等待完成 或者 超时的信道

		go func() {
			fun()
			chDone <- true
		}()

		select {
		case <-time.After(time.Duration(timeout) * time.Millisecond): // 等待指定 毫秒
			return fmt.Errorf("wait timeout(%v ms).", timeout)
		case <-chDone:
		}
	} else {
		fun()
	}

	return nil
}

// 等待 超时，单位：毫秒[ <=0 表示不超时]
func WaitTimeoutError(timeout int64, fun func() error) error {
	if timeout > 0 {
		chDone := make(chan error, 1)

		go func() {
			chDone <- fun()
		}()

		select {
		case <-time.After(time.Duration(timeout) * time.Millisecond):
			return fmt.Errorf("wait timeout(%v ms).", timeout)
		case doErr := <-chDone:
			return doErr
		}
	} else {
		return fun()
	}
}

//-----------------------------------------------------------------------------------------------------------//
