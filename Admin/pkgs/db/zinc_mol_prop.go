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

func (self *ZincMolProp) Init() {
	self.ModelPublic.Init(self)

	self.Model.FuncMakeMatch = func() (bson.M, error) {
		return bson.M{"zincmolname": self.ZincMolName}, nil
	}

}

//-----------------------------------------------------------------------------------------------------------//

// 查询
type QueryZincMolProp struct {
	xorm.QueryColl
}

func NewQueryZincMolProp() *QueryZincMolProp {
	v := new(ZincMolProp)
	v.Init()
	return &QueryZincMolProp{
		*xorm.NewQueryColl(v.GetCollName()),
	}
}

func (self *QueryZincMolProp) QueryData() (rows []ZincMolProp, err error) {
	if err = self.Query(&rows); err == nil {
		count := len(rows)
		for i := 0; i < count; i++ {
			rows[i].Init()
		}
	}
	return
}

//-----------------------------------------------------------------------------------------------------------//
