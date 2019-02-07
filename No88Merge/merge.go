package main

import "fmt"

func main() {
	nums1 := []int{4,5,6,0,0,0}
	nums2 := []int{1,2,3}
	merge(nums1, 3, nums2, 3)
}

func merge(nums1 []int, m int, nums2 []int, n int)  {
	if n == 0 {
		return
	}
	if m == 0 {
		for i := 0; i < n; i++ {
			nums1[i]= nums2[i]
		}
		return
	}
	for index, num2 := range nums2 {
		for i := 0; i < m+n; i++ {
			// 如果是最后一个
			if i == m+n-1 {
				nums1[i] = num2
				continue
			}
			// 如果num2不在num1的i和i+1之间
			if num2 >= nums1[i+1] && i+1<m+index {
				continue
			}
			// 如果第二个数组的某一个数在第一个数组的两个数字之间或者小于第一个数组的第一个数字
			if num2 >= nums1[i] && num2 < nums1[i+1]  {
				numsResult := ProcessNums(i+1, m+index, num2, nums1)
				nums1 = numsResult
				break
			}
			if num2 < nums1[0] {
				numsResult := ProcessNums(0, m+index, num2, nums1)
				nums1 = numsResult
				break
			}
			// 如果当前数字大于数组1中的所有数字
			nums1[m+index] = num2
			break
		}
	}

	return
}

func ProcessNums(start, end, insertNum int, nums []int) []int {
	for j := end; j > start; j-- {
		nums[j] = nums[j-1]
	}
	nums[start] = insertNum
	return nums
}
