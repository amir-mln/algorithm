package problems_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/amir-mln/algorithm/grokking/problems"
)

func TestProblem44(t *testing.T) {
	var cases []struct {
		Name  string
		Input struct {
			S string
			C int
		}
		Output int
	}

	f, err := os.Open("problem_44.json")
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
			res := problems.Problem44(c.Input).Solve()

			if res != c.Output {
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
