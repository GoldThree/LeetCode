package No1162maxDistance

// 时间超时
// func maxDistance(grid [][]int) int {

//     //找到所有的陆地，然后每个海洋都和陆地进行计算，找到最远的海洋位置

//     if len(grid) == 1 && len(grid[0]) == 1 {
//         return -1
//     }

//     row := len(grid)

//     landArray := make([][]int, 0)
//     seaArray := make([][]int, 0)

//     for i:= 0;i<row;i++ {
//         for j:=0;j<row; j++ {
//             if grid[i][j] == 0 {
//                 seaArray = append(seaArray, []int{i, j})
//                 continue
//             }
//             if grid[i][j] == 1 {
//                 landArray = append(landArray, []int{i, j})
//                 continue
//             }
//         }
//     }

//     if len(landArray) == 0 || len(seaArray) == 0 {
//         return -1
//     }

//     result := 0
//     for _, sea := range seaArray {
//         // 对于每一个海洋来说，每一个海洋距离最近的陆地就是结果，然后从这些距离最近的中选出最大的就是结果
//         minDistance := row*row
//         for _, land := range landArray {
//             tempResult := abs(sea[0]-land[0])+abs(sea[1]-land[1])
//             if tempResult <= minDistance  {
//                 minDistance = tempResult
//             }

//         }
//         if minDistance >= result {
//             result = minDistance
//         }
//     }
//     return result

// }

// func abs(n int) int {
//     if n < 0 {
//         return -n
//     }
//     return n
// }

// 广度优先遍历

// 每一个点的四周位置
var d = [4][2]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}

func maxDistance(grid [][]int) int {
	data := make([][]int, 0)
	row, column := len(grid), len(grid[0])
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if grid[i][j] == 1 {
				data = append(data, []int{i, j})
			}
		}
	}

	if row*column == len(data) || 0 == len(data) {
		return -1
	}
	res := 0
	for len(data) > 0 {
		dataLen := len(data)
		for k := 0; k < dataLen; k++ {
			pre := data[0]
			data = data[1:]
			for _, it := range d {
				x, y := it[0] + pre[0], it[1] + pre[1]
				if x >= 0 && x < row && y >= 0 && y < column && grid[x][y] == 0 {
					grid[x][y] = 1
					data = append(data, []int{x, y})
				}
			}
		}
		res++
	}
	return res - 1
}
