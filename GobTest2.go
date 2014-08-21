package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

type Q struct {
	N int
	B int
}

func main() {
	var s Q
	pipe, _ := os.OpenFile("gobEx", os.O_RDWR, 0666)
	enc := gob.NewEncoder(pipe)
	err := enc.Encode(Q{21,33})
	pipe.Seek(0, 0)
	if err != nil {
		fmt.Println(err)
	}
	dec := gob.NewDecoder(pipe)
	err2 := dec.Decode(&s)

	if err2 != nil {
		fmt.Println(err2)
	}

	fmt.Println(s)

}
