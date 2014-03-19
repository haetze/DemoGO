package main

import (
	"fmt"
	"bufio"
	"net/http"
	"io"
)

type Hello struct {
	count int
}

func (h *Hello) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
		if r.Method == "GET"{
		h.count++
		fmt.Fprint(w, h.count)
	}else{
		reader := bufio.NewReader(r.Body)
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println(str)
			fmt.Fprint(w, h.count)
		}
	}

}
func main() {
	h := Hello{0}
	http.ListenAndServe("localhost:4000", &h)
}
