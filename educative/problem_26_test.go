package educative_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/amir-mln/algorithm/datastructure"
	"github.com/amir-mln/algorithm/educative"
)

func TestProblem26(t *testing.T) {
	var cases []struct {
		Name   string
		Input  datastructure.Matrix[int]
		Output datastructure.Matrix[int]
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

	for title, solution := range educative.P26Solutions {
		for _, c := range cases {
			solution(c.Input)

			if !c.Output.Equals(c.Input) {
				t.Errorf("\n in %q solution: for %q with %v, expected %v and got %v",
					title,
					c.Name,
					c.Input,
					c.Output,
					c.Input,
				)
			}
		}
	}
}
