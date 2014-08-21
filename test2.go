package main

import "fmt"

func main(){
	test(1,2,1,1,1,1,"q")
}

func test(b ...interface{}){
	fmt.Println(b)
}

