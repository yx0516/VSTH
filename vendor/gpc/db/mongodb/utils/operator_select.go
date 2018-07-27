package utils

import (
	"strings"

	"gopkg.in/mgo.v2/bson"

	plib "gpc/publib"
)

//-----------------------------------------------------------------------------------------------------------//

// 选择 操作
type OperatorSelect bson.M

func NewOperatorSelect() OperatorSelect {
	return make(OperatorSelect)
}

//-----------------------------------------------------------------------------------------------------------//

func (self OperatorSelect) JsonString() string {
	return plib.JsonMarshalPrettyToString(self)
}

// 转成 数组
func (self OperatorSelect) ToArray() (array []bson.M) {
	for k, v := range self {
		array = append(array, bson.M{k: v})
	}
	return
}

// 清空数据
func (self OperatorSelect) Clear() OperatorSelect {
	for k, _ := range self {
		delete(self, k)
	}
	return self
}

// 添加 匹配，默认 $and 关系
func (self OperatorSelect) Add(field string, value interface{}) OperatorSelect {
	self[field] = value
	return self
}

//-----------------------------------------------------------------------------------------------------------//

// 添加 exists 操作
func (self OperatorSelect) Exists(field string, ok bool) OperatorSelect {
	self[field] = bson.M{"$exists": ok}
	return self
}

// 匹配字段 null 值
func (self OperatorSelect) Null(field string) OperatorSelect {
	self[field] = nil
	return self
}

/*
	i 如果设置了这个修饰符，模式中的字母会进行大小写不敏感匹配。
	m 默认情况下，PCRE 认为目标字符串是由单行字符组成的(然而实际上它可能会包含多行).
	  如果目标字符串 中没有 "\n"字符，或者模式中没有出现“行首”/“行末”字符，设置这个修饰符不产生任何影响。
	s 如果设置了这个修饰符，模式中的点号元字符匹配所有字符，包含换行符。如果没有这个修饰符，点号不匹配换行符。
	x 如果设置了这个修饰符，模式中的没有经过转义的或不在字符类中的空白数据字符总会被忽略，
	  并且位于一个未转义的字符类外部的#字符和下一个换行符之间的字符也被忽略。这个修饰符使被编译模式中可以包含注释。
	  注意：这仅用于数据字符。 空白字符 还是不能在模式的特殊字符序列中出现，比如序列。
*/
// field 数字键名 , expression 正则表达式 ，options 选项：i | m | s | x
func (self OperatorSelect) Regex(field, expression string, options ...string) OperatorSelect {
	m := bson.M{"$regex": expression}
	if len(options) > 0 {
		if option := strings.TrimSpace(options[0]); option != "" {
			m["$options"] = option
		}
	}
	self[field] = m
	return self
}

// 使用 JavaScript 函数 查询，例如："this.apistatus == this.status" 【expression 不能为空】【this | obj】
func (self OperatorSelect) Where(expression string) OperatorSelect {
	self["$where"] = expression
	return self
}

// 限制查询返回的数组字段的内容只包含匹配elemMatch条件的数组元素
// 注意：数组中元素是内嵌文档。如果多个元素匹配$elemMatch条件，操作符返回数组中第一个匹配条件的元素。
// 注意：返回的这个 *OperatorQuery 是 ElemMatch 的
func (self OperatorSelect) ElemMatch(field string) OperatorSelect {
	oq := NewOperatorSelect()
	self[field] = bson.M{"$elemMatch": oq}
	return oq
}

//-----------------------------------------------------------------------------------------------------------//

// 支持 not 操作的类型
func (self OperatorSelect) addNot(field string, operator string, value interface{},
	isNot bool, isExists ...bool) OperatorSelect {
	m := make(bson.M)
	if isNot {
		m["$not"] = bson.M{"$" + operator: value}
	} else {
		m["$"+operator] = value
	}
	if len(isExists) > 0 {
		m["$exists"] = isExists[0]
	}
	self[field] = m
	return self
}

// value 数字类型【大于 >】
func (self OperatorSelect) Gt(isNot bool, field string, value interface{}, isExists ...bool) OperatorSelect {
	return self.addNot(field, "gt", value, isNot, isExists...)
}

// value 数字类型【大于等于 >=】
func (self OperatorSelect) Gte(isNot bool, field string, value interface{}, isExists ...bool) OperatorSelect {
	return self.addNot(field, "gte", value, isNot, isExists...)
}

// value 数字类型【小于 <】
func (self OperatorSelect) Lt(isNot bool, field string, value interface{}, isExists ...bool) OperatorSelect {
	return self.addNot(field, "lt", value, isNot, isExists...)
}

// value 数字类型【小于等于 <=】
func (self OperatorSelect) Lte(isNot bool, field string, value interface{}, isExists ...bool) OperatorSelect {
	return self.addNot(field, "lte", value, isNot, isExists...)
}

// value 数组类型【必须全部匹配】
func (self OperatorSelect) All(isNot bool, field string, value interface{}, isExists ...bool) OperatorSelect {
	return self.addNot(field, "all", value, isNot, isExists...)
}

// value 数组类型【只要匹配一个】
func (self OperatorSelect) In(isNot bool, field string, value interface{}, isExists ...bool) OperatorSelect {
	return self.addNot(field, "in", value, isNot, isExists...)
}

// value 数组类型
func (self OperatorSelect) Nin(isNot bool, field string, value interface{}, isExists ...bool) OperatorSelect {
	return self.addNot(field, "nin", value, isNot, isExists...)
}

// field 数组键名 , value 数字类型
func (self OperatorSelect) Size(isNot bool, field string, value interface{}, isExists ...bool) OperatorSelect {
	return self.addNot(field, "size", value, isNot, isExists...)
}

// field 数字键名 , value 数字数组 []int{ 模除数 , 余数 }
func (self OperatorSelect) Mod(isNot bool, field string, value interface{}, isExists ...bool) OperatorSelect {
	return self.addNot(field, "mod", value, isNot, isExists...)
}

// value 单值，排除指定内容外的其它
func (self OperatorSelect) Ne(isNot bool, field string, value interface{}, isExists ...bool) OperatorSelect {
	return self.addNot(field, "ne", value, isNot, isExists...)
}

//-----------------------------------------------------------------------------------------------------------//

// 添加 $or 操作
// values: bson.M {"name":"value", ...}
func (self OperatorSelect) Or(values ...interface{}) OperatorSelect {
	if arr, ok := self["$or"]; !ok {
		self["$or"] = values
	} else if _, ok = arr.([]interface{}); !ok {
		self["$or"] = values
	} else {
		self["$or"] = append(self["$or"].([]interface{}), values...)
	}

	return self
}

// 添加 $and 操作，默认就是 $and
// values: bson.M {"name":"value", ...}
func (self OperatorSelect) And(values ...interface{}) OperatorSelect {
	if arr, ok := self["$and"]; !ok {
		self["$and"] = values
	} else if _, ok = arr.([]interface{}); !ok {
		self["$and"] = values
	} else {
		self["$and"] = append(self["$and"].([]interface{}), values...)
	}

	return self
}

//-----------------------------------------------------------------------------------------------------------//
