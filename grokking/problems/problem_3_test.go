package problems_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/amir-mln/algorithm/grokking/problems"
)

func TestProblem3(t *testing.T) {
	var cases []struct {
		Input  string
		Output bool
	}

	f, err := os.Open("problem_3.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&cases)
	if err != nil {
		panic(err)
	}

	for k, v := range problems.Problem3 {
		t.Run(k, func(t *testing.T) {
			for i, c := range cases {
				title := fmt.Sprintf("test case %d", i+1)
				t.Run(title, func(t *testing.T) {
					result := v(c.Input)

					if result != c.Output {
						t.Errorf(
							"for input %s.\n expected %t got %t",
							c.Input,
							c.Output,
							result,
						)
					}
				})
			}
		})
	}
}
