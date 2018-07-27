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

func (self *LigandActivityRef) Init() {
	self.ModelPublic.Init(self)

	self.Model.FuncMakeMatch = func() (bson.M, error) {
		return bson.M{"ligandactivityid": self.LigandActivityId}, nil
	}

	self.Model.FuncMakeMatch = func() (bson.M, error) {
		return bson.M{"name": self.Name}, nil
	}
}

//-----------------------------------------------------------------------------------------------------------//

// 查询
type QueryLigandActivityRef struct {
	xorm.QueryColl
}

func NewQueryLigandActivityRef() *QueryLigandActivityRef {
	v := new(LigandActivityRef)
	v.Init()
	return &QueryLigandActivityRef{
		*xorm.NewQueryColl(v.GetCollName()),
	}
}

func (self *QueryLigandActivityRef) QueryData() (rows []LigandActivityRef, err error) {
	if err = self.Query(&rows); err == nil {
		count := len(rows)
		for i := 0; i < count; i++ {
			rows[i].Init()
		}
	}
	return
}

//-----------------------------------------------------------------------------------------------------------//
