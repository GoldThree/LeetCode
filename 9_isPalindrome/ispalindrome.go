package main

import (
	"fmt"
	"strconv"
)

func main() {
	result := isPalindrome(-31413)
	fmt.Println(result)
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 && x >= 0 {
		return true
	}
	s := strconv.Itoa(x)
	length := len(s)
	i := 0
	j := length - 1
	if length%2 == 0 {
		for {
			if s[i] != s[j] {
				return false
			}
			if i+1 == j {
				return true
			}
			i = i + 1
			j = j - 1
		}

	}

	for {
		if s[i] != s[j] {
			return false
		}
		if i+2 == j {
			return true
		}
		i = i + 1
		j = j - 1
	}

	return true
}
