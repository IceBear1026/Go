package main

import (
	"fmt"
	"maps"
)

func main() {
	// Create a map with string KEYS and int VALUES
	// maps are key : value pairs
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m) // map: map[k1:7 k2:13]

	v1 := m["k1"]
	fmt.Println("v1:", v1) // v1: 7

	v3 := m["k3"]          // returns zero value for int which is 0
	fmt.Println("v3:", v3) // v3: 0
	// why? Because "k3" does not exist in the map, so it returns the zero value for the type of the map's values, which is int.
	// The zero value for int is 0, so v3 is assigned the value 0.

	// To check if a key exists in the map, we can use the "comma ok" idiom
	v2, ok := m["k2"]                      // v2 will hold the value of m["k2"] which is 13, and ok will be true because "k2" exists in the map
	fmt.Println("v2:", v2, "present?", ok) // v2: 13 present? true

	// if I wanted just check if a key exists without caring about the value, I can do this:
	_, ok = m["k2"]             // we use the blank identifier _ to ignore the value, and just check if "k2" exists in the map
	fmt.Println("present?", ok) // present? true
	// "_" means I know this returns a value but I don't need it.

	// btw we are setting ok to ok = instead of ok := because we already declared ok in the previous line, so we just want to update its value.

	fmt.Println("len:", len(m)) // len: 2

	delete(m, "k2")        // delete the key "k2" from the map
	fmt.Println("map:", m) // map: map[k1:7]

	clear(m)               // clear the map by deleting all keys
	fmt.Println("map:", m) // map: map[]

	_, prs := m["k2"]        // check if "k2" exists in the map after we cleared it.
	fmt.Println("prs:", prs) // prs: false

	// ==================================== //

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n) // map: map[bar:2 foo:1]

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
