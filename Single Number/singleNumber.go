package main

import "fmt"

var nums = []int{2, 2, 1}

func main() {
	fmt.Println(singleNumber(nums))
}

func singleNumber(nums []int) int {
	res := 0
	for _, v := range nums {
		res ^= v
	}
	return res
}
