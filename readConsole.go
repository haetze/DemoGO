package main

import( 
	"fmt"
	"os"
	"bufio"
)


func main (){
	read := bufio.NewReader(os.Stdin)
	str, _ := read.ReadString('\n')
	fmt.Println(str)
}

