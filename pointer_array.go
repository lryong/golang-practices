// pointer_array.go
package main

import (
	"fmt"
)

const MAX int = 3

func main() {
	var a = [3]int{10, 100, 1000}
	var ptr [MAX]*int
	for i := 0; i < MAX; i++ {
		ptr[i] = &a[i] //整数地址赋值给指针数组
	}

	for i := 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr[i])
	}

}
