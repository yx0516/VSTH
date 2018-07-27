package db

import (
	//"errors"
	//"strings"
	//"time"

	//"gopkg.in/mgo.v2/bson"

	//"gpc/db/mongodb/cst"
	"gpc/db/mongodb/xorm"
)

//-----------------------------------------------------------------------------------------------------------//

func (self *ZincMolSupply) Init() {
	self.ModelPublic.Init(self)

}

//-----------------------------------------------------------------------------------------------------------//

// 查询
type QueryZincMolSupply struct {
	xorm.QueryColl
}

func NewQueryZincMolSupply() *QueryZincMolSupply {
	v := new(ZincMolSupply)
	v.Init()
	return &QueryZincMolSupply{
		*xorm.NewQueryColl(v.GetCollName()),
	}
}

func (self *QueryZincMolSupply) QueryData() (rows []ZincMolSupply, err error) {
	if err = self.Query(&rows); err == nil {
		count := len(rows)
		for i := 0; i < count; i++ {
			rows[i].Init()
		}
	}
	return
}

//-----------------------------------------------------------------------------------------------------------//
