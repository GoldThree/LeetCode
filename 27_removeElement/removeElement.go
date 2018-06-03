package main

import "fmt"

func  main(){
	ints := []int{1,2,3,4,5}
	result := removeElement(ints, 2	)
	fmt.Println(result)
}

func removeElement(nums []int, val int) int {
	// j指向最后一个不为val的位置
	// i指向第一个值为val的位置
	i, j := 0, len(nums)-1
	for {
		for i < len(nums) && nums[i] != val {
			i++
		}

		for j >= 0 && nums[j] == val {
			j--
		}

		if i >= j {
			break
		}

		nums[i], nums[j] = nums[j], nums[i]
	}

	return i
}

