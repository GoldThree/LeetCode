package main

import "fmt"

func main() {
	a := []int{-2,1,-3,4,-1,2,1,-5,4}
	fmt.Println(maxSubArray(a))
}
func maxSubArray(nums []int) int {

	length := len(nums)
	if 0 == length {
		return 0
	}

	sum := nums[0]
	maxSum := sum
	for i := 1; i < length; i++ {
		if sum < 0 {
			sum = nums[i]
		} else {
			sum = nums[i] + sum
		}

		if maxSum < sum {
			maxSum = sum
		}
	}

	return maxSum
}
