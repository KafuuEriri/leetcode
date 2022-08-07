package main

import "fmt"

var nums []int = []int{-10, -3, 0, 5, 9}

func main() {
	result := sortedArrayToBST(nums)
	fmt.Println(result.Val)
	fmt.Println(result.Left)
	fmt.Println(result.Right)
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return parseTree(nums, 0, len(nums)-1)
}

func parseTree(nums []int, i, j int) *TreeNode {
	if i > j {
		return nil
	}
	mid := (i + j) / 2
	var root *TreeNode = &TreeNode{}
	root.Val = nums[mid]
	root.Left = parseTree(nums, i, mid-1)
	root.Right = parseTree(nums, mid+1, j)
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
