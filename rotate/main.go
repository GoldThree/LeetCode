package rotate

// func rotate(matrix [][]int)  {

//     row := len(matrix)
//     if row <= 1 {
//         return
//     }

//     tempMatrix := make([][]int, row)
//     for k:=0;k<row;k++ {
//         tempMatrix[k] = make([]int, row)
//     }

//     for i, ma := range matrix {
//         for j, value := range ma {
//             tempMatrix[i][j] = value
//         }
//     }


//     for i:= 0;i<row;i++ {
//         needJ := i
//         for j:=0;j<row;j++ {

//             needI := row-1-j
//              matrix[i][j] = tempMatrix[needI][needJ]
//         }
//     }



//     // 00->02  00-> 20
//     // 01->12  01-> 10
//     // 02->22  02-> 00

//     // 10->01  10-> 21
//     // 11->11  11-> 11
//     // 12->21  12-> 01

//     // 20->00  20-> 22
//     // 21->10  21-> 12
//     // 22->20  22-> 02

// }

func rotate(matrix [][]int)  {
	if len(matrix) == 0 {
		return
	}
	m, n := len(matrix), len(matrix[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if j > i {      // 当 j > i 时进行转置
				matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
			}
			if j >= n / 2 {     // 遍历到每一行的后半部分时进行反转
				matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
			}
		}
	}
}
