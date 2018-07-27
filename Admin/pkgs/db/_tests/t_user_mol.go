package main

import (
	"fmt"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	plib "gpc/publib"

	rs "gpc/util/randstr"

	"gpc/db/mongodb/conn"
	"gpc/db/mongodb/conn/gsession"
	"gpc/db/mongodb/xorm"

	"rcdd/project/Chem/Admin/pkgs/db"
)

var (
	pp  = fmt.Println
	ppp = func() { pp("######################################") }
	_   = plib.Debug
	_   = spew.Config
	_   = log.Ldate
	_   = make(bson.M)
	_   = xorm.Log
	_   = os.ErrNotExist

	session  *conn.Session
	database *mgo.Database

	randStrNum func(n int) string
)

func init() {
	//	mgo.SetDebug(true)
	//	mgo.SetLogger(log.New(os.Stdout, "[mgo] ", log.Ldate|log.Ltime))

	randStrNum = rs.RandAlphaDigit

	c := conn.NewConn(
		"127.0.0.1",
		27017,
	)
	c.DB = "chem"
	c.UserName = "chemUser"
	c.UserPwd = "chemUserPwd"
	c.Timeout = 3

	var err error
	if session, err = c.Login(); err != nil {
		plib.OsExitPrint(1, err.Error())
	}
	database = session.DB(c.DB)

	// 初始化 orm session
	gsession.InitSession(session.Session, c.DB)
}

func t1() {
	mol := db.NewUserMol("测试分子1")
	//	mol.Name = ""
	mol.Data = "  1212 222 3333   "
	mol.ChemId = 1
	mol.UserId = 1

	pp(mol.Insert())

	plib.Dump(mol)
}

func t2() {
	mol := db.NewUserMol("user2")
	mol.Name = ""

	//	err := mol.ReadByName()
	//	err := mol.Read()

	mol.Id = 2

	err := mol.ReadById()

	if err != nil {
		pp(err)
	} else {
		plib.Dump(mol)
	}
}

func t3() {
	mol := db.NewUserMol("user1")
	mol.Name = ""
	mol.Id = 5

	mol.Status = 2

	pp(mol.Update())
	//	pp(mol.UpdateByFields([]string{"status"}))
}

func t4() {
	mol := db.NewUserMol("user2")

	//	mol.Name = ""
	//	pp(mol.Delete())
	//	pp(mol.DeleteById())
	pp(mol.DeleteByName())
}

func t5() {
	for i := 0; i < 20; i++ {
		//		mol := db.NewUserMol(randStrNum(10))
		mol := db.NewUserMol(fmt.Sprintf("mol%v", i))
		mol.UserId = 1
		mol.ChemId = 1
		mol.Data = "xx111"
		//	mol.Name = ""

		pp(mol.Insert())
	}
}

func t6() {
	v := db.NewQueryUserMol()

	v.Limit = 2
	v.Sort = []string{"id"}
	v.Match.Gt(false, "id", 1)

	pp(v.Count())
	pp(v.CountAll())

	if rows, err := v.QueryData(); err != nil {
		pp(err)
	} else {
		plib.Dump(rows)
	}

	if rows, err := v.GetQuery(); err != nil {
		pp(err)
	} else {
		plib.Dump(rows)
	}
}

func main() {
	defer session.SafeClose()

	//	t1()
	//	t2()
	//	t3()
	//	t4()
	//	t5()
	t6()
}
