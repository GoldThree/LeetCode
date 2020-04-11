package No96numTrees

// 对于[1, n]，如果以 i(1<i<n) 为root，那么根据BST的性质，[1, i-1] 的元素落在其左子树，[i+1, n]落在右子树，
//这就得到了一个递归关系，以i为root的bst数量为左子树和右子树进行任意组合的总数量


func numTrees(n int) int {
	G := make([]int, n+1)
	G[0], G[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			G[i] += G[j-1] * G[i-j]
		}
	}
	return G[n]
}




