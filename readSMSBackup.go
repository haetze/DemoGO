package main

import (
	"fmt"
	"XMLParse"
)

func main(){
	data := XMLParse.GetDataFieldFromFile("smsBackup.xml", "locked")
	for _, m :=  range data{
		fmt.Println(m)
	}
}
