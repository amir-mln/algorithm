package problems_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/amir-mln/algorithm/grokking/problems"
)

func TestProblem22(t *testing.T) {
	var cases []struct {
		Name   string
		Input  string
		Output string
	}

	f, err := os.Open("problem_22.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&cases)
	if err != nil {
		panic(err)
	}

	for title, solution := range problems.Problem22 {
		for _, c := range cases {
			res := solution(c.Input)

			if res != c.Output {
				t.Errorf("in %q for %q with %v, expected %v and got %v",
					title,
					c.Name,
					c.Input,
					c.Output,
					res,
				)
			}
		}
	}
}
