package No100SameTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	isLeftSame := isSameTree(p.Left, q.Left)
	if !isLeftSame {
		return false
	}

	isRightSame := isSameTree(p.Right, q.Right)
	if !isRightSame {
		return false
	}
	return true

}
