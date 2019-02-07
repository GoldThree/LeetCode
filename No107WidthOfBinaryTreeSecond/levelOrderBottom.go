package No107WidthOfBinaryTreeSecond

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	result := make([][]int, 0)
	GenResult(&result, root, 0)
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}

func GenResult(result *[][]int, root *TreeNode, height int) {
	if root == nil {
		return
	}
	if height >= len(*result) {
		*result = append(*result, make([]int, 0))
	}

	(*result)[height] = append((*result)[height], root.Val)
	GenResult(result, root.Left, height+1)
	GenResult(result, root.Right, height+1)
}
