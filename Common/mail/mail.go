package mail

import (
	//"fmt"
	//"os/exec"
	//"strconv"
	//	"strings"
	"gopkg.in/gomail.v2"
)

/**
QQ 邮箱
POP3 服务器地址：qq.com（端口：995）
SMTP 服务器地址：smtp.qq.com（端口：465/587）

163 邮箱：
POP3 服务器地址：pop.163.com（端口：110）
SMTP 服务器地址：smtp.163.com（端口：25）

126 邮箱：
POP3 服务器地址：pop.126.com（端口：110）
SMTP 服务器地址：smtp.126.com（端口：25）
**/

func SentMailBy163(from string, to string, subject string, body string, authority string) error {

	m := gomail.NewMessage()

	m.SetHeader("From", from)
	m.SetHeader("To", to) //send email to multipul persons

	m.SetHeader("Subject", subject)

	m.SetBody("text/html", body)
	d := gomail.NewPlainDialer("smtp.163.com", 25, from, authority)
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

func SentMailBy163FromVSTH(to string, subject string, body string) error {

	from := "VSTH_RCDD@163.com"
	auth := "rcddvsth2018"

	err := SentMailBy163(from, to, subject, body, auth)
	return err
}

func SentMailByQQ(from string, to string, subject string, body string, authority string) error {

	m := gomail.NewMessage()

	m.SetHeader("From", from)
	m.SetHeader("To", to) //send email to multipul persons

	m.SetHeader("Subject", subject)

	m.SetBody("text/html", body)
	d := gomail.NewPlainDialer("smtp.qq.com", 465, from, authority)
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
