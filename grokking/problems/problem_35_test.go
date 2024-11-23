package problems_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/amir-mln/algorithm/grokking/problems"
)

func TestProblem35(t *testing.T) {
	var cases []struct {
		Name   string
		Input  string
		Output bool
	}

	f, err := os.Open("problem_35.json")
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
			res := problems.Problem35(c.Input).Solve()
			if c.Output != res {
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
