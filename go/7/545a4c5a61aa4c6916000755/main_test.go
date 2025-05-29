package main

import (
	"testing"
)

func TestGimme(t *testing.T) {
	var tests = []struct {
		input [3]int
		want  int
	}{
		{[3]int{1, 2, 3}, 1},
		{[3]int{2, 1, 3}, 0},
		{[3]int{1, 3, 2}, 2},
		{[3]int{2, 3, 1}, 0},
	}
	for _, TT := range tests {
		ans := Gimme(TT.input)
		if ans != TT.want {
			t.Errorf("%v should return %d, not %d", TT.input, TT.want, ans)
		}
	}
}
