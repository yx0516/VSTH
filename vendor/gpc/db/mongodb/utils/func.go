package utils

import (
	"errors"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//-----------------------------------------------------------------------------------------------------------//

// 检查 指定 【集合】 里的 查询表达式 是否存在
func Exist(coll *mgo.Collection, match interface{}) (ok bool, err error) {
	var n int
	if n, err = coll.Find(match).Limit(1).Count(); err == nil && n > 0 {
		ok = true
	}
	return
}

// 检查 指定 【集合】 里的 【属性】 【值】 是否存在
func ExistValue(coll *mgo.Collection, field string, value interface{}) (ok bool, err error) {
	return Exist(coll, bson.M{field: value})
}

//-----------------------------------------------------------------------------------------------------------//

// 生成 获取下个自增Id的函数
func GenNextId(collCounter *mgo.Collection, collName string) (int, error) {
	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"id": 1}},
		Upsert:    true,
		ReturnNew: true,
	}
	doc := struct{ Id int }{}
	if _, err := collCounter.Find(bson.M{"_id": collName}).Apply(change, &doc); err != nil {
		return 0, errors.New("GenNextId error:" + err.Error())
	}
	return doc.Id, nil
}

//-----------------------------------------------------------------------------------------------------------//
