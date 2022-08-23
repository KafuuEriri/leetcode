package main

import "fmt"

var prices []int = []int{7, 1, 5, 3, 6, 4}

func main() {
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	res := 0
	// 记录一个最小值
	min := prices[0]
	// 每日遍历如果当前值比最小值还低，变更最小值
	for i := 1; i < len(prices); i++ {
		if min > prices[i] {
			min = prices[i]
		} else { // 如果当前值大于最小值，记录当前值和最小值的差额
			if res < prices[i]-min {
				res = prices[i] - min
			}
		}
	}
	return res
}
