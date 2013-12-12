package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func newFib() func() int {
	a, b := 0, 1
	return (func() int {
		a, b = b, a+b
		return a
	})
}

func main() {
	fib := newFib()
	t1 := time.Now().UnixNano()
	a := "40"
	if len(os.Args) > 1 {
		a = os.Args[1]
	}

	n, err := strconv.ParseInt(a, 10, 64)
	if err != nil {
		return
	}
	//b := 0;
	var i int64 = 1
	for ; i < n; i++ {
		_ = fib()
	}
	//fmt.Println(b);

	fmt.Println(fib())
	t2 := time.Now().UnixNano()
	fmt.Println((t2 - t1) / 1000000)
}
