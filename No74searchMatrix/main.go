package No74searchMatrix

func searchMatrix(matrix [][]int, target int) bool {

	row := len(matrix)
	if row == 0{
		return false
	}
	column := len(matrix[0])
	if column == 0 {
		return false
	}

	for i:=0;i<row;i++ {

		if i == row-1 && target <matrix[i][0]{
			return false
		}

		if i == row-1 && target >= matrix[i][0]{
			for j:=0;j<column;j++ {
				if matrix[i][j] == target {
					return true
				}
			}
			return false
		}

		if matrix[i][0]<=target && target < matrix[i+1][0] {
			for j:=0;j<column;j++ {
				if matrix[i][j] == target {
					return true
				}
			}
			return false
		}
	}

	return false

}
