package main

import "fmt"

type pair[A any, B any] struct {
	fst A
	snd B
}

func swap[A any, B any](p pair[A,B]) pair[B,A] {
	q := pair[B,A]{fst: p.snd, snd: p.fst}
	return q
}


func maper[A any, B any](f func(a A) B, arr []A) []B {
	arr2 := make([]B,len(arr))
	for i,e := range arr {
		arr2[i] = f(e)
	}
	return arr2
}

func main() {
	p := pair[int, byte]{fst: 1, snd: 'h'}
	
	fmt.Println(p)

	q :=  swap(p)

	fmt.Println(q)

	arr1 := make([]float32, 100)
	for i,_ := range arr1 {
		if i == 0 {
		} else if i == 1 {
			arr1[i] = 1.0
		} else {
			arr1[i] = arr1[i-2] + arr1[i-1]
		}
	}

	fmt.Println(arr1)

	arr2 := maper(func(x float32) float64 {return float64(x*2)}, arr1)

	fmt.Println(arr2)
}
