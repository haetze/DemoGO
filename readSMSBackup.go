package main

import (
	"fmt"
	"XMLParse"
)

func main(){
	data := XMLParse.GetDataFieldFromFile("smsBackup.xml", "body")
	/*for_, m :=  range data{
		fmt.Println(m)
	}*/
	fmt.Println(len(data))
}
