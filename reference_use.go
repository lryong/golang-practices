// reference_use.go

package main

import (
	"fmt"
)

func main() {
	var a int = 100
	var b int = 200
	fmt.Printf("交换前a的值为%d\n", a)
	fmt.Printf("交换前b的值为%d\n", b)

	swap(&a, &b)
	// swap1(a, b)
	fmt.Printf("交换后a的值为%d\n", a)
	fmt.Printf("交换后b的值为%d\n", b)
}

func swap(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp

}

func swap1(x, y int) {
	var temp int
	temp = x
	x = y
	y = temp
}

func swap2(x *int, y *int) {
	var temp int
	temp = x
	x = y
	y = temp
}
