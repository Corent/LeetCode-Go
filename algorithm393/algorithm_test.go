package algorithm393

import "testing"

func TestValidUtf8(t *testing.T) {
	t.Log(validUtf8([]int{197, 130, 1}))
}
