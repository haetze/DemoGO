package main

import (
	"fmt"
	"time"
)

func say(s string, c chan int) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
	c <- 0
}

func main() {
	c1 := make(chan int, 100)
	c2 := make(chan int, 100)
	go say("world", c1)
	fmt.Println("test")
	for {
		select {
		case x := <-c1:
			if x == 0 {
				fmt.Printf("finish\n")
			}
			go say("hello", c2)
		case y := <-c2:
			if y == 0 {
				fmt.Printf("finish2\n")
				return
			}
		}

	}
}
