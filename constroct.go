package main

import (
	"fmt"
)

type MyT struct {
	first int;
	second int
}

func (d *MyT) change(n int){
	d.first = n;
}
type Int int;

func main(){
	fmt.Printf("in main\n");
	t := MyT{12, 13};
	fmt.Println(t);
	fmt.Println(t.first);
	//n := Int(12);
	//fmt.Println(n);
	t.change(99);
	fmt.Println(t);
}
