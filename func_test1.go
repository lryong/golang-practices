package main

import "fmt"

func max(num1, num2 int) int {
	// var result int
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}

func main() {
	fmt.Printf("result is %d", max(4, 5))

}
