package main

import (
	"reflect"
	"testing"
)

func TestFourSum(t *testing.T) {
	cases := []struct {
		in     []int
		target int
		want   [][]int
	}{
		{[]int{1, 0, -1, 0, -2, 2}, 0, [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}},
		{[]int{2, 2, 2, 2, 2}, 8, [][]int{{2, 2, 2, 2}}},
		{[]int{1, 2, 3}, 6, [][]int{}},
		{[]int{}, 0, [][]int{}},
		{[]int{1000000000, 1000000000, 1000000000, 1000000000}, 4000000000 - 4, [][]int{}},
	}
	for _, c := range cases {
		if got := FourSum(c.in, c.target); !reflect.DeepEqual(got, c.want) {
			t.Errorf("FourSum(%v,%d) = %v, want %v", c.in, c.target, got, c.want)
		}
	}
}
