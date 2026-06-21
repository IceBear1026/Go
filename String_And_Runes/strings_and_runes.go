package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// why isn't this const := s?
	// because const is a keyword, not a variable declaration operator.
	// const is used to declare a constant value, while := is used to declare and initialize a variable.
	const s = "สวัสดี"

	fmt.Println("Len:", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}

	// equivalent for loop
	// for i, c := range s
	fmt.Println()

	// the reason why this rune count is different from the len(s) is because
	// len(s) returns the number of bytes in the string, while utf8.RuneCountInString(s)
	// returns the number of runes (Unicode code points) in the string.
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	/*
		A Go string is a read-only slice of bytes. The language and the standard library treat strings spcially - as container of text encoded in UTF-8. In other languages, strings are made of "characters". In Go, the concept of character is called a rune - it's an integerthat represents a Unicode code point.

		s is a string assigned a literal value representing the word "hello" in Thai language. Go string literals are UTF-8 encoded text.

		Some strings are equivalent to []byte, this will produce the length of the raw bytes stored within.

		Indexing into a string produces the raw byte value at each index. This loop generates the hex values of all the bytes that constiute the code points in s.

		To count how many runes are in a string, we can use the utf8 package. Note that the run-time of RuneCountInString depends onthe size of the string, because it has to decode each UTF-8 rune sequentially. Some Thai characters are represented by UTF-8 code points that can span multiple bytes, so the reuslt of this count may be surprising.
	*/

	for idx, runeValue := range s {
		// the %#U verb prints the rune in a human-readable format, along with its Unicode code point value.
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println("\nUsing DecodeRuneInString")
	// i and w are both instantiated as 0. i is the index of the current byte in the string, and w is the width of the current rune in bytes.
	// hence we are doing i+=w to move to the next rune in the string.
	for i, w := 0, 0; i < len(s); i += w {
		// we are instantiating these values using the utf8.DecodeRuneInString function, which decodes the first rune in the string starting at index i and returns the rune value and its width in bytes.

		// I don't understand this s[i:]. Doesn't this mean the full string from that index is going to be used? Yes, s[i:] is a slice of the string s starting from index i to the end of the string. This allows DecodeRuneInString to decode the next rune in the string starting from the current index i.
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}

}

// Values enclosed in single quotes are rune literals. We can compare a rune value to a rune literal directly.
func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found Thai character, so sua")
	}
}
