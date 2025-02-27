package main

import "fmt"

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	for n%2 == 0 {
		n = n / 2
	}
	for n%3 == 0 {
		n = n / 3
	}
	for n%5 == 0 {
		n = n / 5
	}
	return n == 1
}

func main() {
	fmt.Println(isUgly(6))  // 输出: true
	fmt.Println(isUgly(14)) // 输出: false
	fmt.Println(isUgly(1))  // 输出: true
	fmt.Println(isUgly(30)) // 输出: true
	fmt.Println(isUgly(21)) // 输出: false
}
