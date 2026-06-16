package main

import (
	"fmt"
)

func main() {
	i := 1
	for i <= 3 {
		fmt.Print(i)
		i++
	}
	// this outputs 1, 2, 3.

	fmt.Println() // adding a new line for better formatting.

	for j := 0; j < 3; j++ {
		fmt.Printf("%d", j) // this is another way to write a print statement. the %d is a placeholder for an integer. we can also use %s for strings, %f for floating point numbers, etc.
	}
	// this outputs 0, 1, 2.

	fmt.Println()

	// integer k here is being assigned the value of the range of 3, which is 0, 1, 2.
	for k := range 3 {
		fmt.Print(k)
	}

	fmt.Println()

	/*
		but how would I make k start at a certain value instead of 0?
		start := 5

		for k := range 3 {
			fmt.Println(start + k)
		}
	*/

	// this will print "loop" once, because the for loop will run indefinitely until we break out of it. the break statement will exit the loop after the first iteration.
	for {
		fmt.Println("loop")
		break
	}

	for n := range 6 {
		// this will check even numbers
		if n%2 == 0 {
			continue // this will skip the rest of the loop and go to the next iteration if n is even.
		}
		fmt.Print(n, " ")
	}

	// Sprintf does not PRINT, it returns a STRING.

}
