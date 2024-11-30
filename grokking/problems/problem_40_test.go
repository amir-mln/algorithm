package problems_test

import (
	"encoding/json"
	"os"
	"slices"
	"testing"

	"github.com/amir-mln/algorithm/grokking/problems"
)

func TestProblem40(t *testing.T) {
	var cases []struct {
		Name  string
		Input struct {
			DNA string
			K   int
		}
		Output []string
	}

	f, err := os.Open("problem_40.json")
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
			res := problems.Problem40(c.Input).Solve()
			slices.Sort(res)
			slices.Sort(c.Output)
			if !slices.Equal(res, c.Output) {
				t.Errorf("\nfor %q and input %v, expected: %v but got: %v",
					c.Name,
					c.Input,
					c.Output,
					res,
				)
			}
		})

	}
}
