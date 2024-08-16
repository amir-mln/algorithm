package educative_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/amir-mln/algorithm/educative"
)

func TestProblem25(t *testing.T) {
	var cases []struct {
		Name   string
		Input  []int
		Output int
	}

	f, err := os.Open("problem_25.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&cases)
	if err != nil {
		panic(err)
	}

	for title, solution := range educative.P25Solutions {
		for _, c := range cases {
			res := solution(c.Input)

			if res != c.Output {
				t.Errorf("\n in %q solution: for %q with %v, expected %v and got %v",
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
