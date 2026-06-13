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
		// It adds spaces between arguments only
		fmt.Println(hellopackage.Say(os.Args[1]))
		// else
	} else {
		fmt.Println(hellopackage.Say("world"))
	}
}
