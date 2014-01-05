package main

import (
	"fmt"
)

type MyT struct {
	first int;
	second int
}

type Int int;

func main(){
	fmt.Printf("in main\n");
	t := MyT{12, 13};
	fmt.Println(t);
	fmt.Println(t.first);
	n := Int(12);
	fmt.Println(n);

}
