// pointer_to_pointer.go
package main

import (
	"fmt"
)

func main() {
	var a int
	var ptr *int
	var pptr **int

	a = 555
	ptr = &a
	pptr = &ptr

	fmt.Printf("变量a = %d\n", a)
	fmt.Printf("指针变量*ptr = %d\n", *ptr)
	fmt.Printf("指针变量**ptr = %d\n", **pptr)
}
