package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head.Next
	for fast != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		if fast.Next == nil {
			break
		}
		fast = fast.Next.Next
	}
	return false
}

func main() {
	// 测试用例1: 链表不含环
	node1 := &ListNode{3, nil}
	node2 := &ListNode{2, nil}
	node3 := &ListNode{0, nil}
	node4 := &ListNode{-4, nil}
	//node5 := &ListNode{5, nil}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = nil

	result1 := hasCycle(node1)
	fmt.Println(result1) // false

	// 测试用例2: 链表头节点就是环的起点
	node1.Next = node1
	result2 := hasCycle(node1)
	fmt.Println(result2) // true

	// 测试用例3: 链表中间部分形成一个环
	// node5.Next = node2
	// result3 := hasCycle(node1)
	// fmt.Println(result3) // true
}
