package email

import (
	"gopkg.in/gomail.v2"
	"strconv"
)

func SendMail(mailTo []string, subject string, body string) error {
	// 设置邮箱主体
	mailConn := map[string]string{
		"user": "q2506152074@163.com",
		"pass": "TPDAYYESMYRHLSDU",
		"host": "smtp.163.com",
		"port": "465",
	}

	port, _ := strconv.Atoi(mailConn["port"])
	m := gomail.NewMessage()
	// 添加别名
	m.SetHeader("From", m.FormatAddress(mailConn["user"], "发验证码滴!!!"))
	// 发送给用户(可以多个)
	m.SetHeader("To", mailTo...)
	// 设置邮件主题
	m.SetHeader("Subject", subject)
	// 设置邮件正文
	m.SetBody("text/html", body)
	// 设置邮件正文
	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])
	err := d.DialAndSend(m)
	return err
}
