package publib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

//-----------------------------------------------------------------------------------------------------------//

func JsonMarshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func JsonMarshalString(v interface{}) (string, error) {
	if bytes, err := json.Marshal(v); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}

// 转成 json 格式 字符串
func JsonMarshalToString(v interface{}) string {
	if bytes, err := json.Marshal(v); err != nil {
		return ""
	} else {
		return string(bytes)
	}
}

// 转成 json 格式：美化输出
func JsonMarshalPretty(v interface{}) ([]byte, error) {
	return json.MarshalIndent(v, "", "    ")
}

// 转成 json 格式：美化输出
func JsonMarshalPrettyString(v interface{}) (string, error) {
	bytes, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}

// 转成 json 格式：字节数组 美化输出
func JsonMarshalPrettyToByte(v interface{}) []byte {
	if bytes, err := json.MarshalIndent(v, "", "    "); err != nil {
		return nil
	} else {
		return bytes
	}
}

// 转成 json 格式：字符串 美化输出
func JsonMarshalPrettyToString(v interface{}) string {
	if bytes, err := json.MarshalIndent(v, "", "    "); err != nil {
		return ""
	} else {
		return string(bytes)
	}
}

// 把 json 格式的 字符串 转成 struct
func JsonUnmarshalUseNumberString(jsonString string, v interface{}) error {
	d := json.NewDecoder(strings.NewReader(jsonString))
	d.UseNumber()
	return d.Decode(&v)
}

// 把 json 格式的 []byte 转成 struct
func JsonUnmarshalUseNumber(jsonByte []byte, v interface{}) error {
	d := json.NewDecoder(bytes.NewReader(jsonByte))
	d.UseNumber()
	return d.Decode(&v)
}

// 把 json 格式的 []byte 转成 struct
func JsonUnmarshal(jsonByte []byte, v interface{}) error {
	return json.Unmarshal(jsonByte, v)
}

// 接口 通过 json 解析
func JsonIntfToStruct(src, dest interface{}) error {
	if jsonByte, err := json.Marshal(src); err == nil {
		return json.Unmarshal(jsonByte, dest)
	} else {
		return err
	}
}

// 接口 通过 json 解析到 struct
func JsonIntfUseNumberToStruct(src, dest interface{}) error {
	if jsonByte, err := json.Marshal(src); err == nil {
		return JsonUnmarshalUseNumber(jsonByte, dest)
	} else {
		return err
	}
}

//-----------------------------------------------------------------------------------------------------------//

// 打印输出
func Dump(v interface{}) {
	if bytes, err := json.MarshalIndent(v, "", "    "); err == nil {
		fmt.Println(string(bytes))
	} else {
		fmt.Println("Dump error:" + err.Error())
	}
}

// json 化后存入文件
func JsonMarshalPrettyToJsonFile(v interface{}, filePath string) error {
	if bytes, err := JsonMarshalPretty(v); err != nil {
		return err
	} else {
		if _, err = OutByteToFile(filePath, bytes); err != nil {
			return err
		}
	}
	return nil
}

// 读取 json 文件后，反序列化
func ReadJsonFileUnmarshalUseNumber(filePath string, v interface{}) error {
	if bytes, err := ioutil.ReadFile(filePath); err != nil {
		return err
	} else {
		return JsonUnmarshalUseNumber(bytes, v)
	}
}

//-----------------------------------------------------------------------------------------------------------//
