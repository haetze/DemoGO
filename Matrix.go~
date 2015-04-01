package main

import (
	"fmt"

	"github.com/haetze/Matrix"
)

func main() {
	a := Matrix.M{{0.25, 0, 0}, {0.7, 0.55, 0}, {0, 0.4, 0.95}}
	b := Matrix.M{{0.8, 0, 0}, {0, 0.7, 0}, {0, 0, 0.55}}
	c := Matrix.M{{1.25, 0.4285714286, 0.81818181818181818182}, {0, 1, 0}, {0, 0, 1}}
	d := Matrix.Mult(a, b)
	d2 := Matrix.CreateM(d)
	x := Matrix.Mult(d2, c)
	fmt.Println(x)
}
