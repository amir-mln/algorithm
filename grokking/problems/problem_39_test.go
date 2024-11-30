package problems_test

import (
	"encoding/json"
	"os"
	"slices"
	"testing"

	"github.com/amir-mln/algorithm/grokking/problems"
)

func TestProblem39(t *testing.T) {
	var cases []struct {
		Name   string
		Input  []string
		Output [][]string
	}

	f, err := os.Open("problem_39.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&cases)
	if err != nil {
		panic(err)
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			groups := problems.Problem39(c.Input).Solve()

			for _, g := range groups {
				found := false
				slices.Sort(g)
				for _, slc := range c.Output {
					slices.Sort(slc)
					if slices.Equal(slc, g) {
						found = true
						break
					}
				}

				if !found {
					t.Errorf("\nfor case %q and input %q, expected: %v but got: %v",
						c.Name,
						c.Input,
						c.Output,
						groups,
					)
				}
			}
		})

	}
}
