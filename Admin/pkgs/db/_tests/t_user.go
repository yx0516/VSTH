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

	"rcdd/project/VSTH/Admin/pkgs/db"
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
	c.DB = "app"
	c.UserName = "appUser"
	c.UserPwd = "appPwd"
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
	user := db.NewUser("user2")
	//	user.Name = ""
	user.Status = 3

	pp(user.Insert())

	plib.Dump(user)
}

func t2() {
	user := db.NewUser("user1")
	//	user.Name = ""

	//	err := user.ReadByName()
	err := user.Read()

	//	user.Id = 2

	//	err := user.ReadById()

	if err != nil {
		pp(err)
	} else {
		plib.Dump(user)
	}
}

func t3() {
	user := db.NewUser("user1")
	//	user.Name = ""
	//	user.Id = 5

	user.Status = 1

	pp(user.Update())
	//	pp(user.UpdateByFields([]string{"status"}))
}

func t4() {
	user := db.NewUser("user2")

	//	user.Name = ""
	//	pp(user.Delete())
	//	pp(user.DeleteById())
	pp(user.DeleteByName())
}

func t5() {
	v := db.NewQueryUser()

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

	//	if rows, err := v.GetQuery(); err != nil {
	//		pp(err)
	//	} else {
	//		plib.Dump(rows)
	//	}
}

func main() {
	defer session.SafeClose()

	//	t1()
	//	t2()
	t3()
	//	t4()
	//	t5()
}
