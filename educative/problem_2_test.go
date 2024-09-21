package educative_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/amir-mln/algorithm/educative"
)

func TestProblem2(t *testing.T) {
	cases := []struct {
		title  string
		input  string
		output int32
	}{}

	f, err := os.Open("problem_2.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&cases)
	if err != nil {
		panic(err)
	}

	for k, v := range educative.Problem2 {
		t.Run(k, func(t *testing.T) {
			for _, c := range cases {
				t.Run(c.title, func(t *testing.T) {
					result := v(c.input)
					if c.output != result {
						t.Errorf(
							"for input %q, expected %d got %d",
							c.input,
							c.output,
							result,
						)

					}
				})
			}
		})
	}

}
