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
	fmt.Println(isSymmetric(root))
}

func isSymmetric(root *TreeNode) bool {
	// 左树和右数相等，直接沿用上一题的树比较
	if root == nil {
		return true
	}
	return isSymmetricTree(root.Left, root.Right)
}

func isSymmetricTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	} else {
		return isSymmetricTree(p.Left, q.Right) && isSymmetricTree(p.Right, q.Left)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
