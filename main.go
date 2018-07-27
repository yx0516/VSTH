package main

import (
	_ "rcdd/project/VSTH/routers"

	"github.com/astaxie/beego"

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
	c.DB = "vs"
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

func main() {
	beego.Run()
}
