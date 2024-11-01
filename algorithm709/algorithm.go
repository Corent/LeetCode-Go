package algorithm709

func toLowerCase(s string) string {
	bytes, diff := []byte(s), byte('a'-'A')
	for i := range bytes {
		if bytes[i] >= 'A' && bytes[i] <= 'Z' {
			bytes[i] += diff
		}
	}
	return string(bytes)
}
