package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	rs "gpc/util/randstr"

	"github.com/davecgh/go-spew/spew"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"gpc/db/mongodb/conn"

	plib "gpc/publib"
)

var (
	pp  = fmt.Println
	ppp = func() { pp("######################################") }
	_   = plib.Debug
	_   = spew.Config
	_   = os.Args
	_   = log.Fatal

	session  *conn.Session
	database *mgo.Database

	objRand    *rand.Rand
	randStrNum func(n int) string

	_ = make(bson.M)
)

func init() {
	//	mgo.SetDebug(true)
	//	mgo.SetLogger(log.New(os.Stdout, "[mgo] ", log.Ldate|log.Ltime))

	objRand = rand.New(rand.NewSource(time.Now().UnixNano())) // 初始化随机变量
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
}

func t1() {

}

func main() {
	defer session.SafeClose()

	t1()
}
