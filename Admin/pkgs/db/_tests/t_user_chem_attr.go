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
	attr := db.NewUserChemAttr("att1")
	attr.UserId = 1
	attr.ChemId = 1
	//	attr.Name = ""

	pp(attr.Insert())

	plib.Dump(attr)
}

func t2() {
	attr := db.NewUserChemAttr("att1")
	//	attr.Id = 1
	attr.UserId = 5
	attr.ChemId = 1
	//	attr.Name = ""

	//	pp(attr.Read())
	pp(attr.ReadByName())
	//	pp(attr.ReadById())

	plib.Dump(attr)
}

func t3() {
	attr := db.NewUserChemAttr("att1")
	attr.Id = 1
	//	attr.UserId=2
	//	attr.ChemId=4
	attr.Name = "att2"

	attr.Type.SetString()
	attr.Status.SetOff()
	attr.Unit = "xxxxxxxxxxxxxx12121212"

	pp(attr.Update())

	plib.Dump(attr)
}

func t4() {
	attr := db.NewUserChemAttr("att3")
	//	attr.Id=2
	//	attr.UserId=3
	attr.ChemId = 4

	pp(attr.Delete())

}

func t5() {
	for i := 0; i < 20; i++ {
		attr := db.NewUserChemAttr(fmt.Sprintf("attr%v", i))
		attr.UserId = 1
		attr.ChemId = 1

		pp(attr.Insert())
	}
}

func t6() {
	v := db.NewQueryUserChemAttr()

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
