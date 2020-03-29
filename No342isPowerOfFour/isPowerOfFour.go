package main

import "fmt"

func main() {
	fmt.Println(isPowerOfFour(16))
}

func isPowerOfFour(num int) bool {
	if num < 1 {
		return false
	}
	if num == 1 {
		return true
	}
	numFloat := float64(num)
	resultFloat := numFloat/float64(4)
	if resultFloat > float64(int(num/4)) {
		return false
	}
	return isPowerOfFour(num/4)
}
