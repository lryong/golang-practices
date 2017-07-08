// slice_append_copy.go

package main

import (
	"fmt"
)

func main() {
	var numbers []int
	printSlice(numbers)

	numbers = append(numbers, 0) /* 允许追加空切片 */
	printSlice(numbers)

	numbers = append(numbers, 1) /* 向切片添加一个元素 */
	printSlice(numbers)

	numbers = append(numbers, 2, 4, 3) /* 同时添加多个元素 */
	printSlice(numbers)

	numbers1 := make([]int, len(numbers), cap(numbers)*2) /* 创建切片 numbers1 是之前切片的两倍容量*/
	copy(numbers1, numbers)                               /* 拷贝 numbers 的内容到 numbers1 */
	printSlice(numbers1)
}

func printSlice(x []int) {
	fmt.Printf("len = %d cap=%d slice=%v\n", len(x), cap(x), x)

}
