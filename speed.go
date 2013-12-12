package main

import(
	"fmt"
	"time"
)

func main(){
	t1 := time.Now().UnixNano();
	for i:=0; i<100000; i++{
		fmt.Println("done");
	}
	t2 := time.Now().UnixNano();
	fmt.Println((t2-t1)/1000000);

}
