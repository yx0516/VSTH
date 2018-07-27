package ctrls

import (
	"fmt"
	"rcdd/project/VSTH/Admin/pkgs/db"
	"rcdd/project/VSTH/Common/mail"
	"strconv"
	"strings"
	"time"
)

type AccountController struct {
	baseController
}

func (this *AccountController) LoginOrNot() {
	data := make(map[string]interface{})
	data["userid"] = this.userid
	data["username"] = this.username
	this.responseMsg.SuccessMsg("", data)
	this.Data["json"] = this.responseMsg
	this.ServeJSON()
}

//登录
func (this *AccountController) Login() {
	fmt.Println("login")
	username := strings.TrimSpace(this.GetString("username"))
	fmt.Println("username:" + username)
	//email := strings.TrimSpace(this.GetString("useremail"))
	//fmt.Println("Email:" + email)
	password := strings.TrimSpace(this.GetString("password"))
	fmt.Println("Passwd:" + password)
	remember := this.GetString("remember")
	fmt.Println(remember)
	if username != "" && password != "" {
		user := db.NewUser(username)
		//if user.ReadByName() != nil || user.Password != Md5([]byte(password)) {
		if user.ReadByName() != nil || user.Password != (password) {
			fmt.Println(user.Password)
			//fmt.Println(Md5([]byte(password)))
			//fmt.Println("wrong email or password")
			this.responseMsg.ErrorMsg("error username or password", nil)
			this.Data["json"] = this.responseMsg
			this.ServeJSON()

		} else if user.Status == 2 {
			this.responseMsg.ErrorMsg("account is inactive", nil)
			this.Data["json"] = this.responseMsg
			this.ServeJSON()

		} else {
			//fmt.Println("ss")
			user.Logincount += 1
			user.Lastip = this.getClientIp()
			user.Lastlogin = time.Now()
			user.Update()
			authkey := Md5([]byte(this.getClientIp() + "|" + user.Password))
			if remember == "yes" {
				this.Ctx.SetCookie("auth", strconv.FormatInt(int64(user.Id), 10)+"|"+authkey, 7*86400)
			} else {
				this.Ctx.SetCookie("auth", strconv.FormatInt(int64(user.Id), 10)+"|"+authkey)
			}
			this.userid = user.Id
			fmt.Println("account login success")

			data := make(map[string]interface{})
			data["userid"] = user.Id
			data["username"] = user.Name
			this.responseMsg.SuccessMsg("", data)
			this.Data["json"] = this.responseMsg
			this.ServeJSON()

			//this.Redirect("/myJobs", 302)
			//this.TplName = "index.html"
			//return
		}

	}

}

//退出登录
func (this *AccountController) Logout() {
	this.Ctx.SetCookie("auth", "")
	this.userid = -1
	this.username = "anonymous"
	this.Redirect("/", 302)
}

//注册
func (this *AccountController) Register() {
	fmt.Println("register")
	username := strings.TrimSpace(this.GetString("username"))
	fmt.Println("username:" + username)
	email := strings.TrimSpace(this.GetString("email"))
	fmt.Println("Email:" + email)
	password := strings.TrimSpace(this.GetString("password"))
	fmt.Println("Passwd:" + password)

	if username != "" && email != "" && password != "" {

		userQuery := db.NewQueryUser()
		userQuery.Match.Add("name", username)
		if users, err := userQuery.QueryData(); err != nil {
			this.responseMsg.ErrorMsg("Error happens，Please register again.", nil)
			this.Data["json"] = this.responseMsg
			this.ServeJSON()
			return
		} else if len(users) > 0 {
			fmt.Println(users[0])
			this.responseMsg.ErrorMsg("Account has already existed，Please login.", nil)
			this.Data["json"] = this.responseMsg
			this.ServeJSON()
			return
		} else {
			user := db.NewUser(username)

			//check if already used
			//save user infor to database
			user.Status = 1
			//user.Name = username
			//user.Password = Md5([]byte(password))
			user.Password = password
			user.Email = email
			//user.Phone = "1234567890"
			_, err := user.Insert()
			if err != nil {
				fmt.Println(err)
				this.responseMsg.ErrorMsg("Error happens when saving user.", nil)
				this.Data["json"] = this.responseMsg
				this.ServeJSON()
				return
			} else {
				// send userinfo to user by email
				to := email
				subject := "User information of VSTH"
				body := "Hi " + username + ",<br><br>" + "Thanks for registering VSTH, your user information is: <br>" +
					"username: " + username + "<br>" +
					"password: " + password + "<br><br>" +
					"Bests, <br>VSTH"

				if err := mail.SentMailBy163FromVSTH(to, subject, body); err != nil {
					fmt.Println(err)
					this.responseMsg.SuccessMsg("Register successfully, please login.", nil)
					this.Data["json"] = this.responseMsg
					this.ServeJSON()
					return
				}
				this.responseMsg.SuccessMsg("Register successfully, please login. And an email with user information has been sent to you.", nil)
				this.Data["json"] = this.responseMsg
				this.ServeJSON()
				return
			}
		}

	} else {
		this.responseMsg.ErrorMsg("username, password, email can not be empty！", nil)
		this.Data["json"] = this.responseMsg
		this.ServeJSON()
		return
	}

}

// getPassword
func (this *AccountController) GetPassword() {
	fmt.Println("Get password")
	username := strings.TrimSpace(this.GetString("username"))
	fmt.Println("username:" + username)
	email := strings.TrimSpace(this.GetString("useremail"))
	fmt.Println("Email:" + email)

	if username != "" && email != "" {
		user := db.NewUser(username)

		if user.ReadByName() != nil || user.Email != email {
			fmt.Println(user.Email)
			//fmt.Println("Send password to user")
			// send password to user
			to := email
			subject := "Forgot password"
			body := "Hi " + username + ",<br><br>" + "Your password of VSTH is: " +
				user.Password + "<br><br>" +
				"Bests, <br>VSTH"

			if err := mail.SentMailBy163FromVSTH(to, subject, body); err != nil {
				fmt.Println(err)
				this.responseMsg.ErrorMsg("Error username or email", nil)
				this.Data["json"] = this.responseMsg
				this.ServeJSON()
				return

			} else {
				this.responseMsg.SuccessMsg("An email with user information has been sent to you.", nil)
				this.Data["json"] = this.responseMsg
				this.ServeJSON()
				return
			}

		} else {
			this.responseMsg.ErrorMsg("Error username or email", nil)
			this.Data["json"] = this.responseMsg
			this.ServeJSON()
			return
		}
	} else {
		this.responseMsg.ErrorMsg("username and email can not be empty", nil)
		this.Data["json"] = this.responseMsg
		this.ServeJSON()
		return
	}
}
