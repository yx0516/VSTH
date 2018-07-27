package main

import (
	"fmt"
	"rcdd/project/VSTH/Common/mail"

	//"github.com/astaxie/beego"
)

func testSendMailBy163() {
	username := "test"
	password := "test"

	from := "VSTH_RCDD@163.com"
	to := "280878424@qq.com"
	subject := "User information of VSTH"
	body := "Hi " + username + ",<br><br>" + "Thanks for registering VSTH, your user information is: <br>" +
		"username: " + username + "<br>" +
		"password: " + password + "<br><br>" +
		"Bests, <br>VSTH"
	auth := "rcddvsth2018"

	if err := mail.SentMailBy163(from, to, subject, body, auth); err != nil {
		fmt.Println(err)
	}

}

func main() {

	testSendMailBy163()

}
