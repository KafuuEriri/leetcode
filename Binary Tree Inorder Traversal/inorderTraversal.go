package main

import "fmt"

func main() {
	fmt.Println(inorderTraversal(&TreeNode{}))
}

// 在二叉树中，首先遍历左子树，然后访问根结点，最后遍历右子树，在遍历左右子树时，仍然采用这种规则，这样的遍历方式叫做二叉树的中序遍历。
func inorderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	//递归实现
	traverse(root, &result)
	return result
}

// 递归方法
func traverse(root *TreeNode, result *[]int) {
	// 左树
	if root.Left != nil {
		traverse(root.Left, result)
	}
	*result = append(*result, root.Val)
	// 右树
	if root.Right != nil {
		traverse(root.Right, result)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
