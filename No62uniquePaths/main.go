package No62uniquePaths

// func uniquePaths(m int, n int) int {

//     // 递归 向右列减1 向下行减1

//     if m == 1 || n == 1 {
//         return 1
//     }
//     result := 0
//     // 向下走一步
//     if m >= 2 {
//         result += uniquePaths(m-1, n)
//     }
//     // 向右走一步
//     if n >= 2 {
//         result += uniquePaths(m, n-1)
//     }
//     return result
// }

func uniquePaths(m int, n int) int {
	paths := make([][]int, n+1)
	for i := range paths {
		paths[i] = make([]int, m+1)
	}
	paths[1][1] = 1


	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if i == 1 && j == 1 {
				continue
			} else {
				paths[i][j] = paths[i-1][j] + paths[i][j-1]
			}
		}
	}
	return paths[n][m]
}
