package main

import "fmt"

func missingNumber(nums []int) int {
	n := len(nums)
	// 理论总和
	sum := n * (n + 1) / 2
	// 实际综合
	realSum := 0
	for _, v := range nums {
		realSum += v
	}
	return sum - realSum
}

func main() {
	// 测试用例
	nums := []int{3, 0, 1, 4, 2, 6, 7, 8}
	missing := missingNumber(nums)
	fmt.Println("缺失的数字是:", missing)
}
