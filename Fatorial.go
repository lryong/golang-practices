// Fibonacci_sequence.go
package main

import (
	"fmt"
)

func main() {
	var i int = 15
	fmt.Printf("%d的阶乘是 %d\n", i, Fibonacci(i))
}

func Fibonacci(x int) (result int) {
	// result := 0
	if x == 0 {
		result = 1
	} else {
		result = x * Fibonacci(x-1)
	}
	return
}
