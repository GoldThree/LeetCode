package No151reversWords

import "strings"

func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	arr := strings.Split(s, " ")
	n := len(arr)

	words := make([]string, 0, n)
	for i := n - 1; i >= 0; i-- {
		if len(arr[i]) != 0 {
			words = append(words, arr[i])
		}
	}

	return strings.Join(words, " ")
}

func reverseWords1(s string) string {

	length := len(s)
	if length == 0 {
		return ""
	}

	result := ""
	runeStr := []rune(s)

	isWordEnd := false

	temp := ""

	for i:=0;i<length;i++{
		if string(runeStr[i]) == " " {
			if isWordEnd {
				result = " "+temp+result
				temp = ""
				isWordEnd = false
			}
			continue
		}
		temp = temp+string(runeStr[i])

		if i == length-1 {
			result = temp+result
		}
		isWordEnd = true
	}

	if len(result) != 0 && string(result[0]) == " " {
		result = result[1:]
	}

	// fmt.Println(result)

	return result
}

