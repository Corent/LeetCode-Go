package algorithm093

import (
	"strconv"
	"strings"
)

var (
	ip, ans []string
)

func restoreIpAddresses(s string) []string {
	ip, ans = []string{}, []string{}
	restoreIpAddressesCore(s, 0)
	return ans
}

func restoreIpAddressesCore(s string, idx int) {
	n, cnt := len(s), 4-len(ip)
	if n == 0 || n-idx < cnt || n-idx > cnt*3 || !(cnt >= 0 && cnt <= 4) {
		return
	}
	if idx == n && cnt == 0 {
		ans = append(ans, strings.Join(ip, "."))
		return
	}

	if idx < n {
		ip = append(ip, s[idx:idx+1])
		restoreIpAddressesCore(s, idx+1)
		ip = ip[:len(ip)-1]

		if idx+1 < n {
			v, _ := strconv.Atoi(s[idx : idx+2])
			if v >= 10 && v <= 99 {
				ip = append(ip, s[idx:idx+2])
				restoreIpAddressesCore(s, idx+2)
				ip = ip[:len(ip)-1]
			}

			if idx+2 < n {
				v, _ = strconv.Atoi(s[idx : idx+3])
				if v >= 100 && v <= 255 {
					ip = append(ip, s[idx:idx+3])
					restoreIpAddressesCore(s, idx+3)
					ip = ip[:len(ip)-1]
				}
			}
		}
	}
}
