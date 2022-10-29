package lc295

import (
	"testing"
)

func TestMedianFinder(t *testing.T) {
	mf := Constructor()
	mf.AddNum(1)
	mf.AddNum(2)
	m := mf.FindMedian()
	if m != 1.5 {
		t.Errorf("MedianFinder.FindMedian() = %v, want %v", m, 1.5)
	}
	mf.AddNum(3)
	m = mf.FindMedian()
	if m != 2 {
		t.Errorf("MedianFinder.FindMedian() = %v, want %v", m, 2)
	}
}
