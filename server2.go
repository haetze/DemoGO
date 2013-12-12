package main

import (
    "net/http"
    "fmt"
)
type Struct struct {
    Greeting string
    Punct    string
    Who      string
}

type String string

func (str *Struct) ServeHTTP(w http.ResponseWriter, r *http.Request){
    fmt.Fprint(w, str.Greeting);
    fmt.Printf("got aa mess");
}

func (str String)ServeHTTP(w http.ResponseWriter, r *http.Request){
    fmt.Fprint(w, str);
}
func main() {
    // your http.Handle calls here
    var  struc Struct = Struct{"hi",":","ben"};
    var str String = String("123");
    http.Handle("/string", str);
    http.Handle("/struct", &struc);
    http.ListenAndServe("localhost:4000", nil);
}
