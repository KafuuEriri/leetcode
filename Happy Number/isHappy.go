package main

import "fmt"

func isHappy(n int) bool {

	fast, slow := n, n
	for {
		fast = getNext(getNext(fast))
		slow = getNext(slow)
		if fast == slow {
			break
		}
	}
	return fast == 1
}

func getNext(n int) int {
	res := 0
	for n > 0 {
		t := n % 10
		res += t * t
		n = n / 10
	}
	return res
}

func main() {
	n := 19
	fmt.Println(isHappy(n))
}
