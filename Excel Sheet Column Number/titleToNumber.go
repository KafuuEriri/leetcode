package main

import (
	"fmt"
	"math"
)

var charToNumber map[rune]int = map[rune]int{
	'A': 1, 'B': 2, 'C': 3, 'D': 4, 'E': 5, 'F': 6, 'G': 7, 'H': 8, 'I': 9, 'J': 10, 'K': 11, 'L': 12, 'M': 13, 'N': 14,
	'O': 15, 'P': 16, 'Q': 17, 'R': 18, 'S': 19, 'T': 20, 'U': 21, 'V': 22, 'W': 23, 'X': 24, 'Y': 25, 'Z': 26,
}

func titleToNumber(columnTitle string) int {
	res := 0
	length := len(columnTitle)
	for k, v := range columnTitle {
		res += charToNumber[v] * int(math.Pow(float64(26), float64(length-k-1)))
	}
	return res
}

func main() {
	var columnTitle = "ZY"
	res := titleToNumber(columnTitle)
	fmt.Println(res)
}
