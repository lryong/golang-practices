// cal_usage.go

package main

import (
	"fmt"
)

func main() {

	var a int = 21
	var b int = 10
	var c int

	c = a + b
	fmt.Printf("c is %d\n", c)
	c = a - b
	fmt.Printf("c is %d\n", c)
	c = a * b
	fmt.Printf("c is %d\n", c)
	c = a / b
	fmt.Printf("c is %d\n", c)
	c = a % b
	fmt.Printf("c is %d\n", c)
	c++
	fmt.Printf("c is %d\n", c)
	c--
	fmt.Printf("c is %d\n", c)
}
