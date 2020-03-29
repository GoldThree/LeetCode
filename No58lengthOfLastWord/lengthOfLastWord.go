package main

import "fmt"

func main() {
	fmt.Println(lengthOfLastWord("   "))
}

func lengthOfLastWord(s string) int {
	runeS := []rune(s)
	if len(runeS) == 1{
		if string(runeS[0]) == " " {
			return 0
		}
		return 1
	}
	length := len(runeS)
	isFirstMeet := false

	end := length -1
	for i := length-1; i >= 0; i-- {
		if string(runeS[i]) != " " && !isFirstMeet {
			isFirstMeet = true
			end = i
			continue
		}
		if string(runeS[i]) == " " && isFirstMeet {

			return end - i
		}
		continue
	}
	if !isFirstMeet {
		return 0
	}
	return end+1
}
