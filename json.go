package main

//json test

import (
	"encoding/json"
	"fmt"
)

type Mes struct {
	Name string
	Age  int
}

func strin(str string) string {
	return str[1:]
}

func main() {
	m := Mes{"test", 12}
	a, _ := json.Marshal(m)
	fmt.Println(string(a))
	var n Mes
	json.Unmarshal(a, &n)
	t := n.Name
	fmt.Println(t)
}
