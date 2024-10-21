package algorithm564

import (
	"math"
	"strconv"
)

func nearestPalindromic(n string) string {
	length, candidates := len(n), make(map[int64]struct{})
	candidates[int64(math.Pow10(length))+1] = struct{}{}
	candidates[int64(math.Pow10(length-1))-1] = struct{}{}
	prefix, _ := strconv.ParseInt(n[:(length+1)>>1], 10, 64)
	for i := int64(-1); i <= 1; i++ {
		p := strconv.FormatInt(prefix+i, 10)
		if length&1 == 0 {
			p += reverse(p)
		} else {
			p += reverse(p[:len(p)-1])
		}
		pp, _ := strconv.ParseInt(p, 10, 64)
		candidates[pp] = struct{}{}
	}
	num, _ := strconv.ParseInt(n, 10, 64)
	delete(candidates, num)
	minDiff, rst := int64(math.MaxInt64), int64(math.MaxInt64)
	for c := range candidates {
		if diff := absInt64(c - num); diff < minDiff {
			minDiff, rst = diff, c
		} else if diff == minDiff {
			rst = minInt64(rst, c)
		}
	}
	return strconv.FormatInt(rst, 10)
}

func reverse(s string) string {
	if len(s) < 2 {
		return s
	}
	bytes := []byte(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	return string(bytes)
}

func absInt64(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
}

func minInt64(m, n int64) int64 {
	if m < n {
		return m
	}
	return n
}
