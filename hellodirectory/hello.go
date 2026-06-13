// every .go file must start with a package declaration.
// it tells Go: this file belongs to this package.
// This file belongs to a package named hello
// For example, imagine this folder structure:

package hellopackage

import "fmt"

// func FunctionName(parameterName parameterType) returnType
func Say(name string) string {
	// Sprintf does not PRINT, it returns a STRING.
	return fmt.Sprintf("Hello, %s!", name)

	// this returns PRINT. When a function returns nothing, you simply omit the return type.
	// func Say(name string) {
	// fmt.Printf("Hello, %s!", name)
}

/*
var message string = "Hello"
The above is the original way to declare a variable in Go.
message := "Hello"
This means create a new variable named message, automatically figure out its type, and assign "Hello" to it.
similar to in C++
auto message = "Hello";
*/
