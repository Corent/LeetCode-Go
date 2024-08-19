package algorithm352

import "testing"

func TestSummaryRanges(t *testing.T) {
	sr := Constructor()
	t.Log("null")
	sr.AddNum(6)
	t.Log("null")
	t.Log(sr.GetIntervals())
	sr.AddNum(6)
	t.Log("null")
	t.Log(sr.GetIntervals())
	sr.AddNum(0)
	t.Log("null")
	t.Log(sr.GetIntervals())
	sr.AddNum(4)
	t.Log("null")
	t.Log(sr.GetIntervals())
}
