package main

import "testing"

func TestGetCount(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  int
	}{
		{"aaooiiuu should be 8", "aaooiiuu", 8},
		{"afgfgaoofgfgiifguu should be 8", "afgfgaoofgfgiifguu", 8},
		{"aaodfdfoi;p;pu should be 6", "aaodfdfoi;p;pu", 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := GetCount(tt.input)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
