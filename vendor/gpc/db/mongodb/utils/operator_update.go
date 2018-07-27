package utils

import (
	"gopkg.in/mgo.v2/bson"

	plib "gpc/publib"
)

//-----------------------------------------------------------------------------------------------------------//

// 更新 操作
type OperatorUpdate bson.M

func NewOperatorUpdate() OperatorUpdate {
	return make(OperatorUpdate)
}

//-----------------------------------------------------------------------------------------------------------//

func (self OperatorUpdate) JsonString() string {
	return plib.JsonMarshalPrettyToString(self)
}

// 转成 数组
func (self OperatorUpdate) ToArray() (array []bson.M) {
	for k, v := range self {
		array = append(array, bson.M{k: v})
	}
	return
}

// 清空数据
func (self OperatorUpdate) Clear() OperatorUpdate {
	for k, _ := range self {
		delete(self, k)
	}
	return self
}

// 添加 匹配【替换整个文档值】
func (self OperatorUpdate) Add(field string, value interface{}) OperatorUpdate {
	self[field] = value
	return self
}

//-----------------------------------------------------------------------------------------------------------//

// 修改值
func (self OperatorUpdate) Set(field string, value interface{}) OperatorUpdate {
	return self.addMulti("$set", field, value)
}

//-----------------------------------------------------------------------------------------------------------//

// 支持多值的操作
func (self OperatorUpdate) addMulti(operator, field string, value interface{}) OperatorUpdate {
	if m, ok := self[operator]; ok {
		if bm, ok := m.(bson.M); ok {
			bm[field] = value
		} else {
			self[operator] = bson.M{field: value}
		}
	} else {
		self[operator] = bson.M{field: value}
	}
	return self
}

//-----------------------------------------------------------------------------------------------------------//

// 往数组 尾部添加文档，如果 field 为数组则往数组尾部添加，否则为新加字段，内容为数组类型
func (self OperatorUpdate) Push(field string, value interface{}) OperatorUpdate {
	return self.addMulti("$push", field, value)
}

// value 必须为数组类型，往数组 尾部添加多个文档，如果 field 为数组则往数组尾部添加，否则为新加字段，内容为数组类型
func (self OperatorUpdate) PushEach(field string, value interface{}) OperatorUpdate {
	return self.addMulti("$push", field, bson.M{"$each": value})
}

// 往数组 尾部添加文档【不重复才添加-只对单元素的数组才有效，存在时不会再添加，不会有错误】【参考Push】
func (self OperatorUpdate) AddToSet(field string, value interface{}) OperatorUpdate {
	return self.addMulti("$addToSet", field, value)
}

// value 必须为数组类型，往数组 尾部添加多个文档【不重复才添加-只对单元素的数组才有效，存在时不会再添加，不会有错误】【参考PushEach】
func (self OperatorUpdate) AddToSetEach(field string, value interface{}) OperatorUpdate {
	return self.addMulti("$addToSet", field, bson.M{"$each": value})
}

// 删除 数组 头或尾的元素【不存在都不会报错】
func (self OperatorUpdate) Pop(field string, isHeader bool) OperatorUpdate {
	idx := 1 // 【index>=0 删除尾部最后一个, index<0 删除头第一个】
	if isHeader {
		idx = -1
	}
	return self.addMulti("$pop", field, idx)
}

// 删除 数组 里匹配的元素【不存在都不会报错】
func (self OperatorUpdate) Pull(field string, value interface{}) OperatorUpdate {
	return self.addMulti("$pull", field, value)
}

// value 必须为数组类型，删除数组里多个匹配的元素【不存在都不会报错】
func (self OperatorUpdate) PullAll(field string, value interface{}) OperatorUpdate {
	return self.addMulti("$pullAll", field, value)
}

//-----------------------------------------------------------------------------------------------------------//

// 重命名 键名 或 子文档键名【数组无法修改,子文档键名用.定位】
func (self OperatorUpdate) Rename(oldKeyName, newKeyName string) OperatorUpdate {
	return self.addMulti("$rename", oldKeyName, newKeyName)
}

// 删除 键名 或 子文档键名【数组无法删除,子文档键名用.定位】
func (self OperatorUpdate) Unset(keyNames ...string) OperatorUpdate {
	if count := len(keyNames); count > 0 {
		m := make(bson.M)
		for _, keyName := range keyNames {
			m[keyName] = 1 // 键值占位的
		}
		self["$unset"] = m
	}
	return self
}

// 原子性递增/递减 n 值
func (self OperatorUpdate) Inc(field string, n int) OperatorUpdate {
	return self.addMulti("$inc", field, n)
}

//-----------------------------------------------------------------------------------------------------------//
