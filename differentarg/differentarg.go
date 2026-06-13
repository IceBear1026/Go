package main

import (
	"fmt"
	"os"
)

func main() {
	// so if we receive any arguments, we will print.
	if len(os.Args) > 1 {
		fmt.Println("Hello,", os.Args[1])
		// else
	} else {
		fmt.Println("Hello, world!")
	}
}
