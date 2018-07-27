package utils

import (
	"gopkg.in/mgo.v2/bson"
)

//-----------------------------------------------------------------------------------------------------------//

// Query.Select 的参数助手【Query.Select 调用多次只有最后一次有效，所以可以缓存下这个参数】
type QuerySelect bson.M

func NewQuerySelect() QuerySelect {
	return make(QuerySelect)
}

func (self QuerySelect) add(ok bool, fields ...string) QuerySelect {
	count := len(fields)
	if count > 0 {
		for i := 0; i < count; i++ {
			self[fields[i]] = ok
		}
	}
	return self
}

// 需要 返回的字段
func (self QuerySelect) True(fields ...string) QuerySelect {
	return self.add(true, fields...)
}

// 不需要 返回的字段
func (self QuerySelect) False(fields ...string) QuerySelect {
	return self.add(false, fields...)
}

// 控制查询返回的数组中元素的个数
// field 为数组键值名, value 为数字(正数取前面几个，负数取后面几个) 或 数字数组[]int{10,30}实现类似切片
func (self QuerySelect) Slice(field string, value interface{}) QuerySelect {
	self[field] = bson.M{"$slice": value}
	return self
}

//-----------------------------------------------------------------------------------------------------------//
