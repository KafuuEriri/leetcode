package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left = invertTree(root.Left)
	root.Right = invertTree(root.Right)
	temp := root.Left
	root.Left = root.Right
	root.Right = temp

	return root
}

// 中序遍历函数，用于验证反转结果
func inorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	inorderTraversal(root.Left)
	fmt.Printf("%d ", root.Val)
	inorderTraversal(root.Right)
}

func main() {
	// 构建二叉树
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}

	// 打印原始二叉树中序遍历结果
	fmt.Println("Original Tree Inorder Traversal:")
	inorderTraversal(root)
	fmt.Println()

	// 反转二叉树
	invertedRoot := invertTree(root)

	// 打印反转后的二叉树中序遍历结果
	fmt.Println("Inverted Tree Inorder Traversal:")
	inorderTraversal(invertedRoot)
	fmt.Println()
}
