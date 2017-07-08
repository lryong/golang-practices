// continue_statement.go

package main

import (
	"fmt"
)

func main() {
	for test := 0; test <= 20; test++ {
		if test%2 == 0 {
			fmt.Printf("The number is an even number %d\n", test)
		}
		if test == 19 {
			fmt.Printf("bypass 19\n")
			continue
		}
	}

}
