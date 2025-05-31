package kata

import "testing"

func TestNearestSq(t *testing.T) {
	var tests = []struct {
		input int
		want  int
	}{
		{1, 1},
		{2, 1},
		{10, 9},
		{111, 121},
		{9999, 10000},
	}

	for _, TT := range tests {
		ans := NearestSq(TT.input)
		if ans != TT.want {
			t.Fatalf("%d: want %d, not %d", TT.input, TT.want, ans)
		}
	}
}
