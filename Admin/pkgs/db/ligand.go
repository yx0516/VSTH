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

func (self *Ligand) Init() {
	self.ModelPublic.Init(self)

	self.Model.FuncMakeMatch = func() (bson.M, error) {
		return bson.M{"name": self.Name}, nil
	}

	self.Model.FuncUpdateByFields = func(fvs xorm.MapSI, vSelf interface{}) (err error) {
		if fvs.Len() > 0 {

		}
		return
	}
}

//-----------------------------------------------------------------------------------------------------------//

// 查询
type QueryLigand struct {
	xorm.QueryColl
}

func NewQueryLigand() *QueryLigand {
	v := new(Ligand)
	v.Init()
	return &QueryLigand{
		*xorm.NewQueryColl(v.GetCollName()),
	}
}

func (self *QueryLigand) QueryData() (rows []Ligand, err error) {
	if err = self.Query(&rows); err == nil {
		count := len(rows)
		for i := 0; i < count; i++ {
			rows[i].Init()
		}
	}
	return
}

//-----------------------------------------------------------------------------------------------------------//
