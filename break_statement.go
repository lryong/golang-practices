// break_statement.go

package main

import (
	"fmt"
)

func main() {
	var a int = 10

	for a < 20 {
		fmt.Printf("This is the test %d\n", a)
		a++
		if a == 18 {
			fmt.Printf("Break Statement\n")
			break
		}
	}

}
