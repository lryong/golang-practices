// relational_usage.go

package main

import (
	"fmt"
)

func main() {
	var a int = 21
	var b int = 10

	if a == b {
		fmt.Printf("a equals b\n")
	} else {
		fmt.Printf("a doesn't equal b\n")
	}

	if a < b {
		fmt.Printf("a is smaller than b\n")
	} else {
		fmt.Printf("a is bigger than b\n")
	}

}
