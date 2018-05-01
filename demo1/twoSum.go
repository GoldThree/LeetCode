package main

func main() {
	nums := []int{2, 7, 0, 11}
	twoSum(nums, 9)
}

func twoSum(nums []int, target int) []int {
	length := len(nums)
	a := make([]int, 0)
	for k, v := range nums{
		for i := k + 1; i < length; i++ {
			if v + nums[i] == target {
				a = append(a, k)
				a = append(a, i)
			}
		}
	}
	return a
}
