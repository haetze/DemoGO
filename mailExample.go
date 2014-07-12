package main

import (
	"fmt"
	"net/smtp"
)

func main() {
	auth := smtp.PlainAuth(
		"",
		"richy.sting@gmail.com",
		"V3J2@Quitlam",
		"smtp.gmail.com")

	fmt.Println("aut set up")

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"richy.sting@gmail.com",
		[]string{"richy.sting@googlemail.com"},
		[]byte("test test"))

	fmt.Println("send")

	if err != nil {
		fmt.Println(err)
	}
}
