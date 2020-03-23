package main


func generateMatrix(n int) [][]int {

	// 第一行是顺序递增
	// 当某一个数组从左到右进行遍历时候，当遇到index==n-1的时候，index不变，数组的index+1,
	// 当按照列遍历的时候，index不变，当遇到index==n-1的时候，行的index倒叙


	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}

	if n == 1 {
		result[0][0] = 1
		return result
	}

	if n == 2 {
		result[0][0] = 1
		result[0][1] = 2
		result[1][1] = 3
		result[1][0] = 4
		return result
	}

	max := n*n

	row := 0
	index := 0
	minColumn := 0
	maxColumn := n-1
	minRow := 0
	maxRow := n-1

	if (n+1)%2==0 {
		mid := (n+1)/2-1
		result[mid][mid] = max
	}

	for i:=1;i<=max;i++ {

		if index == minColumn && minRow > maxRow {
			result[row][index] = i
			continue
		}

		// 走到左上角 从左往右
		if row == minRow && index < maxColumn {
			result[row][index] = i
			index+=1
			if index == maxColumn {
				minRow += 1
				continue
			}
		}

		// 走到右上角 从上往下
		if index == maxColumn && row < maxRow {
			result[row][index] = i
			row+=1
			if row == maxRow {
				maxColumn-=1
				continue
			}
		}

		// 走到右下角 从右往左
		if row == maxRow && index > minColumn {
			result[row][index] = i
			index -= 1
			if index == minColumn {
				maxRow -= 1
				continue
			}
		}


		// 走到左下角 从下往上
		if index == minColumn && row > minRow {
			result[row][index] = i
			row -=1
			if row == minRow {
				minColumn += 1
				continue
			}
		}
	}
	return result

}