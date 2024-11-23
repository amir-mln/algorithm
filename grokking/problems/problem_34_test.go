package problems_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/amir-mln/algorithm/grokking/problems"
)

func TestProblem34(t *testing.T) {
	var cases []struct {
		Name   string
		Inputs []int
		Output bool
	}

	f, err := os.Open("problem_34.json")
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
			res := problems.Problem34{c.Inputs}.Solve()
			if c.Output != res {
				t.Errorf("\nfor case %q inputs %v, expected: %v but got: %v",
					c.Name,
					c.Inputs,
					c.Output,
					res,
				)
			}
		})

	}
}
