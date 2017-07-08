// assignment_operator.go

package main

import (
	"fmt"
)

func main() {
	var a int = 20
	var b int = 33
	a += 1
	fmt.Printf("a is %d\n", a)
	a -= 1
	fmt.Printf("a is %d\n", a)
	a *= 1
	fmt.Printf("a is %d\n", a)
	a /= 1
	fmt.Printf("a is %d\n", a)
	a %= 1
	fmt.Printf("a is %d\n", a)
	b <<= 1
	fmt.Printf("a is %d\n", b)
	b >>= 1
	fmt.Printf("a is %d\n", b)
	b &= 1
	fmt.Printf("a is %d\n", b)
	b |= 1
	fmt.Printf("a is %d\n", b)
	b ^= 1
	fmt.Printf("a is %d\n", b)

}
