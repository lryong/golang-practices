// send_list.go
package main

import "fmt"

func sum_avg(arr []float32, size int) float32 {
	var sum float32
	for v := 0; v < len(arr); v++ {
		sum += arr[v]
	}
	return sum / float32(size)

}

func main() {
	var test_list = []float32{1.2, 1.1, 2.3309, 1.0, 6.6, 4.4}
	res := sum_avg(test_list, 10)
	fmt.Printf("The avg_result is %f\n", res)
}
