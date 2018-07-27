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
	//	"gpc/db/mongodb/orm"
	"gpc/db/mongodb/xorm"

	plib "gpc/publib"
)

var (
	pp  = fmt.Println
	ppp = func() { pp("######################################") }
	_   = plib.Debug
	_   = spew.Config
	_   = os.Args
	_   = log.Ldate

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

	// 初始化 orm session
	gsession.InitSession(session.Session, c.DB)

	model := xorm.NewModelInit()
	model.AddDefModels(
		new(App),
	)
	model.Init("test")
}

type App struct {
	xorm.ModelPublic `bson:"-" json:"-"`
	Id               int
	Name             string
	Remark           string
	CreateTime       time.Time
	Status           uint8 // 状态值(0预留):1为正常,2为禁用
	Arr              []int
}

func (self App) CollName() string {
	return "app"
}

func (self *App) Init() {
	self.ModelPublic.Init(self)

	self.Model.FuncMakeMatch = func() (bson.M, error) {
		return bson.M{"name": "app2"}, nil
	}
}

func NewApp(name string) *App {
	app := &App{
		Name:       name,
		CreateTime: time.Now(),
	}
	app.Init()

	return app
}

func t1() {
	app := NewApp("app2")
	app.Status = 2
	app.Remark = "测试备注"
	app.Id = 10

	//	pp(app.Model.CheckObject())

	//	if v, err := app.Model.Copy(); err != nil {
	//		log.Fatal(err)
	//	} else {
	//		app.Name = "x111"
	//		pp(v.Insert())
	//		plib.Dump(v)
	//	}

	//	pp(app.Model.GetCollName())

	//	if v, err := app.Model.MakeModelPtr(); err != nil {
	//		log.Fatal(err)
	//	} else {
	//		plib.Dump(v)
	//	}

	//	if v, err := app.Model.GetStructType(); err != nil {
	//		log.Fatal(err)
	//	} else {
	//		pp(v)
	//	}

	//	pp(app.Model.GetStructFullName())

	//	if m, err := app.Model.GetFieldValues(); err != nil {
	//		log.Fatal(err)
	//	} else {
	//		plib.Dump(m)
	//	}

	//	if m, err := app.Model.GetFieldValues("name", "id", "remark"); err != nil {
	//		log.Fatal(err)
	//	} else {
	//		plib.Dump(m)
	//	}

	//	if id, err := app.Model.GetFieldIdValue(app); err != nil {
	//		log.Fatal(err)
	//	} else {
	//		pp(id)
	//	}

	//	pp(app.Model.CheckId(app.Id))

	//	if id, err := app.Model.VerifyId(); err != nil {
	//		log.Fatal(err)
	//	} else {
	//		pp(id)
	//	}

	//	if name, err := app.Model.GetFieldNameValue(app); err != nil {
	//		log.Fatal(err)
	//	} else {
	//		pp(name)
	//	}

	//	pp(app.Id)
	//	if err := app.Model.SetFieldIdValue(app, 3); err != nil {
	//		log.Fatal(err)
	//	} else {
	//		pp(app.Id)
	//	}

	//	pp(app.Name)
	//	if err := app.Model.SetFieldNameValue(app, "xxx111"); err != nil {
	//		log.Fatal(err)
	//	} else {
	//		pp(app.Name)
	//	}

	//	pp(app.Model.Exist(bson.M{"name": "app21"}))
	//	pp(app.Model.Exist(nil))

	//	pp(app.Model.ExistValue("name", "app21"))
	//	pp(app.Model.ExistFieldValue("name", "remark"))
	//	pp(app.Model.ExistName("app2", nil))
	//	pp(app.Model.ExistId())

	//	if v, err := app.Model.CopyRead(); err != nil {
	//		//	if v, err := app.Model.CopyRead("remark"); err != nil {
	//		log.Fatal(err)
	//	} else {
	//		plib.Dump(v)
	//	}

	//	if err := app.Model.ReadById(10, "Remark"); err != nil {
	//		log.Fatal(err)
	//	} else {
	//		plib.Dump(app)
	//	}

	//	if err := app.Model.ReadByName("app2", nil, "Ramrk", "arr"); err != nil {
	//		log.Fatal(err)
	//	} else {
	//		plib.Dump(app)
	//	}

	//	app.Name = "Ew2Nxdoa05"
	//	app.Status = 0
	//	if err := app.Model.ReadByFields("name", "status"); err != nil {
	//		log.Fatal(err)
	//	} else {
	//		pp(app.Model.CheckObject())
	//		plib.Dump(app)
	//	}

	//	if err := app.Model.Read(bson.M{"name": "Ew2Nxdoa05"}, "name", "status"); err != nil {
	//		log.Fatal(err)
	//	} else {
	//		pp(app.Model.CheckObject())
	//		plib.Dump(app)
	//	}

	//	var v []interface{}
	//	if err := app.Model.Query("app", bson.M{"name": "Ew2Nxdoa05"}, &v, "name", "status"); err != nil {
	//		log.Fatal(err)
	//	} else {
	//		plib.Dump(v)
	//	}

	//	pp(app.Model.UpdateById(app.Id, bson.M{"remark": "测试111", "Status": 9}))

	//	app.Remark = "更新测试21"
	//	app.Status = 12
	//	pp(app.Model.UpdateFieldsById(app.Id, "remark", "Status"))
	//	pp(app.Model.UpdateFields(bson.M{"id": app.Id}, "remark", "Status"))

	//	pp(app.Model.Update(bson.M{"id": app.Id}, bson.M{"remark": "测试a", "status": 8}))

	//	pp(app.Id)
	//	pp(app.Model.DeleteById(app.Id))
	//	pp(app.Model.DeleteByName("9lT0CZEkNR", nil))

	//	app.Name = "a8NIDGhq0X"
	//	app.Status = 0
	//	pp(app.Model.DeleteByFields("name", "Status"))

	ids := []int{13, 14}

	//	var v []interface{}
	//	if err := app.Model.ReadInIds("app", &v, ids...); err != nil {
	//		log.Fatal(err)
	//	} else {
	//		plib.Dump(v)
	//	}

	if info, err := app.Model.DeleteInIds("app", ids...); err != nil {
		log.Fatal(err)
	} else {
		plib.Dump(info)
	}
}

func t2() {
	app := NewApp("app2")
	app.Remark = "测试备注"
	app.Id = 15
	//	app.Id = bson.ObjectIdHex("57e0b29d6c91dd2b1004a6db")
	//	app.Id = bson.ObjectIdHex("57e0b29d6c91dd2b1004a6d1")

	//	app.DisableInsert(fmt.Errorf("禁止插入！"))

	//	if err := app.Insert(); err != nil {
	//		pp(app.StringPretty())
	//		log.Fatal(err)
	//	}

	pp(app.ExistId())

	//	if v, err := app.CopyRead("remark"); err != nil {
	//	if v, err := app.CopyRead(); err != nil {
	//		log.Fatal(err)
	//	} else {
	//		plib.Dump(v)
	//	}

	//	if err := app.Read("Remark"); err != nil {
	//		log.Fatal(err)
	//	}

	//	if err := app.ReadById("Remark"); err != nil {
	//		log.Fatal(err)
	//	}

	//	for i := 0; i < 10; i++ {
	//		app.Name = randStrNum(10)
	//		pp(app.Insert())
	//	}

	pp(app.StringPretty())
}

func main() {
	//	t1()
	t2()
}
