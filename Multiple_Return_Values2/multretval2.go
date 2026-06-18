package main

import "fmt"

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	result, err := divide(10, 2)

	result2, err2 := divide(10, 0)

	fmt.Println("result:", result, "error:", err)    // result: 5 error: <nil>
	fmt.Println("result2:", result2, "error:", err2) // result2: 0 error: cannot divide by zero

	_, err3 := divide(10, 0)      // we only care about the error, so we ignore the result with _
	fmt.Println("result3:", err3) // result3: cannot divide by zero

	result4, _ := divide(10, 2) // we only care about the result, so we ignore the error with _
	fmt.Println("result4:", result4)
}
