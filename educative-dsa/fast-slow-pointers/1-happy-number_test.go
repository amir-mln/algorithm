package fastslowpointers_test

import (
	fastslowpointers "educative-dsa/fast-slow-pointers"
	"testing"
)

func TestTrappingRainWater(t *testing.T) {
	cases := []struct {
		name   string
		number uint32
		output bool
	}{
		{
			name:   "test case 1",
			number: 2147483646,
			output: false,
		},
		{
			name:   "test case 2",
			number: 1,
			output: true,
		},
		{
			name:   "test case 3",
			number: 19,
			output: true,
		},
		{
			name:   "test case 4",
			number: 8,
			output: false,
		},
		{
			name:   "test case 5",
			number: 7,
			output: true,
		},
		{
			name:   "test case 6",
			number: 10,
			output: true,
		},
		{
			name:   "test case 7",
			number: 11,
			output: false,
		},
	}

	for k, s := range fastslowpointers.HappyNumber {
		t.Run(k, func(t *testing.T) {
			for _, c := range cases {
				res := s(c.number)

				if res != c.output {
					t.Errorf(
						"in %q with number: %d, expected %v; got %v ",
						c.name,
						c.number,
						c.output,
						res,
					)
				}
			}
		})
	}
}
