package main 

import(
//	"fmt"
	"TCP"
)

func main(){
	TCP.Client(TCP.IP{127,0,0,1}, 8080)
}

