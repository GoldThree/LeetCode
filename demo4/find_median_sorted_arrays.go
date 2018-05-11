package main

import (
	"sort"
	"fmt"
)

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	result := FindMedianSortedArrays(nums1, nums2)
	fmt.Println(result)
}

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	for _, v := range nums2 {
		nums1 = append(nums1, v)
	}
	sort.Ints(nums1)
	length1 := len(nums1)
	if length1%2 == 0 {
		index1 := (length1 - 1) / 2
		index2 := index1 + 1
		result := (float64(nums1[index1]) + float64(nums1[index2])) / 2
		return result
	}
	index3 := (length1 - 1) / 2
	return float64(nums1[index3])

}
