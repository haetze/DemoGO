//2014, 12, 30
//haetze
//demo1.go

package main

import (
	"fmt"
)

type T struct {
	a int
	b int
}

func main() {
	a := "hahha"
	b := T{1, 1}
	changeArg(&a)
	changeArg2(&b)
	fmt.Println(a)
	fmt.Println(b)

}

func changeArg(a *string) {
	*a = "qwqwqwq"
}
func changeArg2(a *T) {
	a.a = 2
	a.b = 2
}
