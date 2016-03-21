package main

import "fmt"

type Number interface {
	add(Number) Number
	sub(Number) Number
	mult(Number) Number
	div(Number) Number
	NUMBER () float64
}

type NUMBER struct {
	n float64
}

func main() {
	n := NUMBER{12}
	test(n.add(NUMBER{12.12}))
	
}

func test(x Number) {
	fmt.Println(x)
}

func (x NUMBER) add(y Number) Number{
	return NUMBER{x.n + y.NUMBER()}
}

func (x NUMBER) sub(y Number) Number{
	return NUMBER{x.n - y.NUMBER()}
}


func (x NUMBER) mult(y Number) Number{
	return NUMBER{x.n * y.NUMBER()}
}

func (x NUMBER) div( y Number) Number{
	return NUMBER{x.n / y.NUMBER()}
}

func (x NUMBER) NUMBER () float64{
	return x.n
}
