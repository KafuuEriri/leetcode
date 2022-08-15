package main

import "fmt"

var numRows int = 5

func main() {
	fmt.Println(generate(numRows))
}

func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	// 循环行
	for i := 0; i < numRows; i++ {
		temp := make([]int, i+1)
		temp[0] = 1
		// 循环列，根据每个数字等于它的左上方与右上方两个数字之和（也就是说，第i行第j个数字等于第i-1行的第j-1个数字与第j个数字的和）
		j := 1
		for ; j < i; j++ {
			temp[j] = res[i-1][j-1] + res[i-1][j]
		}
		if i >= 1 {
			temp[j] = 1
		}
		res[i] = temp
	}
	return res
}
