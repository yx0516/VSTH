package db

import (
	//"errors"
	//"strings"
	//"time"

	//	"gopkg.in/mgo.v2/bson"

	//"gpc/db/mongodb/cst"
	"gpc/db/mongodb/xorm"
)

//-----------------------------------------------------------------------------------------------------------//

func (self *JobResult) Init() {
	self.ModelPublic.Init(self)

}

//-----------------------------------------------------------------------------------------------------------//

// 查询
type QueryJobResult struct {
	xorm.QueryColl
}

func NewQueryJobResult() *QueryJobResult {
	v := new(JobResult)
	v.Init()
	return &QueryJobResult{
		*xorm.NewQueryColl(v.GetCollName()),
	}
}

func (self *QueryJobResult) QueryData() (rows []JobResult, err error) {
	if err = self.Query(&rows); err == nil {
		count := len(rows)
		for i := 0; i < count; i++ {
			rows[i].Init()
		}
	}
	return
}

//-----------------------------------------------------------------------------------------------------------//
