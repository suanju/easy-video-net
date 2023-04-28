package email

import (
	"easy-video-net/global"
	"gopkg.in/gomail.v2"
	"strconv"
)

func SendMail(mailTo []string, subject string, body string) error {
	// 设置邮箱主体
	mailConn := map[string]string{
		"user": global.Config.EmailConfig.User,
		"pass": global.Config.EmailConfig.Pass,
		"host": global.Config.EmailConfig.Host,
		"port": global.Config.EmailConfig.Port,
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
	if err != nil {
		global.Logger.Errorf("发送至%s邮箱验证码失败,内容 ： %s 错误原因 :%d", mailTo, body, err)
	}
	return err
}
