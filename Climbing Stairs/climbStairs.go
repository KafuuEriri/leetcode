package main

import (
	"fmt"
)

var n int = 1

func main() {
	fmt.Println(climbStairs(n))
}

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	// 记录过去爬过的次数
	arr := make([]int, n+1)
	arr[1] = 1 // 第一个台阶一种
	arr[2] = 2 // 第二个台阶两种
	for i := 3; i <= n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr[n]
}
