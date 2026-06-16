// makes this the main package
package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
}

// we write "go run" to run this program.
// sometimes we want to build our programs into binaries.
// we can do this using "go build"
/*
If you run go build in a folder with package main, you get an executable file (same name as folder/module, platform-dependent extension like .exe on Windows).
If the package is not main (library package), it compiles for checking/caching but does not produce a runnable executable.
*/
