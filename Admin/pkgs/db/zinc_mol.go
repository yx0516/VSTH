package db

import (
	"errors"
	"strings"
	//"time"

	"gopkg.in/mgo.v2/bson"

	"gpc/db/mongodb/cst"
	"gpc/db/mongodb/xorm"
)

//-----------------------------------------------------------------------------------------------------------//

func (self *ZincMol) Init() {
	self.ModelPublic.Init(self)

	self.Model.FuncCheckInsert = func() (err error) {
		if err = self.checkExistName(); err != nil {
			return
		}

		//插入数据时需要系统初始化的参数放在这里，如：时间等

		return
	}

	self.Model.FuncMakeMatch = func() (bson.M, error) {
		if err := obj.chk.Name(self.Name); err != nil {
			return nil, err
		} else {
			return bson.M{cst.Field_Name: self.Name}, nil
		}
	}

	self.Model.FuncUpdateByFields = func(fvs xorm.MapSI, vSelf interface{}) (err error) {
		v, ok := vSelf.(*ZincMol)
		if !ok {
			return errors.New("vSelf interface convert fail.")
		}

		//plib.Dump(fvs)

		if fvs.Contains(cst.Field_Name) {
			v.Name = self.Name
			if err = v.checkExistName(); err != nil {
				return
			}
		}

		return
	}

}

//-----------------------------------------------------------------------------------------------------------//

//
// same zincId exist
//
func (self *ZincMol) ExistName() (ok bool, err error) {
	self.Name = strings.TrimSpace(self.Name)
	return self.Model.ExistName(self.Name, obj.chk.Name)
}

func (self *ZincMol) checkExistName() error {
	if ok, err := self.ExistName(); err != nil {
		return err
	} else if ok {
		return errors.New(obj.pack.AlreadyExistName(self.Name))
	}
	return nil
}

//-----------------------------------------------------------------------------------------------------------//

func (self *ZincMol) ReadByName() error {
	return self.Model.ReadByName(self.Name, obj.chk.Name)
}

func (self *ZincMol) DeleteByName() (int, error) {
	return self.Model.DeleteByName(self.Name, obj.chk.Name)
}

//-----------------------------------------------------------------------------------------------------------//

// 查询
type QueryZincMol struct {
	xorm.QueryColl
}

func NewQueryZincMol() *QueryZincMol {
	v := new(ZincMol)
	v.Init()
	return &QueryZincMol{
		*xorm.NewQueryColl(v.GetCollName()),
	}
}

func (self *QueryZincMol) QueryData() (rows []ZincMol, err error) {
	if err = self.Query(&rows); err == nil {
		count := len(rows)
		for i := 0; i < count; i++ {
			rows[i].Init()
		}
	}
	return
}

//-----------------------------------------------------------------------------------------------------------//
