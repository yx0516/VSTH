package ctrls

import (
	"rcdd/project/VSTH/Admin/pkgs/db"
	"strconv"
	"strings"
	//	"time"
	"crypto/md5"
	"fmt"

	"github.com/astaxie/beego"
)

type ResponseMsg struct {
	State int // 0:成功,-1：失败
	Msg   string
	Data  interface{}
}

func (this *ResponseMsg) ErrorMsg(msg string, data interface{}) {
	this.State = -1
	this.Msg = msg
	this.Data = data
}

func (this *ResponseMsg) SuccessMsg(msg string, data interface{}) {
	this.State = 0
	this.Msg = msg
	this.Data = data
}

func Md5(buf []byte) string {
	hash := md5.New()
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

type baseController struct {
	beego.Controller
	userid          int
	username        string
	UserTmpFilePath string
	responseMsg     ResponseMsg

	moduleName     string
	controllerName string
	actionName     string
}

func (this *baseController) Prepare() {

	//this.moduleName = "admin" //需要管理员时使用
	controllerName, actionName := this.GetControllerAndAction()
	this.moduleName = ""
	this.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	this.actionName = strings.ToLower(actionName)

	this.responseMsg = *new(ResponseMsg)
	this.auth() // 用户是否登录验证

	this.UserTmpFilePath = this.getUserTmpFilePath()
	//this.checkPermission()

}

//登录状态验证
func (this *baseController) auth() {
	if this.controllerName == "account" && (this.actionName == "login" || this.actionName == "logout") {

	} else {
		arr := strings.Split(this.Ctx.GetCookie("auth"), "|")
		if len(arr) == 2 {
			idstr, password := arr[0], arr[1]
			userid, _ := strconv.ParseInt(idstr, 10, 0)
			if userid > 0 {
				user := db.NewUser("")
				user.Id = int(userid)
				if user.ReadById() == nil && password == Md5([]byte(this.getClientIp()+"|"+user.Password)) {
					this.userid = user.Id
					this.username = user.Name
					fmt.Println("anth login successfully.")
				}
			}
		}
		if this.userid == 0 {
			fmt.Println("user id error, login with anonymous")
			//this.Redirect("/login.html", 302)
			this.userid = -1
			this.username = "anonymous"

			//			return
		}
	}
}

//获取用户IP地址
func (this *baseController) getClientIp() string {
	s := strings.Split(this.Ctx.Request.RemoteAddr, ":")
	return s[0]
}

func (this *baseController) getUserTmpFilePath() string {
	userTmpFilePath := strings.TrimSpace(beego.AppConfig.String("tmpFilePath")) + "/" + this.username
	return userTmpFilePath
}
