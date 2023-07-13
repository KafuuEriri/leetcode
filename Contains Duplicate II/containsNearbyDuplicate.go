package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i, v := range nums {
		j, ok := m[v]
		if ok {
			if i-j <= k {
				return true
			}
		}
		m[v] = i
	}
	return false
}

func main() {
	var nums = []int{1, 2, 3, 1, 2, 3}
	fmt.Println(containsNearbyDuplicate(nums, 2))
}
