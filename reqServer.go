package main

import(
	"fmt"
	"net/http"
	"io"
	"io/ioutil"
	"bufio"
)

type Responder string;

func (res Responder) ServeHTTP ( w http.ResponseWriter, r *http.Request){
	if r.Method == "GET"{
		fmt.Fprint(w, res);
		fmt.Println("in the fun");
	}else{	
		reader:= bufio.NewReader(r.Body);
		str, err := reader.ReadString('\n');
		if err == io.EOF{
			fmt.Println(str);
			fmt.Fprint(w, "lappen\n");
		}
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

