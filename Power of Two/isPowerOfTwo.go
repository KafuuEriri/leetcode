package main

import "fmt"

// func isPowerOfTwo(n int) bool {
// 	if n == 1 {
// 		return true
// 	}
// 	for i := 0; i < 32; i++ {
// 		if 2<<i == n {
// 			return true
// 		}
// 	}
// 	return false
// }

func isPowerOfTwo(n int) bool {
	if n < 1 {
		return false
	}
	return (2<<31)%n == 0
}

func main() {
	n := 1
	fmt.Println(isPowerOfTwo(n))
}
