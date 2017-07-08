// range_test1.go
package main

import (
	"fmt"
)

func main() {
	//使用range求一个slice的和，跟使用数组相似
	var test = []int{2, 4, 6}
	sum := 0

	for _, num := range test { //"_"只是可写变量，省略索引
		sum += num
	}

	fmt.Println("sum result is ", sum)

	for i, v := range test {
		if v == 6 {
			fmt.Println("index:", i)
		}
	}
	//range也可以用在map的键值对上。
	kvs := map[string]string{"a": "apple", "b": `banana`}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	//range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身
	for i, c := range "go" {
		fmt.Printf("%d -> %d\n", i, c)
	}

}
