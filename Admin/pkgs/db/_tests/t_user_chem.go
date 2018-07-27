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
	chem := db.NewUserChem("db02")
	chem.UserId = 1
	//	chem.Name = ""

	pp(chem.Insert())

	plib.Dump(chem)
}

func t2() {
	chem := db.NewUserChem("db01")
	chem.UserId = 1
	//	chem.Id = 1

	pp(chem.Read())
	//	pp(chem.ReadByName())
	//	pp(chem.ReadById())

	plib.Dump(chem)
}

func t3() {
	chem := db.NewUserChem("db0c")
	//	chem.UserId = 2
	chem.Id = 1

	chem.Name = "db03"
	chem.Remark = "测试内容111"
	chem.Status = 0

	//	pp(chem.Update())
	pp(chem.UpdateByFields([]string{"name", "remark"}))
}

func t4() {
	chem := db.NewUserChem("db02")
	chem.UserId = 3

	pp(chem.DeleteByName())
}

func t5() {
	for i := 0; i < 20; i++ {
		//		chem := db.NewUserChem(randStrNum(10))
		chem := db.NewUserChem(fmt.Sprintf("chem%v", i))
		chem.UserId = 1
		//	chem.Name = ""

		pp(chem.Insert())
	}
}

func t6() {
	v := db.NewQueryUserChem()

	//	v.Limit = 2
	v.Sort = []string{"-name"}
	//	v.Match.Gt(false, "id", 1).Add("name", "db02")
	v.Match.Regex("name", "db")
	v.Fields.True("id", "name")

	pp(v.Count())
	pp(v.CountAll())

	//	if rows, err := v.QueryData(); err != nil {
	//		pp(err)
	//	} else {
	//		plib.Dump(rows)
	//	}

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
