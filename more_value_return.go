// more_value_return.go

package main

import (
	"fmt"
)

func test(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := test("Lu", "Ruiyong")
	fmt.Println(a, b)
}
