package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
}

func lengthOfLastWord(s string) int {
	count := 0

	arr := strings.Split(strings.Trim(s, " "), " ")

	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == " " {
			continue
		} else {
			count = len(arr[i])
			break
		}
	}
	return count
}
