package main

import (
	"fmt"
	"XMLParse"
)

func main(){
	data := XMLParse.FileToMapPreLine("smsBackup.xml")
	for _, m :=  range data{
		n := m["body"]
		if n == ""{
			fmt.Println("no field value")
		}else{
			fmt.Println(n)
		}
	}
}
