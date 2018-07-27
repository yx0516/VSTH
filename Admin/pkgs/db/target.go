package db

import (
	//"errors"
	//"strings"
	//"time"

	"gopkg.in/mgo.v2/bson"

	//"gpc/db/mongodb/cst"
	"gpc/db/mongodb/xorm"
)

//-----------------------------------------------------------------------------------------------------------//

func (self *Target) Init() {
	self.ModelPublic.Init(self)

	self.Model.FuncMakeMatch = func() (bson.M, error) {
		return bson.M{"uniprotid": self.UniprotId}, nil
	}

}

//-----------------------------------------------------------------------------------------------------------//

// 查询
type QueryTarget struct {
	xorm.QueryColl
}

func NewQueryTarget() *QueryTarget {
	v := new(Target)
	v.Init()
	return &QueryTarget{
		*xorm.NewQueryColl(v.GetCollName()),
	}
}

func (self *QueryTarget) QueryData() (rows []Target, err error) {
	if err = self.Query(&rows); err == nil {
		count := len(rows)
		for i := 0; i < count; i++ {
			rows[i].Init()
		}
	}
	return
}

//-----------------------------------------------------------------------------------------------------------//
