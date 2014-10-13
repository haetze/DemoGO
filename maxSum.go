package main

import (
	"fmt"
)

func main() {
	sum := maxSum([]int{1, -2, 4, -6, 3 })
	fmt.Println(sum)

}

func maxSum(x []int) []int {
	if allSameSign(x) {
		return sum(x)
	} else if len(x) == 2 {
	}

	for !(testEq(x, shrink(x))) {
		x = shrink(x)
	}
	x = dropZero(x)
	n := findBig(x)
	big := x[n]
	for !(testEq(x, genNew(x))) {
		x = genNew(x)
	}
	x = shrink(x)
	if allSameSign(x) {
		w := sum(x)
		if w[0] < big {
			return []int{big}
		} else {
			return w
		}
	} else {
		b := findBig(x)
		if x[b] < big {
			return []int{big}
		} else {
			return []int{x[b]}
		}
	}

	return x
}

func testEq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func allSameSign(x []int) bool {
	b := x[0] < 0
	for n, i := range x {
		if n == 0 {
		} else {
			if b != (i < 0) {
				return false
			}
		}
	}
	return true
}

func sum(x []int) []int {
	var sum int
	for _, n := range x {
		sum = n + sum
	}
	return []int{sum}
}

func shrink(x []int) []int {
	for n, i := range x {
		if n < len(x)-1 {
			if (i < 0 && x[n+1] < 0) || (i > 0 && x[n+1] > 0) {
				x[n+1] = i + x[n+1]
				x[n] = 0
			}

		}
	}
	x = dropZero(x)
	return x
}

func dropZero(x []int) []int {
	var b []int
	for i := 0; i < len(x); i++ {
		if x[i] != 0 {
			b = append(b, x[i])
		}
	}
	return b
}

func findBig(x []int) int {
	n := 0
	for i := range x {
		if x[i] > x[n] {
			n = i
		}
	}
	return n
}

func genNew(x []int) []int {
	if x[0] < 0 {
		for i := range x {
			if (i%2) == 1 && i < len(x)-2 {
				if (x[i+1] + x[i] + x[i+2]) > (maxSum(x[:i+1]))[0] {
					x[i+1] = x[i] + x[i+1]
					x[i] = 0
				}
			}
		}
	}
	if x[0] > 0 {
		for i := range x {
			if (i%2) == 0 && i < len(x)-2 {
				if (x[i+1] + x[i] + x[i+2]) > (maxSum(x[:i+1]))[0] {
					x[i] = x[i] + x[i+1]
					x[i+1] = 0
					x = dropZero(x)
				}
			}

		}
	}
	x = dropZero(x)
	return x
}
