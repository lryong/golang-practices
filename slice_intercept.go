// slice_intercept.go
package main

import (
	"fmt"
)

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)

}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	printSlice(numbers)

	fmt.Println("numbers == ", numbers)
	fmt.Println("numbers == ", numbers[1:4])
	fmt.Println("numbers == ", numbers[1:])
	fmt.Println("numbers == ", numbers[3:])

	numbers1 := make([]int, 0, 5)
	printSlice(numbers1)

}
