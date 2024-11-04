package algorithm717

func isOneBitCharacter(bits []int) bool {
	if len(bits) < 1 {
		return false
	}
	for idx := 0; idx < len(bits); {
		if idx == len(bits)-1 {
			return bits[idx] == 0
		}
		if bits[idx] == 0 {
			idx++
		} else {
			idx += 2
		}
	}
	return false
}
