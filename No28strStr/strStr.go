package main

import "fmt"

func main() {
	result := strStr("asdfg","asd")
	fmt.Println(result)
}
func strStr(haystack string, needle string) int {
	hlen, nlen := len(haystack), len(needle)
	// 当hlen等于nlen的时候，需要i == 0
	for i := 0; i <= hlen-nlen; i++ {
		if haystack[i:i+nlen] == needle {
			return i
		}
	}

	return -1
}
