package algorithm166

import (
	"strconv"
	"strings"
)

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}
	if denominator == 0 {
		return ""
	}
	builder := &strings.Builder{}
	isNegative := numerator < 0 && denominator > 0 || numerator > 0 && denominator < 0

	num, den := int64(absInt(numerator)), int64(absInt(denominator))
	res, rem := num/den, num%den*10
	builder.WriteString(strconv.FormatInt(res, 10))
	if rem == 0 {
		return buildAns(builder, isNegative)
	}

	m := make(map[int64]int)
	builder.WriteString(".")
	for rem != 0 {
		if beg, ok := m[rem]; ok {
			tmp := builder.String()
			part1, part2 := tmp[:beg], tmp[beg:]
			builder = &strings.Builder{}
			builder.WriteString(part1)
			builder.WriteString("(")
			builder.WriteString(part2)
			builder.WriteString(")")
			return buildAns(builder, isNegative)
		}
		m[rem] = builder.Len()
		res = rem / den
		builder.WriteString(strconv.FormatInt(res, 10))
		rem = (rem % den) * 10
	}
	return buildAns(builder, isNegative)
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func buildAns(builder *strings.Builder, isNegative bool) string {
	ans := builder.String()
	if isNegative {
		return "-" + ans
	}
	return ans
}
