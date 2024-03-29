package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

type String string

func (str *Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, str.Greeting)
	fmt.Printf("got aa mess")
	fmt.Println(r.ContentLength)
}

func (str String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, str)
	fmt.Println(r.ContentLength)
}
func main() {
	// your http.Handle calls here
	fileString, err := ioutil.ReadFile("page.html")
	if err != nil {
		return
	}
	var struc Struct = Struct{"hi", ":", "ben"}
	var str String = String(fileString)
	http.Handle("/string", str)
	http.Handle("/struct", &struc)
	http.ListenAndServe("localhost:4000", nil)
}
