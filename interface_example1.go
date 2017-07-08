// interface_example1.go
package main

import (
	"fmt"
)

type Phone interface {
	call()
}

type NokiaPhone struct{}

func (phone NokiaPhone) call() {
	fmt.Println("I'm Nokia,I can call you!")
}

type IPhone struct{}

func (phone IPhone) call() {
	fmt.Println("I'm iphone,I can call you!")
}

func main() {
	var test_phone Phone
	test_phone = new(NokiaPhone)
	test_phone.call()

	test_phone = new(IPhone)
	test_phone.call()
}
