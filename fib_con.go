package main

import (
	"fmt"
	"time"
)


func fib_seq(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib_seq(n-1) + fib_seq(n-2)
	}
}

func fib_par(n int, c chan int) {
	if n == 0 {
		c <- 0
	} else if n == 1 {
		c <- 1
	} else {
		d := make(chan int)
		go fib_par(n-1, d)
		go fib_par(n-2, d)
		c <- <-d + <-d
	}
}


func main() {
	n := 50
	
	start := time.Now()
	fmt.Println("Fib of ", n," is ", fib_seq(n))
	end := time.Now()
	fmt.Println("fib_seq took ", end.Sub(start).Nanoseconds(), "ns")

	
	c := make(chan int)
	start = time.Now()
	go fib_par(n, c)
	fmt.Println("Fib of ", n, " is ",<-c)
	end = time.Now()
	fmt.Println("fib_seq took ", end.Sub(start).Nanoseconds(), "ns")

}
