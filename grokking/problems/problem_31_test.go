package problems_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/amir-mln/algorithm/grokking/problems"
)

func TestProblem31(t *testing.T) {
	var cases []struct {
		Name   string
		Inputs []string
		Output bool
	}

	f, err := os.Open("problem_31.json")
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
			res := problems.Problem31{c.Inputs[0], c.Inputs[1]}.Solve()

			if c.Output != res {
				t.Errorf("\nfor inputs: %q and %q, expected: %t but got: %t",
					c.Inputs[0],
					c.Inputs[1],
					c.Output,
					res,
				)
			}
		})

	}
}
