package algorithm1233

import (
	"sort"
	"strings"
)

func removeSubfolders(folder []string) []string {
	sort.Strings(folder)
	rst := make([]string, 0)
	for i := 0; i < len(folder); {
		j := i + 1
		for j < len(folder) && isParentFolder(folder[i], folder[j]) {
			j++
		}
		rst = append(rst, folder[i])
		i = j
	}
	return rst
}

func isParentFolder(folder1, folder2 string) bool {
	if folder1 == folder2 {
		return true
	}
	len1, len2 := len(folder1), len(folder2)
	return strings.HasPrefix(folder2, folder1) && len2 > len1 && folder2[len1] == '/'
}
