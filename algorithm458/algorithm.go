package algorithm458

import (
	"math"
)

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	base := minutesToTest/minutesToDie + 1
	tmp := math.Log(float64(buckets)) / math.Log(float64(base))
	return int(math.Ceil(float64(float32(tmp))))
}
