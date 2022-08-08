package main

import "fmt"

var root *TreeNode = &TreeNode{
	Val: 1,
	Left: &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 8,
			},
		},
		Right: &TreeNode{
			Val: 5,
		},
	},
	Right: &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 6,
		},
	},
}

//var root *TreeNode = &TreeNode{Val: 2}

func main() {
	fmt.Println(isBalanced(root))
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left := maxLeftAndRightDepth(root.Left)
	right := maxLeftAndRightDepth(root.Right)
	height := left - right
	switch height {
	case 1, 0, -1:
		return isBalanced(root.Left) && isBalanced(root.Right)
	default:
		return false
	}
}

func maxLeftAndRightDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxLeftAndRightDepth(root.Left), maxLeftAndRightDepth(root.Right)) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
