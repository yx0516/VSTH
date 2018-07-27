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

func (self *LigActTarRelation) Init() {
	self.ModelPublic.Init(self)

}

//-----------------------------------------------------------------------------------------------------------//

// 查询
type QueryLigActTarRelation struct {
	xorm.QueryColl
}

func NewQueryLigActTarRelation() *QueryLigActTarRelation {
	v := new(LigActTarRelation)
	v.Init()
	return &QueryLigActTarRelation{
		*xorm.NewQueryColl(v.GetCollName()),
	}
}

func (self *QueryLigActTarRelation) QueryData() (rows []LigActTarRelation, err error) {
	if err = self.Query(&rows); err == nil {
		count := len(rows)
		for i := 0; i < count; i++ {
			rows[i].Init()
		}
	}
	return
}

//-----------------------------------------------------------------------------------------------------------//
