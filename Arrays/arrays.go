package main

import (
	"fmt"
)

func main() {
	// arrays are fixed length sequences of a particular type.
	var a [5]int           // this is an array of 5 integers. The length is part of the type, so [5]int and [6]int are DIFFERENT types.
	fmt.Println("emp:", a) // this will print the array.
	// emp: [0 0 0 0 0] because the default value for int is 0.
	// it seems like we don't have to specify that it's an array.
	// it also seems like when we print, we don't have to loop through the array to print each element.

	a[4] = 100 // we can set the value of an element in the array using the index. The index starts at 0, so a[4] is the 5th element.
	fmt.Println("set:", a)
	// this prints: set: [0 0 0 0 100] because we set the last element to 100.
	fmt.Println("get:", a[4]) // this prints: get: 100 because we get the value of the last element.

	fmt.Println("len:", len(a)) // this prints: len: 5

	// ===================================================== //

	b := [5]int{1, 2, 3, 4, 5} // this is a shorthand way to declare and initialize an array. We can also omit the length and let it be inferred from the number of elements.
	fmt.Println("dcl:", b)     // this prints: dcl: [1 2 3 4 5]

	b = [...]int{1, 2, 3, 4, 5} // this is another way to declare and initialize an array. The ... tells the compiler to infer the length from the number of elements.
	fmt.Println("dcl:", b)      // this also prints: dcl: [1 2 3 4 5]

	b = [...]int{100, 3: 400, 500}
	fmt.Println("dcl:", b) // this prints: dcl: [100 0 0 400 500] because we set the first element to 100, the 4th element (index 3) to 400, and the 5th element (index 4) to 500. The other elements are set to the default value of 0.

	// ===================================================== //

	var twoD [2][3]int // this is a 2D array, which is an array of arrays. The first dimension has length 2, and the second dimension has length 3.
	// so visually, it looks like this:
	// [[0 0 0]
	//  [0 0 0]]
	// [x][y] where x is amount of arrays and y is amount of elements in each array.

	// this for loop goes 0, 1.
	for i := range 2 {
		// this for loop goes 0, 1, 2.
		for j := range 3 {
			twoD[i][j] = i + j // this sets the value of each element in the 2D array to the sum of its indices.
		}
	}
	fmt.Println("2d:", twoD) // this prints: 2d: [[0 1 2] [1 2 3]]

	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d:", twoD)
	// Note that arrays appear in the form [v1 v2 v3 ...] when printed with fmt.Println.
}
