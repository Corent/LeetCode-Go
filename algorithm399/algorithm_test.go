package algorithm399

import "testing"

func TestCalcEquation(t *testing.T) {
	equations := [][]string{{"a", "c"}, {"b", "e"}, {"c", "d"}, {"e", "d"}}
	values := []float64{2.0, 3.0, 0.5, 5.0}
	queries := [][]string{{"a", "b"}}
	t.Log(calcEquation(equations, values, queries))
}
