package main

import "fmt"

var rowIndex int = 33

func main() {
	fmt.Println(getRow(rowIndex))
}

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	// 题目范围rowIndex 0-33,应构造34层
	triangle := generate(34)
	return triangle[rowIndex]
}

// 题目给出了行范围，不如直接按范围构造出来直接输出
func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	// 循环行
	for i := 0; i < numRows; i++ {
		temp := make([]int, i+1)
		temp[0] = 1
		// 循环列
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
