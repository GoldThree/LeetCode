package main

import "fmt"

func rotate(matrix [][]int)  {
	length := len(matrix)
	if length == 0 || length == 1 {
		return
	}

	maxIndex := length - 1
	tempMap := map[string]int{}
	for i, nums := range matrix {
		fmt.Println(i, nums)
		for j, num := range nums {
			key := fmt.Sprintf("%d.%d", j, maxIndex - i)
			tempMap[key] = num
		}
	}
	for i:= 0; i< length; i++ {
		for j:= 0; j < length; j++ {
			key := fmt.Sprintf("%d.%d", i, j)
			matrix[i][j] = tempMap[key]
		}
	}
	return

}
