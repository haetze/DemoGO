package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
)

type Q struct {
	N int
	B int
}

func main() {
	var pipe bytes.Buffer
	/*	enc := gob.NewEncoder(&pipe)
		err := enc.Encode(Q{12, 12})
		e := ioutil.WriteFile("gobEx", pipe.Bytes(), 0644)
		fmt.Println(e)*/
	var s Q
	b, _ := ioutil.ReadFile("gobEx")
	pipe = *bytes.NewBuffer(b)
	//fmt.Println(pipe)
	dec := gob.NewDecoder(&pipe)
	err2 := dec.Decode(&s)

	if err2 != nil {
		fmt.Println(err2)
	}

	fmt.Println(s)

}
