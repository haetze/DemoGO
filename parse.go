package main

import (
	"strings"
	"fmt"
	"io/ioutil"
)


func main(){
	bytesStr, _ := ioutil.ReadFile("smsBackup.xml")
	str := string(bytesStr)
	mes := strings.SplitAfter(str, "body=\"" )
	for n, m := range mes{
		if n!=1{
			b := strings.SplitAfter(m, "\"")
			b[0] = strings.Replace(b[0], "\"", " ", 1)
			fmt.Println(string(b[0]))
		}
	}


}

