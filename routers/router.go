package routers

import (
	"rcdd/project/VSTH/ctrls"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/index.html", &ctrls.MainController{}, "*:Index")
	//beego.Router("/", &ctrls.MainController{}, "*:Login")
	beego.Router("/", &ctrls.MainController{}, "*:Index")
	beego.Router("/login.html", &ctrls.MainController{}, "*:Login")
	beego.Router("/register.html", &ctrls.MainController{}, "*:Register")

	beego.Router("/account/loginOrNot", &ctrls.AccountController{}, "*:LoginOrNot")
	beego.Router("/account/login", &ctrls.AccountController{}, "*:Login")
	beego.Router("/account/logout", &ctrls.AccountController{}, "*:Logout")
	beego.Router("/account/register", &ctrls.AccountController{}, "*:Register")
	beego.Router("/account/GetPassword", &ctrls.AccountController{}, "*:GetPassword")

	beego.Router("/Myjob.html", &ctrls.MainController{}, "*:Job")
	beego.Router("/sbvs_result.html", &ctrls.MainController{}, "*:SbvsResult")
	beego.Router("/lbvs_result.html", &ctrls.MainController{}, "*:LbvsResult")
	beego.Router("/sbvs.html", &ctrls.MainController{}, "*:SBVS")
	beego.Router("/lbvs.html", &ctrls.MainController{}, "*:LBVS")
	beego.Router("/myLibrary.html", &ctrls.MainController{}, "*:MyLibrary")
	beego.Router("/sysLibrary.html", &ctrls.MainController{}, "*:SysLibrary")
	beego.Router("/GetPassword.html", &ctrls.MainController{}, "*:GetPassword")

	beego.Router("/Myjob", &ctrls.JobController{}, "*:MyJobs")
	beego.Router("/GetQuery", &ctrls.JobController{}, "*:GetLigandOfTarget")
	beego.Router("/jobSubmit", &ctrls.JobController{}, "*:JobSubmit")
	beego.Router("/jobUpdate", &ctrls.JobController{}, "*:JobUpdate")
	beego.Router("/jobResult", &ctrls.JobController{}, "*:JobResults")
	beego.Router("/downloadJobResult?*", &ctrls.JobController{}, "*:DownloadJobResults")
	beego.Router("/downloadJobResult", &ctrls.JobController{}, "*:DownloadJobResults")
	beego.Router("/getZincProp", &ctrls.JobController{}, "*:GetPropsOfZincMol")
	beego.Router("/getZincSupply", &ctrls.JobController{}, "*:GetSupplysOfZincMol")

	beego.Router("/SysLib", &ctrls.JobController{}, "*:SysLibs")
	beego.Router("/MyLib", &ctrls.JobController{}, "*:MyLibs")
	beego.Router("/UploadLib", &ctrls.JobController{}, "*:UploadLib")
	beego.Router("/CheckLibNameExist", &ctrls.JobController{}, "*:CheckLibNameExist")
}
