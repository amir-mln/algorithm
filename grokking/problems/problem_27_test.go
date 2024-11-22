package problems_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/amir-mln/algorithm/grokking"
	"github.com/amir-mln/algorithm/grokking/problems"
)

func TestProblem27(t *testing.T) {
	var cases []struct {
		Name   string
		Input  grokking.Matrix[int]
		Output grokking.Matrix[int]
	}

	f, err := os.Open("problem_27.json")
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
			cp := c.Input.Copy()
			problems.Problem27{c.Input}.Solve()

			if !c.Output.Equals(c.Input) {
				t.Errorf("\n for input: %v, expected %v and got %v",
					cp,
					c.Output,
					c.Input,
				)
			}
		})

	}
}
