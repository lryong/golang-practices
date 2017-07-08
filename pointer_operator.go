// pointer_operator.go
package main

import (
	"fmt"
)

func main() {
	var a int = 4
	var b int32
	var c float32
	var ptr *int

	fmt.Printf("a is %T\n", a)
	fmt.Printf("b is %T\n", b)
	fmt.Printf("c is %T\n", c)

	ptr = &a
	fmt.Printf("a is %d\n", a)
	fmt.Printf("*ptr is %d\n", *ptr)
	fmt.Printf("ptr(addr) is %d\n", ptr)

}
