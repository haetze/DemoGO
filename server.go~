
package main

import ("fmt"
	"net/http"
	"io/ioutil")

type Hello struct{
	count int ;
}

func (h *Hello) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, "Hello!");
	h.count++;
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(body));
}
func main() {
	h := Hello{0};
	http.ListenAndServe("localhost:4000", &h)
}
