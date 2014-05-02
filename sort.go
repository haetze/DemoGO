package main 

import (
	"fmt"
	"sort"
)
func main (){
	y := []int{12,43,1,6,333}
	sort.Ints(y)
	fmt.Println(y)
}
