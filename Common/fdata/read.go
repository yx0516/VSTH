package fdata

import (
	"bufio"
	"io"
	"os"
	"regexp"
	"strings"
)

var (
	regAttr = regexp.MustCompile(`> *?<(?P<value>.*?)>`) // 获取属性名称的正则
	/*

		支持:
			 ><Name123>
			 > <Name123>
			 >  <Name123>
			 >   <Name123>
			 >   < Name123 >
			 >   < Name123 > XX111
	*/
)

//-----------------------------------------------------------------------------------------------------------//

func ReadByFile(sdFile string, readCount int) (SDFs, error) {
	if of, err := os.Open(sdFile); err != nil {
		return nil, err
	} else {
		defer of.Close()
		return ReadByIo(of, readCount)
	}
}

func ReadByIo(rd io.Reader, readCount int) (SDFs, error) {
	return ReadByBuf(bufio.NewReader(rd), readCount)
}

// 判断 是否版本
func isVersion(data string) bool {
	data = strings.ToLower(data)
	return data == "v2000" || data == "v3000"
}

// UTF8 无 BOM  ，有 BOM 尾部会多个 \ufeff 特殊字符，无法识别
func ReadByBuf(reader *bufio.Reader, readCount int) ([]*SDF, error) {
	var sdfs SDFs
	firstLine := ""
	for count := 1; ; count++ {
		sdf := new(SDF)

		// 头文件
		if firstLine == "" {
			for {
				if line, err := reader.ReadString('\n'); err != nil { // line 是原始数据，包含 \n
					if err == io.EOF {
						return sdfs, nil
					} else {
						return nil, err
					}
				} else if str := strings.TrimSpace(line); str == "" { // 空白
					continue // 空白行
				} else if fields := strings.Fields(str); len(fields) == 0 {
					continue // 空白行
				} else if data := fields[len(fields)-1]; isVersion(data) { // 判断版本信息
					sdf.Data += line // mol data
					break
				} else { // 前面三行为头文件:第一行为 名称，第二行为 备注， 第三行为 空白行
					sdf.Name = strings.TrimSpace(line) // 有可能无名称 ，但有备注
					for i := 0; i < 2; i++ {
						if line, err = reader.ReadString('\n'); err != nil {
							if err == io.EOF {
								return sdfs, nil
							} else {
								return nil, err
							}
						} else if str := strings.TrimSpace(line); str != "" {
							if fields = strings.Fields(str); len(fields) > 0 {
								if data := fields[len(fields)-1]; isVersion(data) { // 判断版本信息
									sdf.Data += line       // mol data
									sdf.Comment = sdf.Name // 备注
									sdf.Name = ""          // 无名称
									break
								} else if i == 0 { // 备注
									sdf.Comment = str
								}
							}
						}
					}
					break
				}
			}
		} else if fields := strings.Fields(strings.TrimSpace(firstLine)); len(fields) != 0 { // 有数据，则是 读取属性里返回的
			if data := fields[len(fields)-1]; isVersion(data) { // 判断版本信息
				sdf.Data += firstLine // mol data
			} else { // 前面三行为头文件:第一行为 名称，第二行为 备注， 第三行为 空白行
				sdf.Name = strings.TrimSpace(firstLine) // 有可能无名称 ，但有备注
				for i := 0; i < 2; i++ {
					if line, err := reader.ReadString('\n'); err != nil {
						if err == io.EOF {
							return sdfs, nil
						} else {
							return nil, err
						}
					} else if str := strings.TrimSpace(line); str != "" {
						if fields = strings.Fields(str); len(fields) > 0 {
							if data := fields[len(fields)-1]; isVersion(data) { // 判断版本信息
								sdf.Data += line       // mol data
								sdf.Comment = sdf.Name // 备注
								sdf.Name = ""          // 无名称
								break
							} else if i == 0 { // 备注
								sdf.Comment = str
							}
						}
					}
				}
			}
		}
		firstLine = "" // 清空

		sdfs = append(sdfs, sdf) // 有内容 Name 或 Data 才添加
		// mol 数据
		for {
			if line, err := reader.ReadString('\n'); err != nil { // line 是原始数据，包含 \n
				if err == io.EOF {
					return sdfs, nil
				} else {
					return nil, err
				}
			} else {
				sdf.Data += line
				if strings.HasPrefix(line, "M  END") {
					break // 结束 mol 数据 部分
				}
			}
		}

		// 属性
		for {
			if line, err := reader.ReadString('\n'); err != nil { // line 是原始数据，包含 \n
				if err == io.EOF {
					return sdfs, nil
				} else {
					return nil, err
				}
			} else if data := strings.TrimSpace(line); data == "" {
				continue // 空白行
			} else if data == "$$$$" { // 结束
				break
			} else if sAttr := regAttr.FindStringSubmatch(data); len(sAttr) >= 2 { // 属性 名称
				attr := new(Attribute)
				attr.Name = strings.TrimSpace(sAttr[1]) // 获取 命名 分组值，有可能为空名称
				for {                                   // 属性 值
					if line, err = reader.ReadString('\n'); err != nil { // line 是原始数据，包含 \n
						if err == io.EOF {
							return sdfs, nil
						} else {
							return nil, err
						}
					} else if data = strings.TrimSpace(line); data == "" || data == "$$$$" {
						break // 空白行 或 结束
					} else {
						attr.Value += line
					}
				}
				sdf.AddAttr(attr)
			} else { // 新 对象数据
				firstLine = line // 不能 trim ，可能是 mol data
				break
			}
		}

		if readCount > 0 && count >= readCount { // 是否指定读取多少个
			break
		}
	}

	return sdfs, nil
}

func MolCountsByFile(sdFile string) (int, error) {
	if of, err := os.Open(sdFile); err != nil {
		return -1, err
	} else {
		defer of.Close()
		count := 0
		reader := bufio.NewReader(of)
		for {
			if line, err := reader.ReadString('\n'); err != nil { // line 是原始数据，包含 \n
				if err == io.EOF {
					return count, nil
				} else {
					return -1, err
				}
			} else if data := strings.TrimSpace(line); data == "$$$$" {
				count++
			}
		}
	}

}

//-----------------------------------------------------------------------------------------------------------//
