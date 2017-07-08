// nil_pointer.go
package main

import (
	"fmt"
)

func main() {
	var ptr *int
	/*空指针判断*/
	if ptr == nil { //ptr空指针
		fmt.Println("Has not assigned pointer")
	} else {
		fmt.Println("Has assigned pointer")
	}
	fmt.Printf("ptr's value is %d\n", ptr) //空指针
}
