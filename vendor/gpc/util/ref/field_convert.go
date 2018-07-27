package ref

import (
	"strings"

	"gpc/web/cst"

	plib "gpc/publib"
)

const (
	// 字段名转换函数类型
	CST_FNAME_CONVERT_TO_LOWER uint8 = 1
	CST_FNAME_CONVERT_TO_UPPER uint8 = 2
	CST_FNAME_CONVERT_TO_SNAKE uint8 = 3
	CST_FNAME_CONVERT_TO_CAMEL uint8 = 4
)

//-----------------------------------------------------------------------------------------------------------//

// 字段名 换行 函数
func FieldNameConvertFunc(fNameConvert uint8) func(s string) string {
	switch fNameConvert { // 默认 保持不变
	case CST_FNAME_CONVERT_TO_LOWER: // 转成 全小写
		return strings.ToLower
	case CST_FNAME_CONVERT_TO_UPPER: // 转成 全大写
		return strings.ToUpper
	case CST_FNAME_CONVERT_TO_SNAKE: // 转成 驼峰大小写 字符串 转 下划线连接格式【 XxYy to xx_yy】
		return plib.SnakeString
	case CST_FNAME_CONVERT_TO_CAMEL: // 转成 下划线连接字符串 转 驼峰大小写格式【xx_yy to XxYy】
		return plib.CamelString
	default:
		return func(s string) string { // 默认 字段名 保持不变
			return s
		}
	}
}

//-----------------------------------------------------------------------------------------------------------//

// 字段名称 转换接口
type NameConvert interface {
	Convert(map[string]*StructField) map[string]*StructField
}

// 字段名 转换结构体
type FieldNameConvert struct {
	To          func(string) string
	Tag         string
	TagGet      func(string) string
	IgnoreToAdd bool
}

// 字段名 转换函数【模板】
func (self FieldNameConvert) Convert(src map[string]*StructField) (dst map[string]*StructField) {
	if self.To == nil {
		return src
	}
	dst = make(map[string]*StructField)
	self.Tag = strings.TrimSpace(self.Tag)
	if self.Tag == "" {
		for name, sfi := range src {
			dst[self.To(name)] = sfi
		}
	} else {
		for name, sfi := range src {
			if body, ok := sfi.Tags[self.Tag]; ok {
				if body == cst.TAG_IGNORE {
					if self.IgnoreToAdd {
						dst[name] = sfi
					}
				} else if self.TagGet != nil {
					if body = strings.TrimSpace(self.TagGet(body)); body != "" {
						dst[body] = sfi
					} else {
						dst[self.To(name)] = sfi
					}
				} else {
					dst[body] = sfi
				}
			} else {
				dst[self.To(name)] = sfi
			}
		}
	}
	return
}

//-----------------------------------------------------------------------------------------------------------//

// 字段名转成 全小写【可指定使用 tag 和 处理 tag 获取 最终的字段名】
func NewFiledNameToLower(tag string, get func(string) string, ignoreToAdd ...bool) NameConvert {
	return FieldNameConvert{
		To:          strings.ToLower,
		Tag:         tag,
		TagGet:      get,
		IgnoreToAdd: len(ignoreToAdd) > 0 && ignoreToAdd[0],
	}

}

// 字段名转成 全大写
func NewFiledNameToUpper(tag string, get func(string) string, ignoreToAdd ...bool) NameConvert {
	return FieldNameConvert{
		To:          strings.ToUpper,
		Tag:         tag,
		TagGet:      get,
		IgnoreToAdd: len(ignoreToAdd) > 0 && ignoreToAdd[0],
	}
}

// 字段名 驼峰大小写 字符串转下划线连接格式【 XxYy to xx_yy】
func NewFiledNameToSnake(tag string, get func(string) string, ignoreToAdd ...bool) NameConvert {
	return FieldNameConvert{
		To:          plib.SnakeString,
		Tag:         tag,
		TagGet:      get,
		IgnoreToAdd: len(ignoreToAdd) > 0 && ignoreToAdd[0],
	}
}

// 字段名 下划线连接字符串转驼峰大小写格式【xx_yy to XxYy】
func NewFiledNameToCamel(tag string, get func(string) string, ignoreToAdd ...bool) NameConvert {
	return FieldNameConvert{
		To:          plib.CamelString,
		Tag:         tag,
		TagGet:      get,
		IgnoreToAdd: len(ignoreToAdd) > 0 && ignoreToAdd[0],
	}
}

//-----------------------------------------------------------------------------------------------------------//

// beego orm 结构体的字段名转换
func NewFiledNameBeegoOrmConvert(fNameConvert uint8) NameConvert {
	return FieldNameConvert{
		To:  FieldNameConvertFunc(fNameConvert),
		Tag: cst.TAG_BEEGO_ORM,
		TagGet: func(tag string) (name string) {
			if tag != cst.TAG_IGNORE {
				for _, s := range strings.Split(tag, ";") {
					s = strings.TrimSpace(s)
					if strings.HasPrefix(s, "column(") && strings.HasSuffix(s, ")") {
						return strings.TrimSpace(strings.TrimSuffix(strings.TrimPrefix(s, "column("), ")"))
					}
				}
			}
			return
		},
		IgnoreToAdd: false,
	}
}

// mongodb 结构体的字段名转换
func NewFiledNameMgoConvert(fNameConvert uint8) NameConvert {
	return FieldNameConvert{
		To:  FieldNameConvertFunc(fNameConvert),
		Tag: cst.TAG_MGO_BSON,
		TagGet: func(tag string) (name string) {
			if tag != cst.TAG_IGNORE {
				return tag
			}
			return
		},
		IgnoreToAdd: false,
	}
}

//-----------------------------------------------------------------------------------------------------------//
