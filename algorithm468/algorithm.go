package algorithm468

import (
	"regexp"
)

func validIPAddress(queryIP string) string {
	re := regexp.MustCompile(`^((25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9][0-9]|[0-9])\.){3}(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9][0-9]|[0-9])$`)
	matches := re.FindStringSubmatch(queryIP)
	if len(matches) > 0 && matches[0] == queryIP {
		return "IPv4"
	}
	re = regexp.MustCompile(`^([0-9a-fA-F]{1,4}:){7}([0-9a-fA-F]{1,4})$`)
	matches = re.FindStringSubmatch(queryIP)
	if len(matches) > 0 && matches[0] == queryIP {
		return "IPv6"
	}
	return "Neither"
}
