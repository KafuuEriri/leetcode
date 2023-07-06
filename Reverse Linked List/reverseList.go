package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode // 前一个节点
	cur := head        // 当前节点
	for cur != nil {
		next := cur.Next // 提前记录下一个节点,下面操作会变更next
		cur.Next = prev  // 指针反转
		prev = cur       // 即将进入下一次循环,上一个节点此时变为当前节点,比如1->2->3,此时1->nil, pre=1
		cur = next       // 当前节点跳入下一个节点,cur=2
	}
	return prev
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

	// 调用 reverseList 函数，并打印结果
	newHead := reverseList(head)

	fmt.Println("\nLinked List after reverse elements with value:")
	printLinkedList(newHead)
}
