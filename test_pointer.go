// test_pointer.go
package main

import (
	"fmt"
)

func main() {
	var a int = 20 //声明实际变量
	var ip *int    //声明指针变量

	ip = &a //指针变量的存储地址
	fmt.Printf("a's addr is %x\n", &a)
	fmt.Printf("a's store addr is %x\n", ip) //指针变量的存储地址
	fmt.Printf("*ip 's value is %d\n", *ip)  //使用指针访问值

}
