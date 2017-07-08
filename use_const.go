package main

import "fmt"
import (
	"unsafe"
)

const (
	a1 = "abc"
	b1 = len(a1)
	c1 = unsafe.Sizeof(a1)
)

func main() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str"

	area = LENGTH * WIDTH
	fmt.Printf("面积为：%d", area)
	println()
	println(a, b, c)
	println(a1, b1, c1)

}
