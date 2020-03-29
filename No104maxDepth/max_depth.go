package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	subLeftDepth := 0
	subRightDepth := 0
	if root.Left != nil {
		subLeftDepth = maxDepth(root.Left)
		subLeftDepth += 1
	}

	if root.Right != nil {
		subRightDepth = maxDepth(root.Right)
		subRightDepth += 1
	}

	result := subLeftDepth
	if subLeftDepth < subRightDepth {
		result = subRightDepth
	}
	return result
}
