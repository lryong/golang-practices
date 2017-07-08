// get_list_values.go

package main

import (
	"fmt"
)

func main() {
	var n [10]int
	var i int

	for i = 0; i < 10; i++ {
		n[i] = i + 100
		fmt.Printf("Element[%d] = %d\n", i, n[i])
	}

	// for j = 0; j < 10; j++ {
	// 	fmt.Printf("Element[%d] = %d\n", j, n[j])
	// }

}
