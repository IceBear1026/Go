package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "สวัสดี"
	fmt.Println("Len:", len(s))

	// Count the runes in the string.
	// A rune is a Unicode code point.
	for i := 0; i < len(s); i++ {
		// %x prints the hexadecimal value of the byte.
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	// to count the runes, we can use utf8.RuneCountInString.
	// this counts the number of runes in the string, which is not necessarily the same as the number of bytes.
	// why different? because some runes may be represented by multiple bytes in UTF-8 encoding.
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	/*

	 */
}
