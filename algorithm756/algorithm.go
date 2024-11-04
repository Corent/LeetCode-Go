package algorithm756

var (
	ans        bool
	allowedMap map[string]map[byte]struct{}
)

func pyramidTransition(bottom string, allowed []string) bool {
	if len(bottom) == 0 || len(bottom) > 1 && len(allowed) == 0 {
		return false
	}
	if len(bottom) == 1 {
		return true
	}
	ans, allowedMap = false, map[string]map[byte]struct{}{}
	for _, s := range allowed {
		key := s[:2]
		chs, _ := allowedMap[key]
		if chs == nil {
			chs = map[byte]struct{}{}
		}
		chs[s[2]] = struct{}{}
		allowedMap[key] = chs
	}
	pyramidTransitionHelper1([]byte(bottom))
	return ans
}

func pyramidTransitionHelper1(chs []byte) {
	if ans {
		return
	}
	if len(chs) == 1 {
		ans = true
		return
	}
	sets := make([]map[byte]struct{}, 0)
	for i := 1; i < len(chs); i++ {
		key := string([]byte{chs[i-1], chs[i]})
		if v, ok := allowedMap[key]; ok {
			sets = append(sets, v)
		} else {
			return
		}
	}
	pyramidTransitionHelper2(0, make([]byte, len(chs)-1), sets)
}

func pyramidTransitionHelper2(idx int, chs []byte, sets []map[byte]struct{}) {
	if idx == len(sets) {
		pyramidTransitionHelper1(chs)
		return
	}
	for ch := range sets[idx] {
		chs[idx] = ch
		pyramidTransitionHelper2(idx+1, chs, sets)
	}
}
