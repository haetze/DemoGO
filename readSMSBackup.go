package main

import (
	"fmt"
	"XMLParse"
)

func main(){
	data := XMLParse.FileToMapPreLine("java.xml")
	for _, m :=  range data{
		n := m["android:minSdkVersion"]
		if n == ""{
			fmt.Println("no field value")
		}else{
			fmt.Println(n)
		}
	}
}
