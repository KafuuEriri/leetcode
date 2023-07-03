package main

import "fmt"

func majorityElement(nums []int) int {

	count := 0
	res := 0
	for _, num := range nums {
		if count == 0 {
			res = num
		}
		if res == num {
			count++
		} else {
			count--
		}
	}
	return res
}

func main() {
	// 测试数据
	nums := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}

	// 调用 majorityElement 函数，并打印结果
	result := majorityElement(nums)
	fmt.Println("Result:", result)
}
