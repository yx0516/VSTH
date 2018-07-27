package ctrls

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (this *MainController) Index() {
	this.TplName = "index.html"
}

func (this *MainController) Login() {
	this.TplName = "account_login.html"
}

func (this *MainController) Marvin4js() {
	this.Ctx.Output.Download("marvin4js-license.cxl")
}
func (this *MainController) Register() {
	this.TplName = "account_signup.html"
}

func (this *MainController) Job() {
	this.TplName = "myJob.html"
}

func (this *MainController) SbvsResult() {
	this.TplName = "sbvs_result.html"
}

func (this *MainController) LbvsResult() {
	this.TplName = "lbvs_result.html"
}

func (this *MainController) SBVS() {
	this.TplName = "sbvs.html"
}

func (this *MainController) LBVS() {
	this.TplName = "lbvs.html"
}

func (this *MainController) MyLibrary() {
	this.TplName = "myLibrary.html"
}

func (this *MainController) SysLibrary() {
	this.TplName = "sysLibrary.html"
}

func (this *MainController) GetPassword() {
	this.TplName = "get_pwd.html"
}
