package utils

import (
	"gopkg.in/mgo.v2"
)

//-----------------------------------------------------------------------------------------------------------//

// 查询助手
type QueryHelper struct {
	Match  OperatorSelect // 查询条件
	Fields QuerySelect    // 查询返回的字段
	Sort   []string       // 需排序的字段,字段名前面加减号-表示是 desc 排序,否则就是 asc 排序
	Skip   int            // 起始位置【不能为负数】
	Limit  int            // 限制返回几条记录
}

func NewQueryHelper() *QueryHelper {
	return &QueryHelper{
		Match:  NewOperatorSelect(),
		Fields: NewQuerySelect(),
	}
}

// 生成 查询器
func (self *QueryHelper) MakeQuery(coll *mgo.Collection) *mgo.Query {
	return coll.Find(self.Match).Select(self.Fields).Sort(self.Sort...).Skip(self.Skip).Limit(self.Limit)
}

//-----------------------------------------------------------------------------------------------------------//

// 查询 集合
type QueryColl struct {
	coll *mgo.Collection
	QueryHelper
}

func NewQueryColl(collection *mgo.Collection) *QueryColl {
	return &QueryColl{
		coll:        collection,
		QueryHelper: *NewQueryHelper(),
	}
}

// 查询
func (self *QueryColl) Query() *mgo.Query {
	return self.MakeQuery(self.coll)
}

// 设置新的 集合【记得关闭旧的】
func (self *QueryColl) SetColl(coll *mgo.Collection) *QueryColl {
	self.coll = coll
	return self
}

// 关闭 Session
func (self *QueryColl) Close() {
	if self.coll != nil {
		self.coll.Database.Session.Close()
	}
}

//-----------------------------------------------------------------------------------------------------------//
