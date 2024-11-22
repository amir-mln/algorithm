package problems_test

import (
	"encoding/json"
	"os"
	"slices"
	"testing"

	"github.com/amir-mln/algorithm/grokking"
	"github.com/amir-mln/algorithm/grokking/problems"
)

func TestProblem29(t *testing.T) {
	var cases []struct {
		Name   string
		Input  grokking.Matrix[int]
		Output []int
	}

	f, err := os.Open("problem_29.json")
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
			res := problems.Problem29{c.Input}.Solve()

			if !slices.Equal(c.Output, res) {
				t.Errorf("\nfor input: %v, expected: %v but got: %v",
					c.Input,
					c.Output,
					res,
				)
			}
		})

	}
}
