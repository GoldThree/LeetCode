package main

import "fmt"

func main() {

	test := []int{9, 9, 9}
	fmt.Println(plusOne(test))

}


func plusOne(digits []int) []int {
	length := len(digits)

	if digits[length-1] != 9 {
		result := make([]int, 0)
		result = append(result, digits[:length-1]...)
		result = append(result, digits[length-1]+1)
		return result
	}
	result := make([]int, length+1)
	isPlus := true
	for i := length-1; i >= 0; i-- {
		if !isPlus {
			copy(result[1:i+2], digits[:i+1])
			break
		}
		digits[i] += 1
		if digits[i] == 10 {
			result[i+1] = 0
			continue
		}
		result[i+1] = digits[i]
		isPlus = false
	}
	if isPlus {
		result[0] = 1
	}
	if result[0] == 0 {
		return result[1:]
	}
	return result
}