package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30, 40, 50}
	for index, value := range numbers {
		fmt.Println(index, value)
	}

	// if I only care about the value and not the index, I can do this:
	for _, value := range numbers {
		fmt.Println(value)
	}
}
