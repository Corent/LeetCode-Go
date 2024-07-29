package algorithm068

import "strings"

func fullJustify(words []string, maxWidth int) []string {
	ans := make([]string, 0)
	if len(words) == 0 {
		return ans
	}
	sum, line := 0, make([]string, 0)
	for i := 0; i < len(words); {
		if len(line) == 0 {
			sum, line = len(words[i]), append(line, words[i])
			i++
			continue
		}
		if sum+len(words[i])+1 <= maxWidth {
			sum, line = sum+len(words[i])+1, append(line, words[i])
			i++
			continue
		}
		ans = append(ans, handleLine(line, sum, maxWidth, i == len(words)))
		sum, line = 0, make([]string, 0)
	}
	if len(line) > 0 {
		ans = append(ans, handleLine(line, sum, maxWidth, true))
	}
	return ans
}

func handleLine(line []string, sum, maxWidth int, isLastLine bool) string {
	speCnt, wCnt := maxWidth-sum, len(line)
	if wCnt < 2 {
		builder := strings.Builder{}
		if wCnt != 0 {
			builder.WriteString(line[0])
		}
		for builder.Len() < maxWidth {
			builder.WriteByte(' ')
		}
		return builder.String()
	}
	last := line[wCnt-1]
	line = line[:wCnt-1]
	wCnt--
	m, n := 0, 0
	if !isLastLine {
		m, n = speCnt/wCnt, speCnt%wCnt
	}
	builder := strings.Builder{}
	for i := 0; i < m; i++ {
		builder.WriteByte(' ')
	}
	spe := builder.String()
	builder = strings.Builder{}
	for i := 0; i < wCnt; i++ {
		builder.WriteString(line[i])
		if len(line[i]) > 0 {
			builder.WriteByte(' ')
		}
		builder.WriteString(spe)
		if n > 0 {
			builder.WriteByte(' ')
			n--
		}
	}
	builder.WriteString(last)
	for isLastLine && builder.Len() < maxWidth {
		builder.WriteByte(' ')
	}
	return builder.String()
}
