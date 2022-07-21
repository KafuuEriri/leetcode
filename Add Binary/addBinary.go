package main

import (
	"fmt"
	"strconv"
)

var a, b string = "1", "11"

func main() {
	fmt.Println(addBinary(a, b))
}

func addBinary(a string, b string) string {
	var result string
	var carry uint8 // 进位
	i := len(a) - 1
	j := len(b) - 1
	for i >= 0 && j >= 0 {
		result = strconv.Itoa(int((a[i]-'0'+(b[j]-'0')+carry)%2)) + result
		carry = (a[i] - '0' + (b[j] - '0') + carry) / 2
		i--
		j--
	}
	// i超比j长
	for i >= 0 {
		result = strconv.Itoa(int((a[i]-'0'+carry)%2)) + result
		carry = (a[i] - '0' + carry) / 2
		i--
	}
	// j比i长
	for j >= 0 {
		result = strconv.Itoa(int((b[j]-'0'+carry)%2)) + result
		carry = (b[j] - '0' + carry) / 2
		j--
	}
	// 需要进位
	if carry > 0 {
		result = "1" + result
	}
	return result
}
