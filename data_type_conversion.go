// data_type_conversion.go
package main

import (
	"fmt"
)

func main() {
	var sum int = 123
	var divisor int = 3
	var res float32
	res = float32(sum) / float32(divisor)
	fmt.Printf("The division result is %f", res)
}
