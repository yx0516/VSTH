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

func (self *TargetPDB) Init() {
	self.ModelPublic.Init(self)

	self.Model.FuncMakeMatch = func() (bson.M, error) {
		return bson.M{"pdbcode": self.PDBCode}, nil
	}

}

//-----------------------------------------------------------------------------------------------------------//

// 查询
type QueryTargetPDB struct {
	xorm.QueryColl
}

func NewQueryTargetPDB() *QueryTargetPDB {
	v := new(TargetPDB)
	v.Init()
	return &QueryTargetPDB{
		*xorm.NewQueryColl(v.GetCollName()),
	}
}

func (self *QueryTargetPDB) QueryData() (rows []TargetPDB, err error) {
	if err = self.Query(&rows); err == nil {
		count := len(rows)
		for i := 0; i < count; i++ {
			rows[i].Init()
		}
	}
	return
}

//-----------------------------------------------------------------------------------------------------------//
