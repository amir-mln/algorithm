package problems_test

import (
	"encoding/json"
	"os"
	"slices"
	"testing"

	"github.com/amir-mln/algorithm/grokking/problems"
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

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			res := problems.Problem30{c.Nums1, c.Nums2}.Solve()

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
}
