package help

import (
	"crypto/tls"
	"im/define"
	"net/smtp"

	"github.com/jordan-wright/email"
)

func SendEmailCode(mail, code string) error {
	e := email.NewEmail()
	e.From = "vczs <vczsvs@163.com>"
	e.To = []string{mail}
	e.Subject = "vczs平台验证码"
	e.HTML = []byte("你的验证码为：<h1>" + code + "</h1>")
	err := e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("", "vczsvs@163.com", define.MainPwd, "smtp.163.com"), &tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	if err != nil {
		return err
	}
	return nil
}
