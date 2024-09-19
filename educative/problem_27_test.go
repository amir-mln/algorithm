package educative_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/amir-mln/algorithm/datastructure"
	"github.com/amir-mln/algorithm/educative"
)

func TestProblem27(t *testing.T) {
	var cases []struct {
		Name   string
		Input  datastructure.Matrix[int]
		Output datastructure.Matrix[int]
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

	for title, solution := range educative.P27Solutions {
		t.Run(title, func(t *testing.T) {
			for _, c := range cases {
				t.Run(c.Name, func(t *testing.T) {
					cp := c.Input.Copy()
					solution(c.Input)

					if !c.Output.Equals(c.Input) {
						t.Errorf("\n for input: %v, expected %v and got %v",
							cp,
							c.Output,
							c.Input,
						)
					}
				})

			}
		})
	}
}
