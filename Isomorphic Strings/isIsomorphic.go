package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	d1 := [256]int{}
	d2 := [256]int{}
	for k, v := range s {
		s1 := v
		t1 := rune(t[k])
		if d1[s1] != d2[t1] {
			return false
		}
		d1[s1] = k + 1
		d2[t1] = k + 1
	}
	return true
}

func main() {
	s := "bbbaaaba"
	t := "aaabbbba"
	fmt.Println(isIsomorphic(s, t))
}
