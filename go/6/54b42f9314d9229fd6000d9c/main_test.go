package main

import "testing"

func TestDuplicateEncode(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"din", "((("},
		{"recede", "()()()"},
		{"Success", ")())())"},
		{"(( @", "))(("},
	}

	for _, TT := range tests {
		ans := DuplicateEncode(TT.input)
		if ans != TT.want {
			t.Fatalf("%s: want %s, not %s", TT.input, TT.want, ans)
		}
	}
}
