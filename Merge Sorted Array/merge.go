package main

import "fmt"

var m, n int = 2, 4
var nums1, nums2 []int = []int{1, 3, 0, 0, 0, 0}, []int{2, 4, 5, 6}

func main() {
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	// 将最大的放到最后
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			k--
			i--
		} else {
			nums1[k] = nums2[j]
			k--
			j--
		}
	}
	// 如果数组1全部大于数组2
	for j >= 0 {
		nums1[k] = nums2[j]
		k--
		j--
	}
}
