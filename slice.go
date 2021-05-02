package main


import (
	"fmt"
)

func main() {
	ar := [...]int{1,2,3,4,5,6,7}
	sl := ar[:3]
	fmt.Printf("Array: %v, Slice: %v \n", ar, sl)
	fmt.Printf("Appending to slice. \n")
	sl = append(sl, 100)
	fmt.Printf("Array: %v, Slice: %v \n", ar, sl)
}
