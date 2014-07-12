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
	err := smtp.SendMail(
		"smtp.gmail.com:583",
		auth,
		"richy.sting@gmail.com",
		[]string{"richy.sting@googlemail.com"},
		[]byte{"test test"})

	if err != nil {
		fmt.Println(err)
	}
}
