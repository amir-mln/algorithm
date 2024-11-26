package problems_test

import (
	"encoding/json"
	"os"
	"slices"
	"testing"

	"github.com/amir-mln/algorithm/grokking/problems"
)

func TestProblem37(t *testing.T) {
	var cases []struct {
		Name   string
		Input  []string
		Output []int
	}

	f, err := os.Open("problem_37.json")
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
			res := problems.Problem37(c.Input).Solve()
			if !slices.Equal(c.Output, res) {
				t.Errorf("\nfor case %q and input %q, expected: %v but got: %v",
					c.Name,
					c.Input,
					c.Output,
					res,
				)
			}
		})

	}
}
