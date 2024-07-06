package algorithm030

import "testing"

func TestFindSubstring(t *testing.T) {
	ans := findSubstring("barfoothefoobarman", []string{"foo", "bar"})
	t.Log(ans)
}
