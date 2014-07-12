package main

import (
	"fmt"
	"net/smtp"
)

func main() {
	auth := smtp.PlainAuth(
		"",
		"me@gmail.com",
		"****************",
		"smtp.gmail.com")

	fmt.Println("aut set up")

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"me@gmail.com",
		[]string{"you@gmail.com"},
		[]byte("test test"))

	fmt.Println("send")

	if err != nil {
		fmt.Println(err)
	}

}
