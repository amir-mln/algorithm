package educative_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/amir-mln/algorithm/educative"
)

func TestProblem1(t *testing.T) {
	cases := []struct {
		Input1 string
		Input2 string
		Result string
	}{}

	f, err := os.Open("problem_1.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&cases)
	if err != nil {
		panic(err)
	}

	for k, v := range educative.Problem1 {
		t.Run(k, func(t *testing.T) {
			for _, c := range cases {

				res := v(c.Input1, c.Input2)
				if res != c.Result {
					t.Errorf("for inputs %s and %s; expected %s got %s", c.Input1, c.Input2, c.Result, res)
				}
			}
		})
	}
}
