package main

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}else {
			return -1
		}
	}
	if len(nums) == 2 {
		if nums[0] == target {
			return 0
		}
		if nums[1] == target {
			return 1
		}
		return -1
	}

	var medianIndex int
	length := len(nums)
	if length%2 != 0 {
		medianIndex = (length-1)/2
	} else {
		medianIndex = (length/2) - 1
	}

	preNums := nums[:medianIndex]
	nextNums := nums[medianIndex:]
	index := search(preNums, target)
	if index != -1 {
		return index
	} else {
		index := search(nextNums, target)
		if index != -1 {
			return medianIndex+index
		} else {
			return index
		}
	}
	return -1
}
