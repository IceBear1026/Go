package main

import (
	"fmt"
	hellopackage "hellomodule/hellodirectory"
	"os"
)

func main() {
	// so if we receive any arguments, we will print.
	// len() returns integer
	if len(os.Args) > 1 {
		// This receives the arguments as a slice of strings, and we are passing it to Say2 function.
		// os.Args[0] is the name of the program, so we are slicing it from index 1 to the end to get the actual arguments.
		fmt.Println(hellopackage.Say2(os.Args[1:]))
		// else
	} else {
		fmt.Println(hellopackage.Say("world"))
	}
}
