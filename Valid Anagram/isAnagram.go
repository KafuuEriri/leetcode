package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m1 := make(map[rune]int) // 存储s
	m2 := make(map[rune]int) // 存储t
	for i, v := range s {
		if v == rune(t[i]) {
			continue
		}
		if _, ok := m1[rune(t[i])]; ok {
			m1[rune(t[i])] = m1[rune(t[i])] - 1
		} else {
			m2[rune(t[i])] = m2[rune(t[i])] + 1
		}
		if _, ok := m2[v]; ok {
			m2[v] = m2[v] - 1
		} else {
			m1[v] = m1[v] + 1
		}
		if m1[rune(t[i])] == 0 {
			delete(m1, rune(t[i]))
		}
		if m2[v] == 0 {
			delete(m2, v)
		}

	}
	return len(m1) == 0 && len(m2) == 0
}

func isAnagram1(s string, t string) bool {
	// 如果长度不同，直接返回 false
	if len(s) != len(t) {
		return false
	}

	// 使用一个数组来统计字母频率
	count := [26]int{}

	// 遍历字符串 s，统计字母频率
	for _, char := range s {
		count[char-'a']++
	}

	// 遍历字符串 t，减少字母频率
	for _, char := range t {
		count[char-'a']--
	}

	// 检查字母频率是否全部为 0
	for _, c := range count {
		if c != 0 {
			return false
		}
	}

	return true
}

func main() {
	s := "anagram"
	t := "nagaram"
	fmt.Println(isAnagram(s, t))
}
