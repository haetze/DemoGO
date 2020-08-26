package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
	"math"
	"sort"
)

func split(arr []int) ([]int, []int) {
	n := len(arr)
	a := n / 2
	b := n / 2
	if n % 2 != 0 {
		a++
	}
	arr_a := make([]int, a)
	arr_b := make([]int, b)

	for i := 0; i < a; i++ {
		arr_a[i] = arr[i]
	}
	for i := a; i < n; i++ {
		arr_b[i-a] = arr[i]
	}
	
	return arr_a, arr_b

}

func merge(arr_a []int, arr_b []int) []int {
	a := len(arr_a)
	b := len(arr_b)
	n := a + b
	arr := make([]int, n)

	i_a := 0
	i_b := 0

	for i := 0; i < n; i++ {
		if i_a >= a {
			arr[i] = arr_b[i_b]
			i_b++
		} else if i_b >= b {
			arr[i] = arr_a[i_a]
			i_a++
		} else if arr_a[i_a] > arr_b[i_b] {
			arr[i] = arr_b[i_b]
			i_b++
		} else {
			arr[i] = arr_a[i_a]
			i_a++
		}		
	}
	return arr
	
}

func sort_par(arr []int, out chan []int) {

	if len(arr) <= 1 {
		out <- arr
	} else {
		c := make(chan []int, 2)
		arr_a, arr_b := split(arr)
		go sort_par(arr_a, c)
		go sort_par(arr_b, c)
		arr_s_1 := <- c
		arr_s_2 := <- c
		arr_s := merge(arr_s_1, arr_s_2)
		out <- arr_s
	}
}

func sort_selection(arr []int) []int {
	arr_s := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		j := 0
		for l := 0; l < len(arr); l++ {
			if arr[j] > arr[l] {
				j = l
			}
		}
		arr_s[i] = arr[j]
		arr[j] = math.MaxInt32
	}
	return arr_s
}

func sort_merge(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	arr_a, arr_b := split(arr)
	arr_a_s := sort_merge(arr_a)
	arr_b_s := sort_merge(arr_b)
	return merge(arr_a_s, arr_b_s)
}


func main() {
	m, err := strconv.Atoi(os.Args[1])
	fmt.Println(m)
	if err == nil {
		for i:=1; i <= m; i++ {
			n := int(math.Pow(2, float64(i)));
			arr := make([]int, n)
			for i := 0; i < n; i++ {
				arr[i] = n - i
			}
			c := make(chan []int, 1)
			start := time.Now()
			go sort_par(arr, c)
			end := time.Now()
			<- c
			fmt.Println("Merge Sort (parallel) took ", end.Sub(start).Nanoseconds(), "ns")

			start = time.Now()
			sort_merge(arr)
			end = time.Now()
			fmt.Println("Merge Sort (sequantially) took ", end.Sub(start).Nanoseconds(), "ns")
			
			start = time.Now()
			sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j]})
			end = time.Now()
			fmt.Println("Built Sort took ", end.Sub(start).Nanoseconds(), "ns")
			fmt.Println()
		}
	} else {
		fmt.Println("No proper index for array creation.")
	}
}
