// switch_statement.go
package main

import (
	"fmt"
)

func main() {
	var test string = "lryong"
	var out string = "test"
	switch test {
	case "lryoung":
		out = "so bad"
	case "sdfk":
		out = "not so bad"
	default:
		out = "I'm type for test"
	case "lryong":
		out = "Good"
	}
	fmt.Printf("Your are %s\n", out)

}
