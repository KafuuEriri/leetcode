package main

import "fmt"

var head *ListNode = &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 5, Next: &ListNode{Val: 2, Next: nil}}}}}}

func main() {
	result := partition(head, 3)
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
func partition(head *ListNode, x int) *ListNode {
	if head == nil || nil == head.Next {
		return head
	}
	// 定义左链表比x小，右链表比x大，p，q分别是左右链表的尾节点
	var left, right, p, q *ListNode
	// 循环链表
	for ; head != nil; head = head.Next {
		if head.Val < x {
			if left == nil {
				left = &ListNode{Val: head.Val}
				p = left
			} else {
				p.Next = &ListNode{Val: head.Val}
				p = p.Next
			}

		} else {
			if right == nil {
				right = &ListNode{Val: head.Val}
				q = right
			} else {
				q.Next = &ListNode{Val: head.Val}
				q = q.Next
			}
		}
	}
	// 左链表为空直接返回右链表
	if left == nil {
		return right
	}
	// 右链表拼在左链表之后
	if right != nil {
		p.Next = right
	}
	return left
}

type ListNode struct {
	Val  int
	Next *ListNode
}
