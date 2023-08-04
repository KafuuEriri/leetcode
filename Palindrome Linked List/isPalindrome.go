package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// func isPalindrome(head *ListNode) bool {
// 	if head == nil {
// 		return true
// 	}
// 	nums := make([]int, 0)
// 	for head != nil {
// 		nums = append(nums, head.Val)
// 		head = head.Next
// 	}
// 	l := len(nums)
// 	for i := 0; i <= (l / 2); i++ {
// 		if nums[i] != nums[l-1-i] {
// 			return false
// 		}
// 	}
// 	return true
// }

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	l := 0
	for cur := head; cur != nil; cur = cur.Next {
		l++
	}
	// 反转前半段链表
	step := l / 2
	pre := &ListNode{}
	cur := head
	for i := 1; i <= step; i++ {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	mid := cur
	var left, right *ListNode = pre, nil
	if l%2 == 0 { // 偶数
		right = mid
	} else { // 奇数
		right = mid.Next
	}
	for left != nil && right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}
	return true
}

func main() {
	// 构建测试链表
	// 1 -> 2 -> 6 -> 3 -> 4 -> 5 -> 6
	head := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	// node3 := &ListNode{Val: 2}
	// node4 := &ListNode{Val: 3}
	// node5 := &ListNode{Val: 2}
	// node6 := &ListNode{Val: 2}
	// node7 := &ListNode{Val: 1}

	head.Next = node2
	// node2.Next = node3
	// node3.Next = node4
	// node4.Next = node5
	// node5.Next = node6
	// node6.Next = node7

	fmt.Println(isPalindrome(head))
}
