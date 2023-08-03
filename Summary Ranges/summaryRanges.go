package main

import "fmt"

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}
	first := nums[0]
	pre := first
	res := make([]string, 0)
	for k, v := range nums {
		// 第一个直接跳过
		if k == 0 {
			continue
		}
		if v-nums[k-1] > 1 {
			if first == pre {
				res = append(res, fmt.Sprintf("%d", pre))
			} else {
				res = append(res, fmt.Sprintf("%d->%d", first, pre))
			}
			first = v
			pre = first
		} else {
			pre = v
		}
	}
	// 遍历完最后的不管是区间还是单独的数字肯定会剩下
	if first == pre {
		res = append(res, fmt.Sprintf("%d", pre))
	} else {
		res = append(res, fmt.Sprintf("%d->%d", first, pre))
	}
	return res
}

func main() {
	nums := []int{0, 2, 3, 4, 6, 8, 9}
	fmt.Println(summaryRanges(nums))
}
