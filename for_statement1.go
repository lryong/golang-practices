// for_statement1.go

package main

import (
	"fmt"
)

func main() {
	var b int = 15
	var a int

	numbers := [6]int{1, 2, 3, 5}

	for a := 0; a < 10; a++ {
		fmt.Printf("a is %d\n", a)
	}

	fmt.Printf("statement1 is ok\n")·

	for a < b {
		a++ //a只作用于当前域
		fmt.Printf("A is %d\n", a)
	}

	for i, x := range numbers {
		fmt.Printf("第 %d位x的值= %d\n", i, x)
	}
}
