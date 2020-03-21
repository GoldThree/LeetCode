package main

func spiralOrder(matrix [][]int) []int {
	// 如果列==最大值，列不变，行改变
	// 如果行等于最大值，行不变，列改变
	// 如果列==最小值，列不变，行改变
	// 当最小一列或者最小一行数值全部添加过，则列或者行的最小值加1
	// 当最大一列或者最大一行数据全部添加过，则列或者行的最大值减1

	row := len(matrix)

	result := make([]int, 0)
	if row == 0 {
		return result
	}

	if row == 1 {
		return matrix[0]
	}
	if row == 2 {
		result = append(result, matrix[0]...)
		for i:= len(matrix[1])-1;i>=0;i-- {
			result = append(result, matrix[1][i])
		}
		return result
	}

	// 用递归
	lastIndex := len(matrix[0])-1

	newMartix := make([][]int, 0)

	result = append(result, matrix[0]...)
	// 循环矩阵
	firstNum := make([]int, 0)
	for i:= 1;i<= row-1; i++ {


		if i == row -1 {
			for j := lastIndex;j>=0;j--{
				result = append(result, matrix[i][j])
			}
			continue
		}

		result = append(result, matrix[i][lastIndex])
		if len(matrix[i]) >1 {
			firstNum = append([]int{matrix[i][0]}, firstNum...)
		}

		if len(matrix[i]) > 2 {
			newMartix = append(newMartix, matrix[i][1:lastIndex])
		}

	}
	result = append(result, firstNum...)
	preResult := spiralOrder(newMartix)
	result = append(result, preResult...)
	return result



}

