package main

import "fmt"

func main() {
	var sum = 0
	fmt.Printf("hello, world\n")
	for i := 0; i <= 9; i++ {
		sum += i
	}
	if sum == 45 {
		fmt.Println("you choosed nine")
	}
	fmt.Println(sum)
}
