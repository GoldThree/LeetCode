package main

import "fmt"

func main()  {
	height := []int{1,8,6,2,5,4,8,3,7}
	fmt.Println(maxArea(height))
}

func maxArea(height []int) int {
	lHeight := len(height)
	maxA := 0
	for i, h1 := range height {
		for j := i + 1; j < lHeight; j++ {
			h2 := height[j]
			if h1 <= h2{
				h2 = h1
			}
			maxB := h2 * (j-i)
			if maxB >= maxA {
				maxA = maxB
			}
		}
	}
	return maxA
}