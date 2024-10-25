package main

// 假设 isBadVersion 函数已经定义
func isBadVersion(version int) bool {
	// 这里假设一个简单的实现，实际应用中这个函数会由外部提供
	return version >= 4
}

func firstBadVersion(n int) int {
	left, right := 1, n
	for left < right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
