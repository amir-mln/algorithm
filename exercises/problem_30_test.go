package exercises_test

import (
	"encoding/json"
	"os"
	"slices"
	"testing"

	"github.com/amir-mln/algorithm/exercises"
)

func TestProblem30(t *testing.T) {
	var cases []struct {
		Name   string
		Nums1  []int
		Nums2  []int
		Output []int
	}

	f, err := os.Open("problem_30.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&cases)
	if err != nil {
		panic(err)
	}

	for title, solution := range exercises.Problem30 {
		t.Run(title, func(t *testing.T) {
			for _, c := range cases {
				t.Run(c.Name, func(t *testing.T) {
					res := solution(c.Nums1, c.Nums2)

					if !slices.Equal(c.Output, res) {
						t.Errorf("\nfor inputs: %v and %v, expected: %v but got: %v",
							c.Nums1,
							c.Nums2,
							c.Output,
							res,
						)
					}
				})

			}
		})
	}
}
