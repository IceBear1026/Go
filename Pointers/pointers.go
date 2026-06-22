package main

import (
	"fmt"
)

// a function that takes in an int but no return value.
// it sets whatever that value that inputted to 0.
func zeroval(ival int) {
	ival = 0
}

// similar to C++, we are using * to indicate that this is a pointer to an int, and we are using & to get the memory address of the variable.
// so we are setting the value at the memory address of iptr to 0.
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i) // initial: 1

	// why doesn't this work? Because in Go, when you pass a variable to a function, it is passed by value, meaning that the function receives a copy of the variable. So when zeroval sets ival to 0, it only changes the copy of i that was passed to the function, not the original i in main.
	// therefore zeroval(i) i remains 1 after the function call.
	zeroval(i)
	fmt.Println("zeroval:", i) // zeroval: 1

	// this works because we are passing the memory address of i to zeroptr, so when zeroptr sets the value at that memory address to 0, it changes the original i in main.
	zeroptr(&i)
	fmt.Println("zeroptr:", i) // zeroptr: 0

	fmt.Println("pointer:", &i) // pointer: 0xc0000160b8
}

/*
We'll show how pointers work in contrast to values with 2 functions: zeroval and zeroptr. zeroval has an int parameter, so arguments will be passed to it by value. zeroval will get a copy of ival distinct from the one in the calling function.

zeroptr in contrast has a *int parameter, meaning that it takes an int pointer. The *iptr code in the function body then dereferences the pointer from its memory address to the current value at that address. Assigning a value to a dereferenced pointer changes the value at the referenced address.

The &i syntax gives the memory address of i, i.e. a pointer to i.

Pointers can be printed too.const

zeroval doesn't change the i in main, but zeroptr does because it has a reference to the memory address for that variable.
*/
