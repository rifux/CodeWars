package kata

import "testing"

func TestComp(t *testing.T) {
	var tests = []struct {
		input [2][]int
		want  bool
	}{
		{
			[2][]int{
				{121, 144, 19, 161, 19, 144, 19, 11},
				{11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19},
			},
			true,
		},
		{
			[2][]int{
				{121, 144, 19, 161, 19, 144, 19, 11},
				{11 * 21, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19},
			},
			false,
		},
		{
			[2][]int{
				nil,
				{11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19},
			},
			false,
		},
	}

	for _, test := range tests {
		ans := Comp(test.input[0], test.input[1])
		if ans != test.want {
			t.Fatalf("%v:\nwant: %t, got %t\n", test.input, test.want, ans)
		}
	}
}
