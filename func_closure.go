// func_closure.go

package main

import (
	"fmt"
)

func getSequence() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func main() {
	nextNumber := getSequence()

	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	nextNumber1 := getSequence()

	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())

}
