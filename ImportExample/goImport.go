// src in folder goLib can be imported and used
package main

import (
	"./goLib"
	"fmt"
)

func main() {
	fmt.Println("test")
	goLib.ForLoop()
}
