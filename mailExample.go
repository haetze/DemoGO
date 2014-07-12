package main

import (
	"fmt"
	"net/smtp"
	"strconv"
)

type User struct {
	name     string
	email    string
	password string
	server   string
	port     int
	auth     smtp.Auth
}

func main() {
	user := User{"", "me@gmail.com", "***********", "smtp.gmail.com", 587, nil}
	user.auth = gmailAuth(user)

	fmt.Println("aut set up")

	err := gmailSend(user, "you@gmail.com", "test test")
	fmt.Println("send")

	if err != nil {
		fmt.Println(err)
	}

}

func gmailAuth(user User) smtp.Auth {
	auth := smtp.PlainAuth(
		user.name,
		user.email,
		user.password,
		user.server)
	return auth
}

func gmailSend(user User, to, body string) error {
	err := smtp.SendMail(user.server+":"+strconv.Itoa(user.port), user.auth, user.email, []string{to}, []byte(body))
	return err
}
