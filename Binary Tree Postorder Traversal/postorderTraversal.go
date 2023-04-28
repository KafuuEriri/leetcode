package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 后序遍历:先遍历左子树节点，再遍历右子树节点，最后遍历根节点
func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	traverse(root, &res)
	return res
}

func traverse(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	traverse(root.Left, res)
	traverse(root.Right, res)
	*res = append(*res, root.Val)
}
