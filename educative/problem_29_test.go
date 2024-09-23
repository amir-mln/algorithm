package educative_test

import (
	"encoding/json"
	"os"
	"slices"
	"testing"

	"github.com/amir-mln/algorithm/datastructure"
	"github.com/amir-mln/algorithm/educative"
)

func TestProblem29(t *testing.T) {
	var cases []struct {
		Name   string
		Input  datastructure.Matrix[int]
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

	for title, solution := range educative.Problem29 {
		t.Run(title, func(t *testing.T) {
			for _, c := range cases {
				t.Run(c.Name, func(t *testing.T) {
					res := solution(c.Input)

					if !slices.Equal(c.Output, res) {
						t.Errorf("\nfor input: %v, expected: %v but got: %v",
							c.Input,
							c.Output,
							res,
						)
					}
				})

			}
		})
	}
}
