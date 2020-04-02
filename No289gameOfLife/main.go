package No289gameOfLife

func gameOfLife(board [][]int)  {

	row := len(board)
	if row == 0 {
		return
	}

	column := len(board[0])
	if column == 0 {
		return
	}

	result := make([][]int, row)
	for i, _ := range result{
		result[i] = make([]int, column)
	}

	// oldBoard := board[:]

	for i:=0;i<row;i++ {
		for j:=0;j<column;j++ {

			result[i][j] = changeValue(i, j, row, column, board)
		}
	}

	// fmt.Println(result)

	// board = result[:]
	for i, _ := range result{
		for j:=0;j<=column-1;j++ {
			board[i][j] = result[i][j]
		}
	}
	return

}

func changeValue(i, j, row, column int, board [][]int) int {

	aroundData := make([]int, 0)

	// 上层三个
	// 下层三个
	// 本层两个都要考虑

	// 先考虑左右，如果左右有的话，那么上下层如果存在的话，都是一致的

	// 如果上层存在的话
	if i-1 >= 0 {
		aroundData = append(aroundData, board[i-1][j])
	}
	// 如果下层存在的话
	if i+1 <= row-1 {
		aroundData = append(aroundData, board[i+1][j])
	}

	// 考虑左边
	if j-1 >= 0 {
		aroundData = append(aroundData, board[i][j-1])
		// 如果上层存在的话
		if i-1 >= 0 {
			aroundData = append(aroundData, board[i-1][j-1])
		}
		// 如果下层存在的话
		if i+1 <= row-1 {
			aroundData = append(aroundData,  board[i+1][j-1])
		}
	}

	// 考虑右边
	if j+1 <= column-1 {
		aroundData = append(aroundData, board[i][j+1])
		// 如果上层存在的话
		if i-1 >= 0 {
			aroundData = append(aroundData, board[i-1][j+1])
		}
		// 如果下层存在的话
		if i+1 <= row-1 {
			aroundData = append(aroundData,  board[i+1][j+1])
		}
	}

	deadNum := 0
	liveNum := 0
	for _, v := range aroundData {
		if v == 0 {
			deadNum += 1
			continue
		}
		liveNum += 1
	}

	currentValue := board[i][j]
	if currentValue == 1{
		if liveNum < 2 {
			currentValue = 0
		}
		if liveNum > 3 {
			currentValue = 0
		}
	}

	if currentValue == 0 {
		if liveNum == 3 {
			currentValue = 1
		}
	}

	return currentValue
}
