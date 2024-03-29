package main

import (
	"errors"
	"fmt"
)

type List []int

func (l List) String() string {
	return "fool"
}

func (l List) isEqual(l2 List) (bool, error) {
	if len(l2) == 0 {
		return false, errors.New("empty list")
	} else {
		for i := 0; i < len(l2); i++ {
			if l[i] != l2[i] {
				return false, nil
			}
		}
		return true, nil
	}
}

func main() {
	l := List{1, 2, 3, 4, 5, 5, 6, 7}
	fmt.Println(l.isEqual(List{}))
	//l2 := List{};
	//fmt.Println(l == l2);
}
