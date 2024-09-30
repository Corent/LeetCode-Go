package algorithm478

import (
	"math"
	"math/rand/v2"
)

type Solution struct {
	x, y, r float64
}

func Constructor(radius float64, x_center float64, y_center float64) Solution {
	return Solution{x: x_center, y: y_center, r: radius}
}

func (this *Solution) RandPoint() []float64 {
	ca, cr := rand.Float64()*360, math.Sqrt(rand.Float64())*this.r
	return []float64{this.x + cr*math.Cos(ca), this.y + cr*math.Sin(ca)}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(radius, x_center, y_center);
 * param_1 := obj.RandPoint();
 */
