package main

import "fmt"

var temp map[int]string = map[int]string{1: "A", 2: "B", 3: "C", 4: "D", 5: "E", 6: "F", 7: "G", 8: "H", 9: "I", 10: "J", 11: "K", 12: "L",
	13: "M", 14: "N", 15: "O", 16: "P", 17: "Q", 18: "R", 19: "S", 20: "T", 21: "U", 22: "V", 23: "W", 24: "X",
	25: "Y", 26: "Z",
}

func convertToTitle(columnNumber int) string {
	res := ""
	for columnNumber > 0 {
		m := columnNumber % 26
		if m == 0 {
			m = 26
			// 注意转换其实是1-26,而不是0-25,如果刚好26整除,应该-1
			columnNumber--
		}
		res = temp[m] + res
		columnNumber = columnNumber / 26
	}
	return res
}

func main() {
	columnNumber := 701
	res := convertToTitle(columnNumber)
	fmt.Println(res)
}
