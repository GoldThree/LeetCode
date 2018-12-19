package main

import "fmt"

func main() {
	a := []int{0,0,1,1,1,3,3,5,5,6}
	result := removeDuplicates(a)
	fmt.Println(result)

}

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	i, j := 1, 0
	for ; i < len(nums); i++ {
		if nums[i] > nums[j] {
			j = j +1
			nums[j] = nums[i]
		}
	}
	return j+1
}
