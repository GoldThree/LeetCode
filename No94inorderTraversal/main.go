package No94inorderTraversal


//Definition for a binary tree node.
 type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode}

// var res []int

// func inorderTraversal(root *TreeNode) []int {
// 	res := make([]int, 0)
// 	res = dfs(root, res)
// 	return res
// }


func dfs(root *TreeNode, res []int) []int {
	if root != nil {
		res = dfs(root.Left, res)
		res = append(res, root.Val)
		res = dfs(root.Right, res)
	}

	return res
}

func inorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode

	for root != nil || len(stack) > 0 {
		//push
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		//pop
		pre := len(stack) - 1
		res = append(res, stack[pre].Val)
		root = stack[pre].Right
		stack = stack[:pre]
	}
	return res
}

