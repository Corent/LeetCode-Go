package algorithm468

import "testing"

func TestValidIPAddress(t *testing.T) {
	t.Log(validIPAddress("172.16.254.1"))
	t.Log(validIPAddress("2001:0db8:85a3:0:0:8A2E:0370:7334"))
}
