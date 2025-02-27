package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	var result []string
	if root == nil {
		return result
	}
	dfs(root, "", &result)
	return result
}

func dfs(node *TreeNode, path string, result *[]string) {
	if path == "" {
		path = strconv.Itoa(node.Val)
	} else {
		path += "->" + strconv.Itoa(node.Val)
	}
	// 叶子节点
	if node.Left == nil && node.Right == nil {
		*result = append(*result, path)
		return
	}
	// 左子树
	if node.Left != nil {
		dfs(node.Left, path, result)
	}
	// 右子树
	if node.Right != nil {
		dfs(node.Right, path, result)
	}
}

func main() {
	// 示例 1
	root1 := &TreeNode{Val: 1}
	root1.Left = &TreeNode{Val: 2}
	root1.Right = &TreeNode{Val: 3}
	root1.Left.Right = &TreeNode{Val: 5}
	fmt.Println(binaryTreePaths(root1)) // 输出: ["1->2->5", "1->3"]

	// 示例 2
	root2 := &TreeNode{Val: 1}
	fmt.Println(binaryTreePaths(root2)) // 输出: ["1"]

	// 示例 3
	root3 := &TreeNode{Val: 1}
	root3.Left = &TreeNode{Val: 2}
	root3.Left.Left = &TreeNode{Val: 3}
	root3.Left.Right = &TreeNode{Val: 4}
	root3.Right = &TreeNode{Val: 5}
	fmt.Println(binaryTreePaths(root3)) // 输出: ["1->2->3", "1->2->4", "1->5"]
}
