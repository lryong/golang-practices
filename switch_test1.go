// switch_test1.go

package main

import (
	"fmt"
)

func main() {
	var test string = "lel"

	switch {
	case test == "lel":
		fmt.Printf("A")
	case test == "ddd":
		fmt.Printf("A")
	default:
		fmt.Printf("haha")
	}
}
