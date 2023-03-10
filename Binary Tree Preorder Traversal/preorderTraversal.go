package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历,先遍历根节点，再遍历左子树节点，再遍历右子树节点
func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	left := preorderTraversal(root.Left)
	right := preorderTraversal(root.Right)
	res = append(res, left...)
	res = append(res, right...)
	return res
}
