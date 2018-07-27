package main

import (
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"

	_ "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"gpc/db/mongodb/orm/structure"

	plib "gpc/publib"
)

var (
	pp  = fmt.Println
	ppp = func() { pp("######################################") }
	_   = plib.Debug
	_   = spew.Config
)

type App struct {
	Id         bson.ObjectId `bson:"_id" json:"id"`
	Name       string
	Remark     string
	CreateTime time.Time
	Status     uint8 // 状态值(0预留):1为正常,2为禁用
	Arr        []int
}

func NewApp(name string) *App {
	return &App{
		Id:         bson.NewObjectId(),
		Name:       name,
		CreateTime: time.Now(),
	}
}

func (self *App) CollName() string {
	return "app_name"
}

func t1() {
	//	app := NewApp("FirstApp")
	//	app = nil

	//	var app interface{}
	app := *NewApp("第一个App")

	if info, err := structure.Ref(app); err != nil {
		pp(err)
	} else {
		plib.Dump(info)
	}
}

func main() {
	t1()
}
