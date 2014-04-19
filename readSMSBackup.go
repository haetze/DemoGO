package main

import (
	"fmt"
	"XMLParse"
)

func main(){
	data := XMLParse.FileToMapPreLine("smsBackup.xml")
	for _, m :=  range data{
		fmt.Println(m["body"])
	}
}
