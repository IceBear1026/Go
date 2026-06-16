package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s) // constant. this is how we declare a constant in golang. we use the const keyword.

	const n = 500000000 // this is a constant integer. we can also declare constants of other types and also inside a function.

	const d = 3e20 / n // this is a constant floating point number. we can also do math with constants.
	fmt.Println(d)     // 6e+11. this is the result of the math operation we did with constants.

	fmt.Println(int64(d)) // 600000000000. we can also convert constants to other types.
	fmt.Println(math.Sin(n))
	fmt.Println(float64(math.Sin(n))) // -0.28470407323754404. we can also use constants with functions from other packages.
}
