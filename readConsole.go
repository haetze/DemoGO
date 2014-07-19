package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	read := bufio.NewReader(os.Stdin)
	str, er := read.ReadString('\n')
	if er == nil {
		n := len(str)
		i, err := strconv.ParseInt(str[:n-1], 0, 64)
		if err != nil {
			fmt.Println("I wanted a number !!")
		} else {
			fmt.Println(float64(i) / 2)
		}
	}
}
