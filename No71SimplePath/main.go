package No71SimplePath

import "strings"

func simplifyPath(path string) string {
	buf := strings.Split(path, "/")
	var stack []string

	for i := 0; i < len(buf); i++ {
		if buf[i] == "" || buf[i] == "." {
			continue
		}
		if buf[i] == ".." {
			if len(stack) > 0 {
				stack = stack[0 : len(stack)-1]
			}
		} else {
			stack = append(stack, buf[i])
		}
	}

	return "/" + strings.Join(stack, "/")
}
