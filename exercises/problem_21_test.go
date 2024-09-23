package exercises_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/amir-mln/algorithm/exercises"
)

func TestProblem21(t *testing.T) {
	var cases []struct {
		Name     string
		Tokens   []string
		Expected int
	}

	f, err := os.Open("problem_21.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&cases)
	if err != nil {
		panic(err)
	}

	for title, solution := range exercises.Problem21 {
		for _, c := range cases {
			i := len(c.Tokens) - 1
			res := solution(c.Tokens, &i)

			if res != c.Expected {
				t.Errorf("in %q for %q with tokens %v, expected %d and got %d",
					title,
					c.Name,
					c.Tokens,
					c.Expected,
					res,
				)
			}
		}
	}
}
