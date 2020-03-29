package main

import "fmt"

func main() {
	result := reverse(123)
	fmt.Println(result)
}

var MIN int = 0x80000000
var MAX int = 0x7FFFFFFF

func reverse(x int) int {
	sum := 0
	for {
		leftDigits := x / 10
		lastDigit := x % 10
		x = leftDigits

		sum = sum * 10 + lastDigit

		if 0 == leftDigits {
			break
		}
	}

	if sum < -MIN || sum > MAX {
		fmt.Println(sum)
		sum = 0
	}

	return sum
}
