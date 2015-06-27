//2015, 06, 27
//haetze
//structInherenceTest.go

package main

import (
	"fmt"
)

type First struct {
	b int
	m bool
	Second
}

type Second struct {
	m bool
}

type Test interface {
	test() bool
}

func (s Second) test() bool {
	return s.m
}

func main() {
	s := Second{true}
	f := First{12, false, s}

	fmt.Print(m(f))
}

func m(s Test) bool {
	return s.test()
}
