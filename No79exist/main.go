package No79exist

// func exist(board [][]byte, word string) bool {

//     // 回溯 找到选择，添加，返回

//     row := len(board)
//     if row < 1 {
//         return false
//     }

//     column := len(board[0])
//     if column < 1 {
//         return false
//     }

//     usedIndex := make(map[string]bool, 0)

//     for i:=0;i<row;i++ {
//         for j:=0;j<column;j++ {
//             usedIndex = map[string]bool{}
//             if board[i][j] == word[0] {
//                 result := FindWord(board, i, j, 0, word, usedIndex)
//                 fmt.Println(result)
//                 if result == word {
//                     return true
//                 }
//             }
//         }
//     }

//     return false

// }

// func FindWord(board [][]byte, rowIndex, columnIndex, wordIndex int,  needWord string, usedIndex map[string]bool) string {


//     row := len(board)
//     column := len(board[0])

//     current := board[rowIndex][columnIndex]
//     if current != needWord[wordIndex] {
//         return ""
//     }

//      if wordIndex == len(needWord)-1  {
//         return string(current)
//     }

//     wordIndex+=1
//     usedIndex[getIndexStr(rowIndex, columnIndex)] = true


//     // 向右
//     if !usedIndex[getIndexStr(rowIndex, columnIndex+1)] && columnIndex+1 < column {
//         temp := FindWord(board, rowIndex, columnIndex+1, wordIndex, needWord,usedIndex)
//         if temp != "" {
//             return string(current)+temp
//         }
//         if usedIndex[getIndexStr(rowIndex, columnIndex+1)]{
//             usedIndex[getIndexStr(rowIndex, columnIndex+1)] = false
//         }
//     }

//     // 向左
//     if !usedIndex[getIndexStr(rowIndex, columnIndex-1)] && columnIndex-1>=0 {
//         temp := FindWord(board, rowIndex, columnIndex-1, wordIndex, needWord,usedIndex)
//         if temp != "" {
//             return string(current)+temp
//         }
//         if usedIndex[getIndexStr(rowIndex, columnIndex-1)]{
//             usedIndex[getIndexStr(rowIndex, columnIndex-1)] = false
//         }
//     }

//     // 向下
//      if !usedIndex[getIndexStr(rowIndex+1, columnIndex)] && rowIndex+1<row {
//         temp := FindWord(board, rowIndex+1, columnIndex, wordIndex, needWord,usedIndex)
//         if temp != "" {
//             return string(current)+temp
//         }

//         if usedIndex[getIndexStr(rowIndex+1, columnIndex)]{
//             usedIndex[getIndexStr(rowIndex+1, columnIndex)] = false
//         }
//     }

//     // 向上
//      if !usedIndex[getIndexStr(rowIndex-1, columnIndex)] && rowIndex-1 >= 0 {
//         temp := FindWord(board, rowIndex-1, columnIndex, wordIndex, needWord,usedIndex)
//         if temp != "" {
//             return string(current)+temp
//         }

//         if usedIndex[getIndexStr(rowIndex-1, columnIndex)]{
//             usedIndex[getIndexStr(rowIndex-1, columnIndex)] = false
//         }
//     }

//     return ""

// }

// func getIndexStr(rowIndex, columnIndex int) string {

//     return fmt.Sprintf("%d.%d", rowIndex, columnIndex)
// }

// 参考
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(board, i, j, word, 0) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, i int, j int, word string, k int) bool {
	if board[i][j] != word[k] {
		return false
	}
	if k == len(word)-1 {
		return true
	}
	temp := board[i][j]
	board[i][j] = byte(' ')
	if 0 <= i-1 && dfs(board, i-1, j, word, k+1) {
		return true
	}
	if i+1 < len(board) && dfs(board, i+1, j, word, k+1) {
		return true
	}
	if 0 <= j-1 && dfs(board, i, j-1, word, k+1) {
		return true
	}
	if j+1 < len(board[0]) && dfs(board, i, j+1, word, k+1) {
		return true
	}
	board[i][j] = temp
	return false
}

