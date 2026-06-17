package main

import (
	"fmt"
)

func main() {
	var s []string // this is a slice of strings. A slice is a dynamically-sized, flexible view into the elements of an array.
	fmt.Println("uninit: ", s, s == nil, len(s) == 0)
	// this prints: uninit: [] true true because the default value of a slice is nil, and the length of a nil slice is 0.

	s = make([]string, 3) // this creates a slice of strings with length 3. The make function is used to create slices, maps, and channels.
	// but didn't we just create a slice with var s []string? Yes, but that slice is nil and has length 0.
	// the make function allocates and initializes a slice with the specified length and capacity.
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))
	// emp: [  ] len: 3 cap: 3 because we created a slice of length 3, and the capacity is also 3 because we didn't specify a different capacity.

	// we can set and get just like with arrays.
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s) // this prints: set: [a b c] because we set the elements of the slice to "a", "b", and "c".
	// again even though this is a slice, we can still print it without looping through it.
	fmt.Println("get:", s[2]) // this prints: get: c because we get the value of the last element.

}
