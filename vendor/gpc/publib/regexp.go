package publib

import (
	"regexp"
)

// 是否 IPv4 的前缀
func IsIPv4Prefix(ip string) (ok bool) {
	ok, _ = regexp.MatchString(`^((2[0-4]\d|25[0-5]|1\d\d|[1-9]\d|\d)\.){3}(2[0-4]\d|25[0-5]|1\d\d|[1-9]\d|\d)$`, ip)
	if !ok {
		ok, _ = regexp.MatchString(`^((2[0-4]\d|25[0-5]|1\d\d|[1-9]\d|\d)\.){0,3}$`, ip)
	}
	return
}

// 只能是 字母 和 数字【并且不能是空】
func IsCharNum(s string) (ok bool) {
	ok, _ = regexp.MatchString(`^[a-zA-Z0-9]+$`, s)
	return
}

// 只能是 字母开头的，由字母和数字组成【并且不能是空】
func IsCharHeadNum(s string) (ok bool) {
	ok, _ = regexp.MatchString(`^[a-zA-Z]{1}[a-zA-Z0-9]+$`, s)
	return
}

// 只能是 字母【并且不能是空】
func IsChar(s string) (ok bool) {
	ok, _ = regexp.MatchString(`^[a-zA-Z]+$`, s)
	return
}

// 只能是全汉字(UTF-8编码)
func IsHan(s string) (ok bool) {
	ok, _ = regexp.MatchString(`^[\p{Han}]+$`, s)
	return
}

// [a-z0-9A-Z] 或 [a-z0-9A-Z][\-_\.][a-z0-9A-Z]...【中间为连接符】如: app | app_sn | app.sn | app-sn.01
func IsCharNumCond(s string) (ok bool) {
	ok, _ = regexp.MatchString(`^[a-z0-9A-Z]+([\-_\.][a-z0-9A-Z]+)*$`, s)
	return
}

// 字母开头的 [a-z0-9A-Z] 或 [a-z0-9A-Z][\-_\.][a-z0-9A-Z]...【中间为连接符】如: app | app_sn | app.sn | app-sn.01
func IsCharHeadNumCond(s string) (ok bool) {
	ok, _ = regexp.MatchString(`^[a-zA-Z]{1}[a-z0-9A-Z]*?([\-_\.][a-z0-9A-Z]+)*$`, s)
	return
}

// 字母开头的 [a-z0-9A-Z] 或 [a-z0-9A-Z][\-_][a-z0-9A-Z]...【中间为连接符】如: app | app_sn
func IsCharHeadNumCondLine(s string) (ok bool) {
	ok, _ = regexp.MatchString(`^[a-zA-Z]{1}[a-z0-9A-Z]*?([\-_][a-z0-9A-Z]+)*$`, s)
	return
}

// 必须包含【数字】&【小写字母】&【大写字母】&【特殊符号】
func IsDigitAlphaPunct(s string) (ok bool) {
	// 必须包含数字，相当于 [0-9]
	ok, _ = regexp.MatchString(`[[:digit:]]+`, s)
	if !ok {
		return
	}
	// 必须包含小写字母，相当于 [a-z]
	ok, _ = regexp.MatchString(`[[:lower:]]+`, s)
	if !ok {
		return
	}
	// 必须包含大写字母，相当于 [A-Z]
	ok, _ = regexp.MatchString(`[[:upper:]]+`, s)
	if !ok {
		return
	}
	// 所有的特殊符号，相当于 [!-/:-@[-`{-~]
	ok, _ = regexp.MatchString(`[[:punct:]]+`, s)
	return
}

// 电子邮箱的正则匹配, 考虑到各个网站的 mail 要求不一样, 这里匹配比较宽松
// 邮箱用户名可以包含 0-9, A-Z, a-z, -, _, .
// 开头字母不能是 -, _, .
// 结尾字母不能是 -, _, .
// -, _, . 这三个连接字母任意两个不能连续, 如不能出现 --, __, .., -_, -., _.
// 邮箱的域名可以包含 0-9, A-Z, a-z, -
// 连接字符 - 只能出现在中间, 不能连续, 如不能 --
// 支持多级域名, x@y.z, x@y.z.w, x@x.y.z.w.e
func IsMail(s string) (ok bool) {
	mailPattern := `^[a-z0-9A-Z]+([\-_\.][a-z0-9A-Z]+)*@([a-z0-9A-Z]+(-[a-z0-9A-Z]+)*\.)+[a-zA-Z]+$`
	ok, _ = regexp.MatchString(mailPattern, s)
	return
}

// 验证 手机 号码 1**********
func IsMobile(s string) (ok bool) {
	ok, _ = regexp.MatchString(`^1\d{10}$`, s)
	return
}
