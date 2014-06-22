package main

import (
	"fmt"
	"test"
)

func sub(a, b int ) int {
	return test.Add(a, b*-1);
}
func main() {
	fmt.Println(sub(12, -1))
}
