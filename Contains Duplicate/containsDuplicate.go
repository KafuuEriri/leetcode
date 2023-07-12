package main

import "fmt"

func containsDuplicate(nums []int) bool {
	m := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			return ok
		} else {
			m[nums[i]] = struct{}{}
		}
	}
	return false
}

func main() {
	var nums = []int{1, 2, 3, 4}
	fmt.Println(containsDuplicate(nums))
}
