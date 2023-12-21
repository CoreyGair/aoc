package y2023

import "testing"

type Day20Test struct {
	input  string
	output int
}

func TestDay20(t *testing.T) {
	tests := []Day20Test{
		{
			input: `broadcaster -> a, b, c
%a -> b
%b -> c
%c -> inv
&inv -> a`,
			output: 32000000,
		},
		{
			input: `broadcaster -> a
%a -> inv, con
&inv -> b
%b -> con
&con -> output`,
			output: 11687500,
		},
	}

	for _, test := range tests {
		res := Day20(test.input)
		if res != test.output {
			t.Errorf("expected %d got %d", test.output, res)
		}
	}
}
