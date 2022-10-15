package test

import (
	"crypto/tls"
	"github.com/jordan-wright/email"
	"net/smtp"
	"testing"
)

func TestSendEmail(t *testing.T) {
	e := email.NewEmail()
	e.From = "fly321 <944219401@qq.com>"
	e.To = []string{"linqi0349486@163.com"}
	e.Subject = "主题：验证码发送测试"
	//e.Text = []byte("Text Body is, of course, supported!")
	e.HTML = []byte("<h1>我是测试验证码</h1>")
	//err := e.Send("smtp.qq.com:465", smtp.PlainAuth("", "944219401@qq.com", "密钥", "smtp.qq.com"))
	// eof 使用下面的tls
	err := e.SendWithTLS("smtp.qq.com:465",
		smtp.PlainAuth("", "944219401@qq.com", "密钥", "smtp.qq.com"),
		&tls.Config{
			InsecureSkipVerify: true,
			ServerName:         "smtp.qq.com",
		})
	if err != nil {
		t.Error(err.Error())
	}
}
