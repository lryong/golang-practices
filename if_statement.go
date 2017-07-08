// if_statement.go

package main

import (
	"fmt"
)

func main() {
	var a string = "lryong"
	if a == "string" {
		fmt.Printf("input is string")
	} else {
		if a == "lryong" {
			fmt.Printf("input is lryong")
		} else {
			fmt.Printf("input is not string")
		}
	}

}
