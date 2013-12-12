package main

import(
	"fmt"
	"net/http"
);

type Server struct {
	address string;
	channel chan string;
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Hello, world.");
	s.channel<- "got req and handeld";
}


func server(c chan string){
	Server := Server{"127.0.0.1:1234", c};
	c <- "got in server from function ";
	http.HandleFunc("/", startPage)
	http.HandleFunc("/12", page12);
	http.Handle("/server", &Server);
	http.ListenAndServe("127.0.0.1:1234", nil);
}

func page12(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "welcome to page 12");
}

func startPage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Welcome to start page");
}

func mon(c chan string){
	fmt.Println("start monitoring");
	for{
		select{
		case x:= <-c:
			fmt.Println(x);
			fmt.Println("still mon");
		}
	}
}

func main(){
	fmt.Println("test");
	c := make(chan string, 100);
	go server(c);
	go mon(c);
	for{
		select{
		case x:=<-c:
			fmt.Println(x);
		}
	}
}
