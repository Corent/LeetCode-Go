package algorithm332

import "testing"

func TestFindItinerary(t *testing.T) {
	tickets := [][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}}
	t.Log(findItinerary(tickets))
}
