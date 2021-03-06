package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func newFib() func() *big.Int {
	a := big.NewInt(0)
	b := big.NewInt(1)
	return (func() *big.Int {
		a, b = b, a.Add(a, b)
		return a
	})
}

func main() {
	fib := newFib()
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

	y := fib()
	fmt.Println(y)
}
