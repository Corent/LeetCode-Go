package algorithm071

import "strings"

func simplifyPath(path string) string {
	if len(path) == 0 {
		return path
	}
	for strings.HasPrefix(path, "/") {
		path = path[1:]
	}
	for strings.HasSuffix(path, "/") {
		path = path[:len(path)-1]
	}
	ss, stack := strings.Split(path, "/"), make([]string, 0)
	for _, s := range ss {
		switch s {
		case "", ".":
			// do nothing
		case "..":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, s)
		}
	}
	if len(stack) == 0 {
		return "/"
	}
	return "/" + strings.Join(stack, "/")
}
