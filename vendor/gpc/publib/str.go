package publib

import (
	"strings"
	"unicode"
)

//-----------------------------------------------------------------------------------------------------------//

// 填充
func StrFilling(src, tag string, count int) string {
	if len(src) > count {
		return string(src[0:count])
	} else {
		src += strings.Repeat(tag, count-len(src))
		return src
	}
}

// []string 转 []int64
func StrLToInt64List(sList []string) (iList []int64, err error) {
	var i int64
	for _, s := range sList {
		i, err = StrToInt64(s)
		if err != nil {
			return
		}
		iList = append(iList, i)
	}
	return
}

// []inteface{string} 转 []int64
func StrItfLToInt64List(sList []interface{}) (iList []int64, err error) {
	defer func() {
		if rE := recover(); rE != nil {
			err = rE.(error)
		}
	}()
	var i int64
	for _, s := range sList {
		i, err = StrToInt64(s.(string))
		if err != nil {
			return
		}
		iList = append(iList, i)
	}
	return
}

// 字符串前缀是否为大写开头
func CharPrefixIsUpper(s string) bool {
	for _, c := range s {
		if unicode.IsUpper(c) {
			return true
		} else {
			return false
		}
	}
	return false
}

// 把字符串倒序
func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// 驼峰大小写字符串转下划线连接格式【 XxYy to xx_yy】
func SnakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:len(data)]))
}

// 下划线连接字符串转驼峰大小写格式【xx_yy to XxYy】
func CamelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:len(data)])
}

//-----------------------------------------------------------------------------------------------------------//
