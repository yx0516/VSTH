package xorm

import (
	"gpc/db/mongodb/conn/gsession"
	"gpc/db/mongodb/utils"
)

//-----------------------------------------------------------------------------------------------------------//

// 查询 集合
type QueryColl struct {
	collName string
	utils.QueryHelper
}

func NewQueryColl(collName string) *QueryColl {
	return &QueryColl{
		collName:    collName,
		QueryHelper: *utils.NewQueryHelper(),
	}
}

func (self *QueryColl) SetCollName(collName string) *QueryColl {
	self.collName = collName
	return self
}

func (self *QueryColl) GetCollName() string {
	return self.collName
}

// 查询数据【rows必须指针】
func (self *QueryColl) Query(rows interface{}) error {
	db := gsession.NewDB()
	defer db.Session.Close()

	query := self.QueryHelper.MakeQuery(db.C(self.collName))
	return obj.err.AssRead(query.All(rows), self.collName+".query")
}

// 查询数据【保存到接口数组里】
func (self *QueryColl) GetQuery() (array []interface{}, err error) {
	err = self.Query(&array)
	return
}

// 查询条数
func (self *QueryColl) Count() (int, error) {
	db := gsession.NewDB()
	defer db.Session.Close()

	query := self.QueryHelper.MakeQuery(db.C(self.collName))
	if n, err := query.Count(); err != nil {
		return 0, obj.err.AssRead(err, self.collName+".queryCount")
	} else {
		return n, nil
	}
}

// 查询条数【无过滤条件】
func (self *QueryColl) CountAll() (int, error) {
	db := gsession.NewDB()
	defer db.Session.Close()

	if n, err := db.C(self.collName).Find(nil).Count(); err != nil {
		return 0, obj.err.AssRead(err, self.collName+".queryAllCount")
	} else {
		return n, nil
	}
}

//-----------------------------------------------------------------------------------------------------------//
