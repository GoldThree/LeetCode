package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "jbpnbwwd"
	number := LongSubString(s)
	fmt.Println(number)

}

func LongSubString(s string) int {
	max := 0
	for i := 0; i < len(s); i++ {
		for j := i+1; j <= len(s); j++ {
			substr := s[i:j-1]
			if !strings.Contains(substr, s[j-1:j]) {
				if max < j-i {
					max = j - i
				}
			} else {
				break
			}
		}
	}
	return max
}
