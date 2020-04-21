package No106buildTreeII

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 以 后序遍历 root (index=length-1) 寻找 中序遍历 root (index=k)
func buildTree(inorder []int, postorder []int) *TreeNode {

	l := len(postorder) - 1
	for k := range inorder {
		if inorder[k] == postorder[l] {
			return &TreeNode{
				Val:   inorder[k],
				Left:  buildTree(inorder[:k], postorder[:k]),
				Right: buildTree(inorder[k+1:], postorder[k:l]),
			}
		}
	}
	return nil
}
