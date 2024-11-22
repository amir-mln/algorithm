package problems_test

import (
	"encoding/json"
	"os"
	"slices"
	"testing"

	"github.com/amir-mln/algorithm/grokking"
	"github.com/amir-mln/algorithm/grokking/problems"
)

func TestProblem28(t *testing.T) {
	var cases []struct {
		Name   string
		Input  grokking.Matrix[int]
		Output []int
	}

	f, err := os.Open("problem_28.json")
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
			res := problems.Problem28{c.Input}.Solve()

			if !slices.Equal(c.Output, res) {
				t.Errorf("\n for input: %v, expected %v and got %v",
					c.Input,
					c.Output,
					res,
				)
			}
		})

	}
}
