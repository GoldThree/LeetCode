package main

import "fmt"

func main() {
	result := generateParenthesis(3)
	fmt.Println(result)
}

func generateParenthesis(n int) []string {
	res := make([]string, 0, n*n)
	bytes := make([]byte, n*2)
	dfs(n, n, 0, bytes, &res)
	return res
}

func dfs(left, right, idx int, bytes []byte, res *[]string) {
	// 所有符号都添加完毕
	if left == 0 && right == 0 {
		*res = append(*res, string(bytes))
		return
	}

	// "(" 不用担心匹配问题，
	// 只要 left > 0 就可以直接添加
	if left > 0 {
		bytes[idx] = '('
		dfs(left-1, right, idx+1, bytes, res)
	}

	// 想要添加 ")" 时
	// 需要 left < right，
	// 即在 bytes[:idx] 至少有一个 "(" 可以与 这个 ")" 匹配
	if right > 0 && left < right {
		bytes[idx] = ')'
		dfs(left, right-1, idx+1, bytes, res)
	}
}



// https://www.cnblogs.com/cchenyang/p/11418313.html
func generateParenthesis1(n int) []string {
	result := make([]string, 0)
	gen("", 0, 0, n, &result)
	return result
}


func gen(current string, left, right, n int, result *[]string) {
	if left == n && right == n {
		*result = append(*result, current)
		return
	}
	if left < n {
		gen(current+"(", left+1, right, n, result)
	}
	if right < n && left > right {
		gen(current+")", left, right+1, n, result)
	}
}

