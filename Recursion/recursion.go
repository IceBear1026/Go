package main

import (
	"fmt"
)

// this fact function calls itself until it reaches the base case of fact(0) which returns 1.
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	// this is a recursive function that calculates the factorial of n.
	fmt.Println(fact(7))

	// so for this we are instantiating a variable function called fib.
	// anonymous functions can be recursive, but this requires explicitly declaring a variable with var to store the function before it's defined.

	// I still don't understand anonymous functions, but I think the idea is that we can define a function without giving it a name, and we can assign it to a variable. This allows us to create closures and also to have recursive functions without needing to name them.
	var fib func(n int) int

	// and here we are defining what this function will do.
	// this is a recursive function that calculates the nth Fibonacci number.

	// the Fibonacci sequence is defined as:
	// F(0) = 0
	// F(1) = 1
	// F(n) = F(n-1) + F(n-2) for n > 1
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))

	// so is anonymous function basically allowing you to create a function within the main function without having to create a whole new separate function outside? Is that it's main use case? I guess it also allows you to create closures, which is another important concept in Go.
	// closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.
}
