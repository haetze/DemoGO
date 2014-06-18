package main

import (
	"fmt"
	"XMLParseCon"
)

func main(){
	data := XMLParseCon.GetFileToMapPreLine("smsBackup.xml")
	for _, m :=  range data{
		fmt.Println(m.Content["body"])
	}
}
