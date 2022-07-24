package main

import "fmt"

var head *ListNode = &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}}}

func main() {
	result := deleteDuplicates(head)
	for ; result != nil; result = result.Next {
		fmt.Print(result.Val)
		fmt.Print(",")
	}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	// 用map存放是否重复
	temp := make(map[int]interface{})
	var result, p *ListNode
	for ; head != nil; head = head.Next {
		if temp[head.Val] == nil {
			if result == nil {
				result = &ListNode{Val: head.Val}
				p = result
			} else {
				p.Next = &ListNode{Val: head.Val}
				p = p.Next
			}
			temp[head.Val] = true
		}
	}
	return result
}

type ListNode struct {
	Val  int
	Next *ListNode
}
