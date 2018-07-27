package db

import (
	"errors"
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"

	"gpc/db/mongodb/cst"
	"gpc/db/mongodb/xorm"
)

//-----------------------------------------------------------------------------------------------------------//

func (self *User) Init() {
	self.ModelPublic.Init(self)

	self.Model.FuncCheckInsert = func() (err error) {
		if err = self.checkExistName(); err != nil {
			return
		}

		self.CreateTime = time.Now()
		self.UpdateTime = time.Time{}
		return
	}

	self.Model.FuncMakeMatch = func() (bson.M, error) {
		if err := obj.chk.Name(self.Name); err != nil {
			return bson.M{"email": self.Email}, nil
		} else {
			return bson.M{cst.Field_Name: self.Name}, nil
		}

	}

	self.Model.FuncUpdateByFields = func(fvs xorm.MapSI, vSelf interface{}) (err error) {
		if fvs.Len() > 0 {
			fvs[cst.Field_UpdateTime] = time.Now()
		}
		return
	}
}

//-----------------------------------------------------------------------------------------------------------//

func (self *User) ExistName() (ok bool, err error) {
	self.Name = strings.TrimSpace(self.Name)
	return self.Model.ExistName(self.Name, obj.chk.Name)
}

func (self *User) checkExistName() error {
	if ok, err := self.ExistName(); err != nil {
		return err
	} else if ok {
		return errors.New(obj.pack.AlreadyExistName(self.Name))
	}
	return nil
}

//-----------------------------------------------------------------------------------------------------------//

func (self *User) ReadByName() error {
	return self.Model.ReadByName(self.Name, obj.chk.Name)
}

func (self *User) DeleteByName() (int, error) {
	return self.Model.DeleteByName(self.Name, obj.chk.Name)
}

//-----------------------------------------------------------------------------------------------------------//

// 查询
type QueryUser struct {
	xorm.QueryColl
}

func NewQueryUser() *QueryUser {
	v := new(User)
	v.Init()
	return &QueryUser{
		*xorm.NewQueryColl(v.GetCollName()),
	}
}

func (self *QueryUser) QueryData() (rows []User, err error) {
	if err = self.Query(&rows); err == nil {
		count := len(rows)
		for i := 0; i < count; i++ {
			rows[i].Init()
		}
	}
	return
}

//-----------------------------------------------------------------------------------------------------------//
