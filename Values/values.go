package main

import (
	"fmt"
)

func main() {
	fmt.Println("go" + "lang")        // golang. this teaches us that in golang we can add strings like this.
	fmt.Println("1+1 =", 1+1)         // within the Println, when we are adding string and an integer, I must separate this with a comma. And this integer will be converted to a string when getting printed inside println.
	fmt.Println("7.0/3.0 =", 7.0/3.0) // 2.3333333333333335 the answer shows that we are dealing with floating point numbers. If we were to do 7/3, the answer would be 2, because we are using integers.
	fmt.Println(true && false)        // false. this is a boolean expression. the && operator means "and". true and false is false.
	fmt.Println(true || false)        // true. this is a boolean expression as well. the || operator means "or". true or false is true.
	fmt.Println(!true)                // false. Another boolean expression. the ! operator means "not".
}
