package main

import(
	"fmt"
	"net/http"
	"io/ioutil"
)

type Responder string;

func (res Responder) ServeHTTP ( w http.ResponseWriter, r *http.Request){
	if r.Method == "GET"{
		fmt.Fprint(w, res);
		fmt.Println("in the fun");
	}else{
		fmt.Fprint(w, "lappen");
		fmt.Println(r.ContentLength);
	}
}

func main(){
	fileString, err :=  ioutil.ReadFile("page.html");
	if err != nil{
		return;
	}
	responder :=Responder(fileString);
	http.Handle("/page", responder);
	http.ListenAndServe("localhost:4000", nil);
}

