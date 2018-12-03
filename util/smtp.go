package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/smtp"
	"strings"
)

//SendMail send active email
func SendMail(uaddr string) bool {
	addr := "smtp.qq.com:25"
	auth := smtp.PlainAuth(
		"",
		"uchenily@qq.com",
		"mjyqjgywvpzdcach",
		"smtp.qq.com")
	from := "uchenily@qq.com"
	to := []string{uaddr}
	nickname := "auth"
	subject := "utronshop.io 邮箱验证"

	hasher := md5.New()
	hasher.Write([]byte(strings.Join(to, ",")))
	emailToken := hex.EncodeToString(hasher.Sum(nil))

	contentType := "Content-Type: text/html; charset=UTF-8"
	body := "<h1><a href='http://localhost:8080/activemail?emailToken=" + emailToken + "&addr=" + uaddr + "'>现在激活" + "</a></h1><p>如果以上超连接无法访问，请将以下网址复制到浏览器地址栏中</p><h3>http://localhost:8080/activemail?emailToken=" + emailToken + "&addr=" + uaddr + "</h3>"
	msg := []byte(
		"To: " + strings.Join(to, ",") + "\r\n" +
			"From: " + nickname + "<" + from + ">\r\n" +
			"Subject: " + subject + "\r\n" +
			contentType + "\r\n\r\n" +
			body)
	err := smtp.SendMail(addr, auth, from, to, msg)
	if err != nil {
		fmt.Printf("send mail error: %v", err)
		return false
	}
	return true
}
