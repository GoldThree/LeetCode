package No64minPaths

func minPathSum(grid [][]int) int {

	row := len(grid)
	if row == 0 {
		return 0
	}

	column := len(grid[0])
	if column == 0 {
		return 0
	}

	dp := make([][]int, row)
	for i:=0;i<row;i++ {
		dp[i] = make([]int, column)
	}

	dp[0][0] = grid[0][0]
	for i := 0; i< row;i++ {
		for j:=0;j< column; j++ {
			if i ==0 && j ==0 {
				continue
			}
			if i == 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
				continue
			}

			if j == 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
				continue
			}
			minNum := min(dp[i-1][j], dp[i][j-1])
			dp[i][j] = grid[i][j] + minNum
		}
	}

	return dp[row-1][column-1]

}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
