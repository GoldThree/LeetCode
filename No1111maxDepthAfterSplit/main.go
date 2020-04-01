package No1111maxDepthAfterSplit

// func maxDepthAfterSplit(seq string) []int {

//     length := len(seq)
//     if length <= 1 {
//         return []int{}
//     }

//     leftSeq := make([]string, 0)

//     result := make([]int, length)

//     leftIndexes := make([]int, 0)
//     runeSeq := []rune(seq)
//     for index, value := range runeSeq {
//         if string(value) == "(" {
//             leftSeq = append(leftSeq, "(")
//             leftIndexes = append(leftIndexes, index)
//             continue
//         }
//         if string(value) == ")" {
//             fmt.Println("-----", leftSeq)
//             deep := len(leftSeq)-1
//             result[index] = deep
//             leftIndex := leftIndexes[len(leftIndexes)-1]
//             result[leftIndex] = deep

//             leftSeq = leftSeq[:len(leftSeq)-1]
//             leftIndexes = leftIndexes[:len(leftIndexes)-1]
//         }
//     }
//     // fmt.Println(result)
//     return result
// }

func maxDepthAfterSplit(seq string) []int {

	length := len(seq)
	if length <= 1 {
		return []int{}
	}

	// leftSeq := make([]string, 0)

	result := make([]int, length)
	runeSeq := []rune(seq)
	for index, value := range runeSeq {
		if string(value) == "(" {
			result[index] = index & 1
		}
		if string(value) == ")" {
			result[index] = (index + 1) & 1
		}
	}

	return result
}
