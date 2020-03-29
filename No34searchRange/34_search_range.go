package main

import "fmt"

func main() {
	num := []int{5,7,7,8,8,8,10}
	result := searchRange(num, 8)
	fmt.Println(result)

}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	i,j := 0, len(nums)-1
	start, end := -1, -1
	result := make([]int, 0)
	for {
		if start != -1 && end != -1 {
			result = append(result, start)
			result = append(result, end)
			return result
		}
		if i == j {
			if nums[i] == target {
				start = i
				end = i
			}
			result = append(result, start)
			result = append(result, end)
			return result
		}
		if nums[i] == target {
			start = i
			if nums[j] == target {
				end = j
				result = append(result, start)
				result = append(result, end)
				return result
			}
			j--
			continue
		}
		if nums[j] == target {
			end = j
			if nums[i] == target {
				start = i
				result = append(result, start)
				result = append(result, end)
				return result
			}
			i++
			continue
		}
		i++
		j--
		if i == len(nums) -1 || j == 0 {
			result = append(result, start)
			result = append(result, end)
			return result
		}
	}
}
