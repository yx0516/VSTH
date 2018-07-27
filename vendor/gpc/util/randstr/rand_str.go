package randstr

import (
	"math/rand"
	"time"
)

var (
	runeDigit = []rune("1234567890")
	lenDigit  = len(runeDigit)

	runeAlpha = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	lenAlpha  = len(runeAlpha)

	runeAlphaDigit = []rune("abcdefghijklmnopqrstuvwxyz1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	lenAlphaDigit  = len(runeAlphaDigit)

	// 删掉一部分特殊符号
	runeAll = []rune(`abcdefghijklmnopqrstuvwxyz1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ~!@#$^&*()_+-={}[]<>`)
	lenAll  = len(runeAll)
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// 生成指定长度(n>=0) 随机数字 组成的字符串[0-9]
func RandDigit(n int) string {
	buffer := make([]rune, n)
	for i := range buffer {
		buffer[i] = runeDigit[rand.Intn(lenDigit)]
	}
	return string(buffer)
}

// 生成指定长度(n>=0) 随机字母 组成的字符串[a-zA-Z]
func RandAlpha(n int) string {
	buffer := make([]rune, n)
	for i := range buffer {
		buffer[i] = runeAlpha[rand.Intn(lenAlpha)]
	}
	return string(buffer)
}

// 生成指定长度(n>=0) 随机字母和数字 组成的字符串
func RandAlphaDigit(n int) string {
	buffer := make([]rune, n)
	for i := range buffer {
		buffer[i] = runeAlphaDigit[rand.Intn(lenAlphaDigit)]
	}
	return string(buffer)
}

// 生成指定长度(n>=0) 随机字母、数字、部分特殊符号 组成的字符串
func RandAll(n int) string {
	buffer := make([]rune, n)
	for i := range buffer {
		buffer[i] = runeAll[rand.Intn(lenAll)]
	}
	return string(buffer)
}

//-----------------------------------------------------------------------------------------------------------//

// 对 [0 - N) 的数字进行随机打散
func RandSliceN(count int) []int {
	arr := make([]int, count)
	for i, _ := range arr {
		arr[i] = i
	}

	ShuffleInt(arr)
	return arr
}

// 打散随机
func ShuffleInt(arr []int) {
	for i := range arr {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
}

// 打散随机
func ShuffleUInt(arr []uint) {
	for i := range arr {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
}

// 打散随机
func ShuffleInt64(arr []int64) {
	for i := range arr {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
}

// 打散随机
func ShuffleUInt64(arr []uint64) {
	for i := range arr {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
}

// 打散随机
func ShuffleFloat64(arr []float64) {
	for i := range arr {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
}

// 打散随机
func ShuffleIntf(arr []interface{}) {
	for i := range arr {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
}

//-----------------------------------------------------------------------------------------------------------//
