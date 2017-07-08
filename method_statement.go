// method_statement.go

package main

import (
	"fmt"
)

type Circle struct {
	radius float64
}

func main() {
	var c1 Circle
	c1.radius = 10.00
	fmt.Println("Area of Circle(c1) = ", c1.getArea())
}

func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}
