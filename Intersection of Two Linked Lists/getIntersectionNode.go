package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p1, p2 := headA, headB
	for p1 != p2 {
		if p1 == nil {
			p1 = headB
		} else {
			p1 = p1.Next
		}
		if p2 == nil {
			p2 = headA
		} else {
			p2 = p2.Next
		}
	}
	return p1
}

func main() {
	// 创建链表A
	node1 := &ListNode{Val: 2}
	node2 := &ListNode{Val: 6}
	node3 := &ListNode{Val: 4}
	// node4 := &ListNode{Val: 4}
	// node5 := &ListNode{Val: 5}
	node1.Next = node2
	node2.Next = node3
	// node3.Next = node4
	// node4.Next = node5

	// 创建链表B
	node6 := &ListNode{Val: 1}
	node7 := &ListNode{Val: 5}
	node6.Next = node7
	//node7.Next = node3

	// 获取两个链表的交点
	intersection := getIntersectionNode(node1, node6)

	// 打印交点的值（如果存在）
	if intersection != nil {
		println(intersection.Val)
	} else {
		println("No intersection")
	}
}
