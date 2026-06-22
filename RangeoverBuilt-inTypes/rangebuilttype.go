package main

import (
	"fmt"
)

func main() {
	// a slice because no value within the [].
	nums := []int{2, 3, 4}
	sum := 0
	// this a for loop that only is intaking the value without its index.
	// and this range nums just takes the length of the numbs slice starting form 0 - nums.length. Convenient way in Go.
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum) // sum: 9

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i) // index>: 1
		}
	}

	// this creating a map that uses string key and string value.
	// and we hardcoded the map with 2 key value pairs.
	kvs := map[string]string{"a": "apple", "b": "banana"}
	// going through the key and value inside the map using range kvs.
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v) // a -> apple, b -> banana
	}

	// When using range on a map in Go, you can either receive both the key and the value, or receive only the key. For example, for k, v := range kvs gives you both values, where k is the key and v is the value. If you only write one variable, like for k := range kvs, Go gives you only the key by default because the key is the first value returned by range on a map. However, if you want only the value, you must use the blank identifier _ to ignore the key, like for _, v := range kvs. This is because the value is the second value returned by range, so you need _ in the first position to skip the key.
	for k := range kvs {
		fmt.Println("key: ", k) // key: a, key: b
	}

	// range on strings iterates over Unicode code points. The first value is the starting byte index of the rune and the second the rune itself.
	for i, c := range "go" {
		fmt.Println(i, c) // 0 103, 1 111
	}

	/*
		i = byte index where the character starts
		c = the character itself, as a rune

		A rune is Go's name for a Unicode code point. Practically, you can think of it as "a character value."

		for "go":
		g has Unicode code point 103
		o has Unicode code point 111

		Why numbers? Because c is a rune, and fmt.Println prints runes as their numeric Unicode values by default.

		To print the actual characters, use %c:

		for i, c := range "go" {
			fmt.Printf("%d: %c\n", i, c) // 0: g, 1: o
		}

		Or to show both the number and character:
		for i, c := range "go" {
			fmt.Printf("index: %d, rune: %c, number: %d\n", i, c, c)
		}

		Output:
		index: 0, rune: g, number: 103
		index: 1, rune: o, number: 111
	*/

}
