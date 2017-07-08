// Type_Switch.go
// switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际存储的变量类型。
package main

import (
	"fmt"
)

func main() {
	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Printf("x type:%T", i)
	case int:
		fmt.Printf("x type is int")
	case float64:
		fmt.Printf("x type is float64")
	case func(int) float64:
		fmt.Printf("x type is func(int)")
	case bool, string:
		fmt.Printf("x type is bool or string")
	default:
		fmt.Printf("x type is unknown")

	}
}
