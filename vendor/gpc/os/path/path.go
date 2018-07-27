package path

import (
	"path/filepath"
	"strings"

	"github.com/kardianos/osext"
)

var (
	exe_dir string = "" // 程序 所在目录
	exe     string = "" // 程序 全路径
)

func init() {
	if exe_dir, _ = osext.ExecutableFolder(); exe_dir != "" {
		exe_dir = strings.TrimRight(filepath.ToSlash(exe_dir), "/") + "/"
	}

	if exe, _ = osext.Executable(); exe != "" {
		exe = filepath.ToSlash(exe)
	}
}

//-----------------------------------------------------------------------------------------------------------//

func ExeDir() string {
	return exe_dir
}

func Exe() string {
	return exe
}

//-----------------------------------------------------------------------------------------------------------//

// 加上当前程序所在目录
func JoinExeDir(p string) string {
	return filepath.ToSlash(exe_dir + strings.TrimLeft(filepath.ToSlash(p), "/"))
}

//　如果 给定的路径不是 绝对路径 则 加上程序当前目录的路径
func NoAbsJoinExeDir(p string) string {
	if !filepath.IsAbs(p) {
		return JoinExeDir(p)
	} else {
		return filepath.ToSlash(p)
	}
}

//-----------------------------------------------------------------------------------------------------------//

// 删除 ../ 返回
func Trim(s string) string {
	s = filepath.ToSlash(filepath.Clean(strings.TrimSpace(s)))
	return strings.TrimSpace(strings.Replace(s, "../", "", -1))
}

//-----------------------------------------------------------------------------------------------------------//
