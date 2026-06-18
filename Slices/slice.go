package main

import (
	"fmt"
	"slices"
)

func main() {
	var x []string // this is a slice of strings. A slice is a dynamically-sized, flexible view into the elements of an array.
	fmt.Println("uninit: ", x, x == nil, len(x) == 0)
	// this prints: uninit: [] true true because the default value of a slice is nil, and the length of a nil slice is 0.

	s := make([]string, 3) // this creates a slice of strings with length 3. The make function is used to create slices, maps, and channels.
	// but didn't we just create a slice with var x []string? Yes, but that slice is nil and has length 0.
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

	fmt.Println("len:", len(s))

	s = append(s, "d")      // this appends the value "d" to the slice s. The append function is used to add elements to a slice.
	s = append(s, "e", "f") // we can also append multiple values at once.
	fmt.Println("apd:", s)  // this prints: apd: [a b c d e f] because we appended "d", "e", and "f" to the slice.

	// ====================================== //
	c := make([]string, len(s)) // this creates a new slice of strings with the same length as s.
	copy(c, s)                  // this copies the elements of s into c. The copy function is used to copy elements from one slice to another.
	// it returns the number of elements copied, which is the minimum of the lengths of the two slices.
	// in this case, it will copy all the elements of s into c because they have the same length.
	fmt.Println("cpy:", c) // this prints cpy: [a b c d e f] because we copied the elements of s into c.

	l := s[2:5] // this creates a slice of s from index 2 to index 5 (exclusive). This is called slicing.
	// the syntax is slice[low:high] where low is the starting index and high is the ending index (exclusive).
	fmt.Println("sl1:", l) // this prints: sl1: [c d e] because we sliced s from index 2 to index 5, which gives us the elements "c", "d", and "e".

	l = s[:5]              // this creates a slice of s from the beginning to index 5 (exclusive). This is a shorthand for s[0:5].
	fmt.Println("sl2:", l) // this prints: sl2: [a b c d e] because we sliced s from the beginning to index 5, which gives us the elements "a", "b", "c", "d", and "e".
	l = s[2:]              // this creates a slice of s from index 2 to the end. This is a shorthand for s[2:len(s)].
	fmt.Println("sl3:", l) // this prints: sl3: [c d e f] because we sliced s from index 2 to the end, which gives us the elements "c", "d", "e", and "f".

	t := []string{"g", "h", "i"} // we just created and initialized a slice with these values.
	fmt.Println("dcl:", t)       // this prints: dcl: [g h i] because we created a slice with the values "g", "h", and "i".

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2") // this prints: t == t2 because the slices are equal.
	}
	// usually in C++, I would have to loop through the vectors and compare each element.
	// but in Go, slices have a built-in function called slices.Equal to compare two slices for equality.

	twoD := make([][]int, 3) // this creates a 2D slice, which is a slice of slices. The outer slice has length 3.
	// this is looping for 0, 1, 2 indexes as variable i.
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen) // this creates an inner slice of integers with length innerLen and assigns it to the i-th element of the outer slice.
		for j := range innerLen {
			twoD[i][j] = i + j // this sets the value of each element in the inner slice to the sum of its indices.
		}
	}
	fmt.Println("2d:", twoD) // this prints: 2d: [[0] [0 1] [0 1 2]] because we created a 2D slice with 3 inner slices of lengths 1, 2, and 3, and we set the values of the elements to the sum of their indices.
}
