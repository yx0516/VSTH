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

func (self *LigandActivity) Init() {
	self.ModelPublic.Init(self)

	self.Model.FuncMakeMatch = func() (bson.M, error) {
		return bson.M{"name": self.Name}, nil
	}
}

//-----------------------------------------------------------------------------------------------------------//

// 查询
type QueryLigandActivity struct {
	xorm.QueryColl
}

func NewQueryLigandActivity() *QueryLigandActivity {
	v := new(LigandActivity)
	v.Init()
	return &QueryLigandActivity{
		*xorm.NewQueryColl(v.GetCollName()),
	}
}

func (self *QueryLigandActivity) QueryData() (rows []LigandActivity, err error) {
	if err = self.Query(&rows); err == nil {
		count := len(rows)
		for i := 0; i < count; i++ {
			rows[i].Init()
		}
	}
	return
}

//-----------------------------------------------------------------------------------------------------------//
