package main

import (
	"fmt"
	"strconv"
)

type MyT struct {
	first int64
	second int64
}

func (d *MyT) change(n int64){
	d.first = n;
}

func (d *MyT) changeS(n string){
	d.first, _ = strconv.ParseInt(n, 10, 64);
}
type Int int;

func main(){
	fmt.Printf("in main\n");
	t := MyT{12, 13};
	fmt.Println(t);
	fmt.Println(t.first);
	//n := Int(12);
	//fmt.Println(n);
	t.changeS("999");
	fmt.Println(t);
}
