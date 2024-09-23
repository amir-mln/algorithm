package exercises_test

import (
	"encoding/json"
	"os"
	"slices"
	"testing"

	"github.com/amir-mln/algorithm/datastructure"
	"github.com/amir-mln/algorithm/exercises"
)

func TestProblem28(t *testing.T) {
	var cases []struct {
		Name   string
		Input  datastructure.Matrix[int]
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

	for title, solution := range exercises.Problem28 {
		t.Run(title, func(t *testing.T) {
			for _, c := range cases {
				t.Run(c.Name, func(t *testing.T) {
					res := solution(c.Input)

					// fmt.Sprint(res) != fmt.Sprint(c.Output)
					if !slices.Equal(c.Output, res) {
						t.Errorf("\n for input: %v, expected %v and got %v",
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
