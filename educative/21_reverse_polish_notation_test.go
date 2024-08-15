package educative_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/amir-mln/algorithm/educative"
)

func TestReversePolishNotation(t *testing.T) {
	t.Log("Testing", "Reverse Polish Notation")
	var cases []struct {
		Name     string
		Tokens   []string
		Expected int
	}

	f, err := os.Open("21_reverse_polish_notation.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&cases)
	if err != nil {
		panic(err)
	}

	for title, solution := range educative.P21Solutions {
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
