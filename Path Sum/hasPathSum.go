package main

import "fmt"

var root *TreeNode = &TreeNode{
	Val: 5,
	Left: &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 11,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: 2,
			},
		},
	},
	Right: &TreeNode{
		Val: 8,
		Left: &TreeNode{
			Val: 13,
		},
		Right: &TreeNode{
			Val: 4,
			Right: &TreeNode{
				Val: 1,
			},
		},
	},
}

func main() {
	fmt.Println(hasPathSum(root, 22))
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil && root.Val == targetSum {
		return true
	}
	// 递归遍历所有路径，target减去每个节点指最后匹配
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
