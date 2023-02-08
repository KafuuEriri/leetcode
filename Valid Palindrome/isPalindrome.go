package main

import (
	"fmt"
	"regexp"
	"strings"
)

var s string = "race a car"

func main() {
	fmt.Println(isPalindrome(s))
}

func isPalindrome(s string) bool {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return false
	}
	processedString := strings.ToLower(reg.ReplaceAllString(s, ""))
	l := len(processedString)
	for i := 0; i < l; i++ {
		if processedString[i] != processedString[l-1-i] {
			return false
		}
	}
	return true
}
