package No75sortColors

func sortColors(nums []int)  {

	if len(nums) == 0 {
		return
	}

	red := 0
	white := 0
	blue := 0
	for i:= 0; i<len(nums); i++ {
		if nums[i] == 0 {
			red += 1
			continue
		}

		if nums[i] == 1 {
			white += 1
			continue
		}

		if nums[i] == 2 {
			blue += 1
			continue
		}
	}

	start := 0

	for r := 0; r < red; r++ {
		nums[start] = 0
		start+=1
	}

	for w := 0; w < white; w++ {
		nums[start] = 1
		start+=1
	}

	for b := 0; b < blue; b++ {
		nums[start] = 2
		start+=1
	}

}
