package main

import "fmt"

func main() {
	a := "1010"
	b := "1011"
	fmt.Println(addBinary(a, b))
}

func addBinary(a string, b string) string {
	reverseA := reverseString(a)
	reverseB := reverseString(b)
	runeA := []rune(reverseA)
	runeB := []rune(reverseB)
	c := reverseA
	if len(reverseA) <= len(reverseB) {
		c  = reverseB
	}

	result := make([]rune, len([]rune(c)))
	isAdd := false
	for i, v := range c {
		if i > len(runeA)-1 || i > len(runeB)-1 {
			if isAdd {
				if string(v) == "1" {
					result[i] = []rune("0")[0]
					continue
				}
				result[i] = []rune("1")[0]
				isAdd = false
				continue
			}
			result[i] = v
			isAdd = false
			continue
		}
		if string(runeA[i]) == string(runeB[i]){
			if string(runeA[i]) == "1" {
				if isAdd {
					result[i] = []rune("1")[0]
					isAdd = true
					continue
				}
				result[i] = []rune("0")[0]
				isAdd = true
				continue
			}
			if isAdd {
				result[i] = []rune("1")[0]
				isAdd = false
				continue
			}
			result[i] = []rune("0")[0]
			isAdd = false
			continue
		}
		if isAdd {
			result[i] = []rune("0")[0]
			continue
		}
		isAdd = false
		result[i] = []rune("1")[0]
	}
	r := reverseString(string(result))
	if isAdd {
		r = fmt.Sprintf("%s%s", "1", r)
	}
	return r
}


func reverseString(s string) string {
	runes := []rune(s)

	for from, to := 0, len(runes)-1; from < to; from, to = from + 1, to - 1 {
		runes[from], runes[to] = runes[to], runes[from]
	}

	return string(runes)
}
