package publib

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
)

//-----------------------------------------------------------------------------------------------------------//

// 判断文件或目录是否存在
func FileIsExist(filePath string) (ok bool) {
	_, err := os.Stat(filePath)
	return err == nil
}

// 获取 文件的 后缀
func FileExt(filePath string) (string, error) {
	if !FileIsExist(filePath) {
		return "", fmt.Errorf("file not found: %s", filePath)
	} else {
		return path.Ext(filePath), nil
	}
}

// 向文件里写入数据
func WriteIoToFile(filePath string, reader io.Reader, mode int) (n int64, err error) {
	f, err := InitOpenFile(filePath, mode, 0600)
	if err != nil {
		return 0, errors.New("open file error:" + err.Error())
	}
	defer f.Close()
	return io.Copy(f, reader)
}

// 向文件里写入数据
func WriteByteToFile(filePath string, data []byte, mode int) (n int, err error) {
	f, err := InitOpenFile(filePath, mode, 0600)
	if err != nil {
		return 0, errors.New("open file error:" + err.Error())
	}
	defer f.Close()
	return f.Write(data)
}

// 向文件里写入数据
func WriteDataToFile(filePath string, data string, mode int) (n int, err error) {
	f, err := InitOpenFile(filePath, mode, 0600)
	if err != nil {
		return 0, errors.New("open file error:" + err.Error())
	}
	defer f.Close()
	return f.WriteString(data)
}

// 向文件里追加模式写入数据 err != nil 表示写入错误
func AppendDataToFile(filePath string, data string) (n int, err error) {
	return WriteDataToFile(filePath, data, os.O_RDWR|os.O_CREATE|os.O_APPEND)
}

// 向文件里覆盖写入数据
func OutIoToFile(filePath string, reader io.Reader) (n int64, err error) {
	return WriteIoToFile(filePath, reader, os.O_RDWR|os.O_CREATE|os.O_TRUNC)
}

// 向文件里覆盖写入数据
func OutDataToFile(filePath string, data string) (int, error) {
	return WriteDataToFile(filePath, data, os.O_RDWR|os.O_CREATE|os.O_TRUNC)
}

// 向文件里覆盖写入数据
func OutByteToFile(filePath string, data []byte) (n int, err error) {
	return WriteByteToFile(filePath, data, os.O_RDWR|os.O_CREATE|os.O_TRUNC)
}

// 目录不存在就创建目录及子目录
func CreateDirs(dir string) (err error) {
	return os.MkdirAll(dir, 0600)
}

// 创建文件所在的目录
func CreateFileInDirs(filePath string) (err error) {
	return CreateDirs(path.Dir(filePath))
}

// 文件所在目录不存在则自动创建
func InitOpenFile(filePath string, mode int, perm os.FileMode) (f *os.File, err error) {
	err = os.MkdirAll(path.Dir(filePath), perm)
	if err == nil {
		f, err = os.OpenFile(filePath, mode, perm)
	}
	return
}

// 文件所在目录不存在则自动创建
func InitCreateOpenFile(filePath string) (f *os.File, err error) {
	return InitOpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
}

// 文件所在目录不存在则自动创建(追加数据)
func InitAppendOpenFile(filePath string) (f *os.File, err error) {
	return InitOpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)
}

// 获取上级目录路径
func GetParentPath(dirPath string) (par string) {
	par = filepath.ToSlash(dirPath)
	par = strings.TrimRight(par, "/")
	par = path.Dir(par)
	return
}

// 创建上级目录
func CreateParentPath(dirPath string) (err error) {
	return os.MkdirAll(GetParentPath(dirPath), 0600)
}

//-----------------------------------------------------------------------------------------------------------//
