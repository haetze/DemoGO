//2014, 10, 24
//haetze
//pipeExample.go

package main

import(
	"fmt"
	"io"
)

func main(){
	reader, writer := io.Pipe()
	go startRead(reader)
	writer.Write([]byte("test"))
	writer.Close()
}

func startRead(reader *io.PipeReader){
	v := make([]byte, 100) 
	_, _ = reader.Read(v)
	fmt.Println(string(v))
}

