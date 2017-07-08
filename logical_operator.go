// logical_operator.go
package main

// import (
// 	"fmt"
// )

func main() {
	var a bool = true
	var b bool = true
	if !a {
		println("a is False\n")
	} else {
		println("a is True\n")
	}

	if a || b {
		println("a or b  at least one is 0\n")
	} else {
		println("a and b are not 0\n")
	}

	if a && b {
		println("a and b are 1\n")
	} else {
		println("a or b at least one is 0\n")
	}

}
