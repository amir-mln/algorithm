package problems_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/amir-mln/algorithm/grokking/problems"
)

func TestProblem32(t *testing.T) {
	var cases []struct {
		Name   string
		Input  string
		Output int
	}

	f, err := os.Open("problem_32.json")
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
			res := problems.Problem32{c.Input}.Solve()

			if c.Output != res {
				t.Errorf("\nfor input %q, expected: %d but got: %d",
					c.Input,
					c.Output,
					res,
				)
			}
		})

	}
}
