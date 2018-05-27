package main

import (
	"math"
	"fmt"
)
const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

func main() {
	nums := []int{-1, 2, 1, 4}
	result := ThreeSumClosest(nums, 1)
	fmt.Println(result)
}

func ThreeSumClosest(nums []int, target int) int {

	twoSum := AllTwoNumbersSumInArray(nums)

	min :=float64(MaxInt)
	for _, v := range nums {
		for _, u := range twoSum {
			if math.Abs(float64(u+v)-float64(target)) <= min {
				min = math.Abs(float64(u + v + target))
			}

		}
	}
	return int(min)

}
func AllTwoNumbersSumInArray(nums []int) []int {
	var result []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			result = append(result, nums[i]+nums[j])
		}
	}
	return result
}
