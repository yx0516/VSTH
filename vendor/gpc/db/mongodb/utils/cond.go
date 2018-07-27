package utils

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	plib "gpc/publib"
)

//-----------------------------------------------------------------------------------------------------------//

// 操作 拼接器
type Cond struct {
	Coll   *mgo.Collection // 查询的集合
	Select OperatorSelect  // 选择
	Set    OperatorUpdate  // 更新
}

func NewCond(collection *mgo.Collection) *Cond {
	return &Cond{
		Coll:   collection,
		Select: NewOperatorSelect(),
		Set:    NewOperatorUpdate(),
	}
}

// 获取 查询条件 信息
func (self *Cond) JsonString() string {
	return plib.JsonMarshalPrettyToString(self)
}

//-----------------------------------------------------------------------------------------------------------//

// 删除 匹配到的第一个文档
func (self *Cond) Remove() error {
	return self.Coll.Remove(self.Select)
}

// 删除 匹配到的所有文档
func (self *Cond) RemoveAll() (*mgo.ChangeInfo, error) {
	return self.Coll.RemoveAll(self.Select)
}

//-----------------------------------------------------------------------------------------------------------//

// 更新 匹配到的第一个文档
func (self *Cond) Update() error {
	return self.Coll.Update(self.Select, self.Set)
}

// 更新 匹配到的所有文档
func (self *Cond) UpdateAll() (*mgo.ChangeInfo, error) {
	return self.Coll.UpdateAll(self.Select, self.Set)
}

// 更新 匹配到的单个文档，当文档不存在时，插入更新的内容
func (self *Cond) Upsert() (*mgo.ChangeInfo, error) {
	return self.Coll.Upsert(self.Select, self.Set)
}

//-----------------------------------------------------------------------------------------------------------//

func (self *Cond) query(query interface{}) *mgo.Query {
	return self.Coll.Find(query)
}

// 获取 查询 对象
func (self *Cond) Query() *mgo.Query {
	return self.query(self.Select)
}

// 获取 or 查询 对象，Select 不能是空数组
func (self *Cond) QueryOr() *mgo.Query {
	return self.query(bson.M{"$or": self.Select.ToArray()})
}

// 获取 nor 查询 对象，Select 不能是空数组【排除掉 or 匹配的其它文档】
func (self *Cond) QueryNor() *mgo.Query {
	return self.query(bson.M{"$nor": self.Select.ToArray()})
}

// 检查 指定 【集合】 里的 查询表达式 是否存在
func (self *Cond) QueryExist() (ok bool, err error) {
	var n int
	if n, err = self.Coll.Find(self.Select).Limit(1).Count(); err == nil && n > 0 {
		ok = true
	}
	return
}

//-----------------------------------------------------------------------------------------------------------//
