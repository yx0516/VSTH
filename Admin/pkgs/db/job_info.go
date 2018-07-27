package db

import (
	//	"time"
	//"errors"
	//"strings"
	//"time"

	"gopkg.in/mgo.v2/bson"

	//"gpc/db/mongodb/cst"
	"gpc/db/mongodb/xorm"
)

//-----------------------------------------------------------------------------------------------------------//

func (self *JobInfo) Init() {
	self.ModelPublic.Init(self)

	self.Model.FuncMakeMatch = func() (bson.M, error) {
		return bson.M{"jobid": self.JobId}, nil
	}

	//	self.Model.FuncCheckInsert = func() (err error) {
	//		if _, err = obj.doc.CheckUserById(self.UserId); err != nil { // 检查 用户 是否存在，并且可用
	//			return
	//		}

	//		// 初始化 参数值
	//		//self.StartTime = time.Now().Format("")
	//		return
	//	}
}

//-----------------------------------------------------------------------------------------------------------//

// 查询
type QueryJobInfo struct {
	xorm.QueryColl
}

func NewQueryJobInfo() *QueryJobInfo {
	v := new(JobInfo)
	v.Init()
	return &QueryJobInfo{
		*xorm.NewQueryColl(v.GetCollName()),
	}
}

func (self *QueryJobInfo) QueryData() (rows []JobInfo, err error) {
	if err = self.Query(&rows); err == nil {
		count := len(rows)
		for i := 0; i < count; i++ {
			rows[i].Init()
		}
	}
	return
}

//-----------------------------------------------------------------------------------------------------------//
