package main

import "fmt"

func addDigits(num int) int {
	if num < 10 {
		return num
	}
	var result int
	for num > 0 {
		dight := num % 10
		result += dight
		num = num / 10
	}
	return addDigits(result)
}

func main() {
	fmt.Println(addDigits(38))
}
