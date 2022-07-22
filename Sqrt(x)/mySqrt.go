package main

import "fmt"

var x int = 8

func main() {
	fmt.Println(mySqrt(x))
}

func mySqrt(x int) int {
	if x == 1 || x == 0 {
		return x
	}
	left, right := 1, x
	for left < right {
		mid := (left + right) / 2
		if mid*mid == x {
			return mid
		}
		if mid*mid > x {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left - 1
}
