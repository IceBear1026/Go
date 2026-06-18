package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func main() {
	a, b := vals()
	fmt.Println(a) // 3
	fmt.Println(b) // 7

	_, c := vals()
	fmt.Println(c) // 7

	c, _ = vals()
	fmt.Println(c) // 3

	/*
		the blank identifier _ does not only mean "check if a value exists"
		It has more general meaning

		It means "I am intentionally ignoring this value"

		I saw it with maps because map lookup returns two things:
		value, ok := m["k2"]

		so when you write:
		_, ok := m["k2"]
		you are saying "Ignore the actual value. Only keept the ok boolean"

		But _ itself is not doing the existence check. The map lookup is what gives you the existence boolean.

		Inside the vals() function, it will always return two integers.
		first value: 3
		second value: 7

		So this a, b:= vals() means:
		a = 3
		b = 7

		And this:
		_, c := vals() means:
		ignore the first value (3) and assign the second value (7) to c.

		The important thing is position.
		_ is in the first position, so it receives and ignores the first return value.

		c is in the second position, so it receives the second return value.

		c, _ := vals()
		would do the opposite: it would assign 3 to c and ignore 7.
	*/

}
