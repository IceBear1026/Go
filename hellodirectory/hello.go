// every .go file must start with a package declaration.
// it tells Go: this file belongs to this package.
// This file belongs to a package named hello
// For example, imagine this folder structure:

package hellopackage

import (
	"fmt"
	"strings"
)

// func FunctionName(parameterName parameterType) returnType
func Say(name string) string {
	// Sprintf does not PRINT, it returns a STRING.
	return fmt.Sprintf("Hello, %s!", name)

	// this returns PRINT. When a function returns nothing, you simply omit the return type.
	// func Say(name string) {
	// fmt.Printf("Hello, %s!", name)
}

func Say2(names []string) string {
	// []string seems to be a slice of strings, which is a more flexible version of an array.
	// and this slice is equivalent to vector in C++.
	// it's a dynamic list. You can append to it, and it can grow as needed.
	// names := []string{"Steven", "Kim", "Alex"}
	// you can access items by index
	// fmt.Println(names[0]) // Steven
	// fmt.Println(names[1]) // Kim
	// fmt.Println(len(names)) // 3
	// names = append(names, "Bob")
	// [3]string is an array with fixed size of 3. You cannot append to it, and it cannot grow as needed.
	// []string is a dynamic array.

	// we are returning a string with strings package's Join function.
	// func Join(s []string, sep string) string
	// another example:
	/*
			// array of strings.
		    str := []string{"Geeks", "For", "Geeks"}

		    // joining the string by separator
		    fmt.Println(strings.Join(str, "-"))

			Output:
			Geeks-For-Geeks
	*/
	return "Hello, " + strings.Join(names, ", ") + "!"
}

/*
var message string = "Hello"
The above is the original way to declare a variable in Go.
message := "Hello"
This means create a new variable named message, automatically figure out its type, and assign "Hello" to it.
similar to in C++
auto message = "Hello";
*/
