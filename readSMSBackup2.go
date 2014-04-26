package main

import (
	"fmt"
	"XMLParseCon"
)

func main(){
	data := XMLParseCon.GetDataFieldFromFile("smsBackup.xml", "body")
	for _, m :=  range data{
		fmt.Println(m)
	}
}
