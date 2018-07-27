package main

import (
	"fmt"
	"rcdd/project/VSTH/Common/cmd"

	"github.com/astaxie/beego"
)

func testRemoteCall() {
	idFile := beego.AppConfig.String("idFile")
	ipAndPort := beego.AppConfig.String("ipAndPort")
	username := beego.AppConfig.String("username")
	script := beego.AppConfig.String("RemoteScript")

	if err := cmd.RemoteCallScript(idFile, username, ipAndPort, script); err != nil {
		fmt.Println(err)
	}

}

func testRemoteVinaCall() {
	idFile := beego.AppConfig.String("idFile")
	ipAndPort := beego.AppConfig.String("ipAndPort")
	username := beego.AppConfig.String("username")
	script := beego.AppConfig.String("VinaScript")

	jobId := "0011"
	pdbCode := "5f5w"
	library := "zinc_test"
	nodes := 10

	err := cmd.RemoteCallVinaScript(idFile, username, ipAndPort, script, jobId, pdbCode, library, nodes)
	if err != nil {
		fmt.Println(err)
	}
}

func testRemoteJobResInsertCall() {
	idFile := beego.AppConfig.String("idFile")
	ipAndPort := beego.AppConfig.String("ipAndPort")
	username := beego.AppConfig.String("username")
	script := beego.AppConfig.String("JobResInsertScript")

	jobId := "0011"
	library := "zinc_test"

	err := cmd.RemoteCallJobResInsertScript(idFile, username, ipAndPort, script, jobId, library)
	if err != nil {
		fmt.Println(err)
	}
}

func testRemoteJobUpdateCall() {
	idFile := beego.AppConfig.String("idFile")
	ipAndPort := beego.AppConfig.String("ipAndPort")
	username := beego.AppConfig.String("username")
	jobUpdateScript := beego.AppConfig.String("JobUpdateScript")
	jobResInsertScript := beego.AppConfig.String("JobResInsertScript")

	jobId := "7d03560b_eab1_4905_bd5a_9ebe8970890c"
	library := "zinc_test"

	err := cmd.RemoteCallJobUpdateScript(idFile, username, ipAndPort, jobUpdateScript, jobResInsertScript, jobId, library)
	if err != nil {
		fmt.Println(err)
	}
}

func testCallScp() {
	idFile := beego.AppConfig.String("idFile")
	ipAndPort := beego.AppConfig.String("ipAndPort")
	username := beego.AppConfig.String("username")
	fromFile := "/home/yanx/tmp/test.sh"
	toFileFile := "/HOME/nscc-gz_vscreening_1/WORKSPACE/"

	err := cmd.CallScp(idFile, username, ipAndPort, fromFile, toFileFile)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	//	testRemoteCall()
	//testRemoteVinaCall()
	//testRemoteJobResInsertCall()
	testRemoteJobUpdateCall()

}
