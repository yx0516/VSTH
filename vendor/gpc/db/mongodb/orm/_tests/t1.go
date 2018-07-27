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
	"gpc/db/mongodb/conn/gsession"
	"gpc/db/mongodb/orm"

	plib "gpc/publib"
)

var (
	pp  = fmt.Println
	ppp = func() { pp("######################################") }
	_   = plib.Debug
	_   = spew.Config
	_   = os.Args

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
		"192.168.200.35",
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

type App struct {
	Id         bson.ObjectId `bson:"_id" json:"id"`
	Name       string
	Remark     string
	CreateTime time.Time
	Status     uint8 // 状态值(0预留):1为正常,2为禁用
	Arr        []int
}

func (self App) CollName() string {
	return "app"
}

func NewApp(name string) *App {
	return &App{
		Id:         bson.NewObjectId(),
		Name:       name,
		CreateTime: time.Now(),
	}
}

func t1() {
	app := NewApp("app12")
	app.Remark = "测试备注"
	app.Id = ""

	if err := orm.Insert(app); err != nil {
		log.Fatal(err)
	}

	//	arr := []*App{NewApp("app3"), NewApp("app4"), NewApp("app5")}
	//	arr := []App{*NewApp("app3"), *NewApp("app4"), *NewApp("app5")}
	//	//	arr := []string{"a1", "a2"}
	//	if err := orm.Insert(arr); err != nil {
	//		log.Fatal(err)
	//	}
}

func t2() {
	app := NewApp("")
	//app.Id = bson.ObjectIdHex("57d8f5c36c91dd2f648012f5")

	app.Name = "app3"
	//	app.Status = 2

	pp(orm.ReadId(app))
	plib.Dump(app)

	return

	//	read := orm.NewRead(app)

	var apps []App
	read := orm.NewRead(&apps)

	read.SetMatchFields("name")
	read.SetMatchFields("status")

	//	read.SetMatchM(bson.M{"name": "app5", "status": 0})
	read.SetMatchM(bson.M{"status": 0})

	err := read.Query()
	//	err := read.Query("name", "createtime")

	if err != nil {
		log.Fatal(err)
	}

	//	plib.Dump(app)
	plib.Dump(apps)
}

func t3() {
	app := NewApp("")
	app.Id = bson.ObjectIdHex("57d909c16c91dd2c80ce69fa")

	app.Name = "app3"
	app.Status = 0

	//	pp(orm.DeleteId(app))

	del := orm.NewDelete(app)
	//	del.SetMatchFields("name")
	//	del.SetMatchFields("status")
	//	del.SetMatchM(bson.M{"status": 0})

	//	del.IsAll = true

	if info, err := del.Remove(); err != nil {
		pp(err == mgo.ErrNotFound)
		log.Fatal(err)
	} else {
		plib.Dump(info)
	}
}

func t4() {
	app := NewApp("")
	//	app.Id = bson.ObjectIdHex("57d91c176c91dd2c1063f2b9")
	app.Name = "app3112"
	app.Remark = "1测试备注1234"
	app.Status = 2

	up := orm.NewUpdate(app)
	//	up.IsAll = true
	//	up.IsUpsert = true
	up.IsReplace = true

	up.SetFields("name")
	//	up.SetFields("status")
	//	up.SetFields("remark")
	//	up.Set.M = bson.M{"name": "app2"}

	up.Match = bson.M{"name": "app3112"}

	if info, err := up.Update(); err != nil {
		pp(err)
	} else {
		plib.Dump(info)
	}
}

func main() {
	t1()
	//	t2()
	//	t3()
	//	t4()
}
