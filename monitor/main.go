package main

import (
	"log"
	"net/smtp"
)

func main() {
	sendNeteaseTextMail("zhaobao_monitor@163.com",
		"2fc3baa90",
		"zhaobao@palmax.com",
		"掐指一算：越南的offer又打不开了",
		"已经一个多小时没有转化了")
}

func sendNeteaseTextMail(from, pass, to, subject, body string) {
	auth := smtp.PlainAuth(
		"",
		from,
		pass,
		"smtp.163.com",
	)
	err := smtp.SendMail(
		"smtp.163.com:25",
		auth,
		from,
		[]string{to},
		[]byte("To: "+to+"\r\n"+"Subject:"+subject+"\r\n"+"\r\n"+body+"\r\n"),
	)
	if err != nil {
		log.Fatal(err)
	}
}
