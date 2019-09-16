package main

import "fmt"

func main()  {
	nums := []int{-1, 2, 1, -4}
	fmt.Println(threeSumClosest(nums, 2))
}

func threeSumClosest(nums []int, target int) int {
	length := len(nums)
	const MaxUint = ^uint(0)
	closest := int(MaxUint >> 1)
	cs := map[int]int{}
	result := 0
	for i, n1 := range nums {
		for j := i + 1; j < length -1; j++ {
			for k := j + 1; k < length; k++ {
				sum := n1 + nums[j] + nums[k]
				var t int
				if target >= sum {
					t = target - sum
				} else {
					t = sum -target
				}
				cs[t] = sum
				if t <= closest {
					closest = t
					result = cs[t]
				}
			}
		}
	}
	return result
}