package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Test of the runtime:\nGo 1.2 runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X...")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}
