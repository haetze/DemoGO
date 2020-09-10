package main


import "fmt"

const size = 3
const next = 0
const end = 1
const print = 3

func thrower(c chan int) {
	for {
		<-c
	}
}

func cell(n, ne, e, es, s, sw, w, wn, me chan int, i, j int, call chan int, state int, end_c chan bool, num int) {

	for command := range call {
		if command == next {
			n <- state
			ne <- state
			e <- state
			es <- state
			s <- state
			sw <- state
			w <- state
			wn <- state
			sum := 0
			for n := 0; n < num; n++ {
				r := <-me
				sum = sum + r
			}
			if sum >= 2 && sum <= 4 {
				state = 1
			} else {
				state = 0
			}
		} else if command == print {
			fmt.Printf("(%d,%d) : %d\n", i, j, state)
		} else {
			end_c <- true
		}		
	}
}

func starter(cs [size][size]chan int, throw_away chan int, i int, j int, caller chan int, end_c chan bool) {
	num := 0
	n := throw_away
	e := throw_away
	s := throw_away
	w := throw_away
	ne := throw_away
	es := throw_away
	sw := throw_away
	wn := throw_away
	if i > 0 {
		num = num + 1
		n = cs[i-1][j]
	}
	if i > 0 && j > 0 {
		num = num + 1
		wn = cs[i-1][j-1]
	}
	if i > 0 && j < size - 1 {
		num = num + 1
		ne = cs[i-1][j+1]
	}
	if j > 0 {
		num = num + 1
		w = cs[i][j-1]
	}
	if j < size - 1 {
		num = num + 1
		e = cs[i][j+1]
	}
	if i < size - 1 {
		num = num + 1
		s = cs[i + 1][j]
	}
	if i < size - 1 && j > 0 {
		num = num + 1
		sw = cs[i+1][j-1]
	}
	if i < size - 1 && j < size - 1 {
		num = num + 1
		es = cs[i+1][j+1]
	}

	init_states := [size][size]int{
	{0,1,1},
	{0,0,1},
	{0,0,1}}

	go cell(n, ne, e, es, s, sw, w, wn, cs[i][j], i, j, caller, init_states[i][j], end_c, num)
	
}

func call_next(c [size][size]chan int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++{
			c[i][j] <- next
		}
	}
}

func call_print(c [size][size]chan int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++{
			c[i][j] <- print
		}
	}
}

func call_end(c [size][size]chan int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++{
			c[i][j] <- end
		}
	}
}

func wait_for_end(c chan bool) {
	for i := 0; i < size * size; i++ {
		<-c
	}
}

func main() {
	throw_away := make(chan int)
	
	end_c := make(chan bool, size*size)
	go thrower(throw_away)

	var cs [size][size]chan int
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++{
			cs[i][j] = make(chan int, 8)
		}
	}
	var caller [size][size]chan int
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++{
			caller[i][j] = make(chan int)
		}
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++{
			starter(cs, throw_away, i, j, caller[i][j], end_c)
		}
	}

	call_print(caller)
	call_next(caller)
	fmt.Println("==================")
	call_print(caller)
	call_end(caller)
	wait_for_end(end_c)

}
