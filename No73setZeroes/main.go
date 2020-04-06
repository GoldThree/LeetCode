package No73setZeroes

func setZeroes(matrix [][]int) {

	row := len(matrix)
	if row == 0 {
		return
	}
	column := len(matrix[0])

	if column == 0 {
		return
	}

	rowZero := map[int]bool{}
	columnZero := map[int]bool{}

	for r := 0; r < row; r++ {
		for c := 0; c < column; c++ {
			if matrix[r][c] == 0 {
				rowZero[r] = true
				columnZero[c] = true
				continue
			}
		}

		for r := 0; r < row; r++ {
			for c := 0; c < column; c++ {
				if rowZero[r] {
					matrix[r][c] = 0
					continue
				}
				if columnZero[c] {
					matrix[r][c] = 0
				}
			}
		}
	}
}
