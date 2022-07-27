package main

import "fmt"

var p *TreeNode = &TreeNode{
	Val: 1,
	Left: &TreeNode{
		Val: 2,
	},
	Right: &TreeNode{
		Val: 3,
	},
}

var q *TreeNode = &TreeNode{
	Val: 1,
	Left: &TreeNode{
		Val: 2,
	},
	Right: &TreeNode{
		Val: 3,
	},
}

func main() {
	fmt.Println(isSameTree(p, q))
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	} else {
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
