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

type Status int

func (self *Status) SetON() {
	*self = 1
}

func (self *Status) SetOff() {
	*self = 2
}

// 用户
type User struct {
	//	Id         bson.ObjectId `bson:"_id" json:"id"`
	Name       string
	Status     Status // 状态值(0预留):1为正常,2为禁用
	CreateTime time.Time
	UpdateTime time.Time
}

func NewUser(name string) *User {
	user := &User{
		// Id:         bson.NewObjectId(), // 不初始化
		Name:       name,
		CreateTime: time.Now(),
	}

	return user
}

func t1() {
	user := NewUser("user1")
	user.Status.SetON()
	//	user.Id.New()

	pp(fmt.Sprintf("%v", user.Status))

	//	pp(user.Id.NewByHex("581e83f099c05b1464c1a791"))
	err := database.C("user").Insert(user)
	if err != nil {
		pp(err)
	}

	plib.Dump(user)
}

func t2() {
	//	user := new(User)
	user := new(interface{})
	err := database.C("user").Find(bson.M{"name": "user1"}).One(user)
	if err != nil {
		pp(err)
	} else {
		plib.Dump(user)
	}
}

func main() {
	defer session.SafeClose()

	//	t1()
	t2()

}
