package exercises_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/amir-mln/algorithm/datastructure"
	"github.com/amir-mln/algorithm/exercises"
)

func TestProblem4(t *testing.T) {
	cases := []struct {
		Name   string
		N      int
		List   []int
		Output []int
	}{}

	f, err := os.Open("problem_4.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&cases)
	if err != nil {
		panic(err)
	}

	for k, v := range exercises.Problem4 {
		t.Run(k, func(t *testing.T) {
			for _, c := range cases {
				l := &datastructure.LinkedList{}
				l.InsertFromList(c.List)

				v(l, c.N)

				got := fmt.Sprint(l)
				expected := fmt.Sprint(c.Output)

				if expected != got {
					t.Errorf("case %q, expected %v; got %v", c.Name, expected, got)
				}
			}
		})
	}
}
