package problems_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/amir-mln/algorithm/grokking/problems"
)

func TestProblem33(t *testing.T) {
	var cases []struct {
		Name   string
		Inputs []int
		Output string
	}

	f, err := os.Open("problem_33.json")
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
			problem := problems.Problem33{c.Inputs[0], c.Inputs[1]}
			res := problem.Solve()
			if c.Output != res {
				t.Errorf("\nfor case %q input %v, expected: %q but got: %q",
					c.Name,
					c.Inputs,
					c.Output,
					res,
				)
			}
		})

	}
}
