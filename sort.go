package main 

import (
	"fmt"
	"sort"
	"os"
	"strconv"
)
func main (){
	var y []int
	if len(os.Args) == 1{
		y = []int{12,43,1,6,333}
	}else{
		for _, m := range os.Args{
			i, err := strconv.Atoi(m)
			if err != nil {
			}else{
				y = append(y, i)
			}
		}
	}
	sort.Ints(y)
	fmt.Println(y)
}
