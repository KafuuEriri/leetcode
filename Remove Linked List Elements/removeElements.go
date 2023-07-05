package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	cur := dummyHead
	for cur.Next != nil {
		next := cur.Next
		if next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummyHead.Next
}

func printLinkedList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
	fmt.Println("nil")
}

func main() {
	// 构建测试链表
	// 1 -> 2 -> 6 -> 3 -> 4 -> 5 -> 6
	head := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 6}
	node4 := &ListNode{Val: 3}
	node5 := &ListNode{Val: 4}
	node6 := &ListNode{Val: 5}
	node7 := &ListNode{Val: 6}

	head.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = node7

	fmt.Println("Original Linked List:")
	printLinkedList(head)

	// 调用 removeElements 函数，并打印结果
	valToRemove := 6
	newHead := removeElements(head, valToRemove)

	fmt.Println("\nLinked List after removing elements with value", valToRemove, ":")
	printLinkedList(newHead)
}
