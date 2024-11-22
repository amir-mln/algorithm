package problems_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/amir-mln/algorithm/grokking"
	"github.com/amir-mln/algorithm/grokking/problems"
)

func TestProblem26(t *testing.T) {
	var cases []struct {
		Name   string
		Input  grokking.Matrix[int]
		Output grokking.Matrix[int]
	}

	f, err := os.Open("problem_26.json")
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
			problems.Problem26{c.Input}.Solve()

			if !c.Output.Equals(c.Input) {
				t.Errorf("\n in %q, with %v, expected %v and got %v",
					c.Name,
					c.Input,
					c.Output,
					c.Input,
				)
			}
		})
	}

}
