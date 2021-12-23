package main

import (
	"fmt"
	"github.com/jordan-wright/email"
	"log"
	"net/smtp"
)

func main() {
	sendEmail("wonkung@163.com")
}

const (
	SMTPHost     = "smtp.163.com"
	SMTPPort     = ":25"
	SMTPUsername = "pzxy_test@163.com"
	SMTPPassword = "PGQKSSRJELESXLFP"
)

//func sendEmail() {
//	newEmail := email.NewQQExmail("wonkung@163.com", "125800")
//	text := newEmail.Info("标题", "作者", []string{"1054470041@qq.com"}).
//		SendText("文本内容")
//	log.Println(text)
//}

//	"github.com/jordan-wright/email"
func sendEmail(receiver string) {
	auth := smtp.PlainAuth("", SMTPUsername, SMTPPassword, SMTPHost)
	e := &email.Email{
		From:    fmt.Sprintf("发送者名字<%s>", SMTPUsername),
		To:      []string{receiver},
		Subject: "这里是标题内容",
		Text:    []byte("这里是正文内容"),
	}
	err := e.Send(SMTPHost+SMTPPort, auth)
	if err != nil {
		log.Fatal(err)
	}
}
