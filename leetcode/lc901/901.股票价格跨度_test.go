package lc901

import (
	"reflect"
	"testing"
)

func TestStockSpanner_Next(t *testing.T) {
	prices := []int{100, 80, 60, 70, 60, 75, 85}
	// prices := []int{31, 41, 48, 59, 79}
	want := []int{1, 1, 1, 2, 1, 4, 6}
	res := make([]int, 0)
	ss := Constructor()
	for _, v := range prices {
		res = append(res, ss.Next(v))
	}
	if !reflect.DeepEqual(want, res) {
		t.Errorf("StockSpanner.Next(), got = %v, want %v", res, want)
	}
}
