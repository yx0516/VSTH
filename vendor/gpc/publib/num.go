package publib

import (
	"fmt"
	"reflect"
	"strconv"
)

// 字符串转整型，err != nil 错误 i = 0 【str = "1.1" 包含小数点转换会错误（不能包含小数点）】
func StrToInt(str string) (int, error) {
	return strconv.Atoi(str)
}

// 字符串 转 uint 【用 uint64 去匹配，因为 int | uint 是跟平台相关的类型】
func StrToUInt(str string) (uint, error) {
	sInt, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err
	} else {
		return uint(sInt), nil
	}
}

// 字符串 转 int64
func StrToInt64(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}

// 字符串 转 uint64 【可能溢出不正确】
func StrToUInt64(str string) (uint64, error) {
	return strconv.ParseUint(str, 10, 64)
}

// 整型转字符串，注意传入的参数是 int
func IntToStr(i int) string {
	return strconv.Itoa(i)
}

// uint 转 string
func UIntToStr(i uint) string {
	return strconv.FormatUint(uint64(i), 10)
}

// []int 转 []string
func IntLToStrL(iList []int) (sList []string) {
	for _, i := range iList {
		sList = append(sList, strconv.Itoa(i))
	}
	return
}

// []uint 转 []string
func UIntLToStrL(iList []uint) (sList []string) {
	for _, i := range iList {
		sList = append(sList, strconv.FormatUint(uint64(i), 10))
	}
	return
}

// uint64 转 string
func UInt64ToStr(i uint64) string {
	return strconv.FormatUint(i, 10)
}

// int64 转 string
func Int64ToStr(i int64) string {
	return strconv.FormatInt(i, 10)
}

// []int64 转 []string
func Int64LToStrL(iList []int64) (sList []string) {
	for _, i := range iList {
		sList = append(sList, strconv.FormatInt(i, 10))
	}
	return
}

// []uint64 转 []string
func UInt64LToStrL(iList []uint64) (sList []string) {
	for _, i := range iList {
		sList = append(sList, strconv.FormatUint(i, 10))
	}
	return
}

// 浮点型转字符串
func FloatToStr(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

// []float64 转 []string
func Float64LToStrL(fList []float64) (sList []string) {
	for _, f := range fList {
		sList = append(sList, strconv.FormatFloat(f, 'f', -1, 64))
	}
	return
}

// 字符串转浮点型，err != nil 错误 i = 0
func StrToFloat(str string) (float64, error) {
	return strconv.ParseFloat(str, 64)
}

// Float6 保留几位小数点
func Float64Sprintf(f float64, pointNum uint) (fp float64) {
	fp, _ = strconv.ParseFloat(fmt.Sprintf("%."+strconv.Itoa(int(pointNum))+"f", f), 64)
	return
}

// interface to int64
func ToInt64(value interface{}) (d int64, err error) {
	val := reflect.ValueOf(value)
	switch value.(type) {
	case int, int8, int16, int32, int64:
		d = val.Int()
	case uint, uint8, uint16, uint32, uint64:
		d = int64(val.Uint())
	default:
		err = fmt.Errorf("ToInt64 need numeric not `%T`", value)
	}
	return
}
