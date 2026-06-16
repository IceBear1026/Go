package main

import (
	"fmt"
)

func main() {
	var a = "initial"
	fmt.Println(a)      // initial. this is how we declare a variable in golang. we use the var keyword, then the name of the variable, then the value of the variable. we can also specify the type of the variable, but golang can infer the type from the value.
	var b, c int = 1, 2 // this is another way to declare variables. we can declare multiple variables at once, and we can specify their types as well.
	fmt.Println(b, c)
	var d = true   // this is a boolean variable.
	fmt.Println(d) // print s true
	var e int
	fmt.Println(e) // 0. when we declare a variable without giving it a value, it gets the zero value for its type. for integers, the zero value is 0. for booleans, the zero value is false. for strings, the zero value is "" (empty string).

	f := "apple"
	fmt.Println(f) // apple. this is a shorthand way to declare and initialize a variable.

}
