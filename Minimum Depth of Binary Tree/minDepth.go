package main

import "fmt"

var root *TreeNode = &TreeNode{
	Val: 2,
	Left: &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 4,
		},
		Right: &TreeNode{
			Val: 5,
		},
	},
	Right: &TreeNode{
		Val: 3,
		Right: &TreeNode{
			Val: 4,
		},
	},
}

func main() {
	fmt.Println(minDepth(root))
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := minDepth(root.Left)
	right := minDepth(root.Right)
	// 分情况，如果左右节点有为空的要计算最大，如果都左右子树都不为空则去最浅的+1
	if left != 0 && right != 0 {
		return min(left, right) + 1
	} else {
		return max(left, right) + 1
	}
}

func min(x, y int) int {
	if x >= y {
		return y
	}
	return x
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
