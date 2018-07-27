package db

import (
	//	"time"
	//"errors"
	//"strings"
	//"time"

	//	"gopkg.in/mgo.v2/bson"

	//"gpc/db/mongodb/cst"
	"gpc/db/mongodb/xorm"
)

//-----------------------------------------------------------------------------------------------------------//

func (self *LibInfo) Init() {
	self.ModelPublic.Init(self)

}

//-----------------------------------------------------------------------------------------------------------//

// 查询
type QueryLibInfo struct {
	xorm.QueryColl
}

func NewQueryLibInfo() *QueryLibInfo {
	v := new(LibInfo)
	v.Init()
	return &QueryLibInfo{
		*xorm.NewQueryColl(v.GetCollName()),
	}
}

func (self *QueryLibInfo) QueryData() (rows []LibInfo, err error) {
	if err = self.Query(&rows); err == nil {
		count := len(rows)
		for i := 0; i < count; i++ {
			rows[i].Init()
		}
	}
	return
}

//-----------------------------------------------------------------------------------------------------------//
