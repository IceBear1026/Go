package main

import "fmt"

// this ...int means that the function sum can take a variable number of int arguments.
// the function will receive these arguments as a slice of ints.
// variadic functions can be called with any number of trailing arguments.
// even fmt.Println is a variadic function that can take any number of arguments and print them out.
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	// we are using "_" because we don't care about the index of the range.
	// also for range, it seems like we don't need to set a length rather just name "nums"
	// and it will automatically iterate over all the elements in the slice.
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...) // we need to use "..." to unpack the slice into individual arguments for the variadic function.
	// if you already have multiple args, apply them to a variadic function by using func(slice...) like the above.
}
