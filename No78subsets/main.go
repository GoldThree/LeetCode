package No78subsets

func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	result = append(result, make([]int, 0))
	for i := 0; i < len(nums); i ++ {
		end := len(result)
		curNum := nums[i]
		result = append(result, append(make([]int,0), curNum))
		for j := 1; j < end; j ++ {
			// fmt.Println("111111111111", end)
			tmp := make([]int, len(result[j]))
			copy(tmp, result[j])
			// fmt.Println(tmp, curNum, result)
			cur := append(tmp, curNum)

			result = append(result, cur)
		}
	}
	return result
}

