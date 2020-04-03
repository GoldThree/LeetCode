package main

import "fmt"

func main() {
	fmt.Println(myAtoi("234"))
}
func myAtoi(str string) int {
	if len(str) == 0 {
		return 0
	}
	var (
		num int
		neg bool
	)
	for i, s := range str {
		if s != ' ' {
			str = str[i:]
			break
		}
	}
	if str[0] == '+' {
		str = str[1:]
	} else if str[0] == '-' {
		str = str[1:]
		neg = true
	}
	for _, s := range str {
		if s < '0' || s > '9' {
			break
		}
		num = num*10 + int(s-'0')
		if num > 2147483648 {
			num = 2147483648
			break
		}
	}
	if neg {
		num *= -1
	} else if num > 2147483647 {
		num = 2147483647
	}
	return num
}

func myAtoiII(str string) int {

	runeString := []rune(str)
	if len(str) == 0{
		return 0
	}

	isNegative := false
	result := 0
	index := -1
	for _, value := range runeString {
		if string(value) == " " {
			if index == -1 {
				continue
			}
			break
		}

		index+=1
		if index == 0 {
			if string(value) == "+" {
				continue
			}

			if string(value) == "-" {
				isNegative = true
				continue
			}

			if isNum(string(value)) == -1 {
				return 0
			}
		}

		if index != 0 && isNum(string(value)) == -1 {
			break
		}
		result = result*10+isNum(string(value))
		if result > 2147483649 {
			break
		}
	}

	if result > 2147483647 && !isNegative {
		result = 2147483647
	}

	if result > 2147483648 && isNegative {
		result = 2147483648
	}

	if isNegative {
		result = 0-result
	}
	return result

}

func isNum(num string) int {
	switch num {
	case "0":
		return 0
	case "1":
		return 1
	case "2":
		return 2
	case "3":
		return 3
	case "4":
		return 4
	case "5":
		return 5
	case "6":
		return 6
	case "7":
		return 7
	case "8":
		return 8
	case "9":
		return 9
	default:
		return -1
	}
}
