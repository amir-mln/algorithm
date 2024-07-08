package binarysearch_test

import (
	binarysearch "educative-dsa/binary-search"
	"testing"
)

type API struct {
	firstBadVersion int
}

func (a *API) IsBad(s int) bool {
	return s >= a.firstBadVersion
}

func TestFirstBadVersion(t *testing.T) {
	cases := []struct {
		name     string
		version  int
		firstBad int
		output   [2]int
	}{
		{
			name:     "test case 1",
			version:  100,
			firstBad: 67,
			output:   [...]int{67, 7},
		},
		{
			name:     "test case 2",
			version:  13,
			firstBad: 10,
			output:   [...]int{10, 3},
		},
		{
			name:     "test case 3",
			version:  29,
			firstBad: 10,
			output:   [...]int{10, 5},
		},
		{
			name:     "test case 4",
			version:  40,
			firstBad: 28,
			output:   [...]int{28, 5},
		},
		{
			name:     "test case 5",
			version:  23,
			firstBad: 19,
			output:   [...]int{19, 5},
		},
		{
			name:     "test case 6",
			version:  8,
			firstBad: 6,
			output:   [...]int{6, 3},
		},
		{
			name:     "test case 7",
			version:  5,
			firstBad: 3,
			output:   [...]int{3, 2},
		},
	}
	api := &API{}

	for k, f := range binarysearch.FirstBadVersion {
		t.Run(k, func(t *testing.T) {
			for _, c := range cases {
				api.firstBadVersion = c.firstBad
				o := f(c.version, api.IsBad)

				if o != c.output {
					t.Errorf(
						"in %q with last version %d and first bad version %d, expected %v; got %v ",
						c.name,
						c.version,
						c.firstBad,
						c.output,
						o,
					)
				}
			}
		})

	}
}
