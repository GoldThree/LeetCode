package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 6}
	result := searchInsert(nums, 5)
	fmt.Println(result)
}

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	for i, num := range nums {
		if num >= target {
			return i
		}
		if i == len(nums)-1 {
			return len(nums)
		}
		if num < target && nums[i+1] > target {
			return i + 1
		}
	}
	return 0
}
