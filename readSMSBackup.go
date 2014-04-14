package main

import (
	"fmt"
	"XMLParse"
)

func main(){
	data := XMLParse.GetDataFieldFromFile("feeds.opml", "xmlUrl")
	for _, m :=  range data{
		fmt.Println(m)
	}
}
