package problems

import (
	"github.com/amir-mln/algorithm/grokking"
)

/*
# Set Matrix Zeros

Given a matrix, `mtx`, if any element within the matrix is zero, set that row and column to zero.

Constraints:
 1. 1 ≤ len(mtx), len(mtx[i]) ≤ 20
 2. -2**31 ≤ mtx[i][j] ≤ 2**31 - 1
*/
type Problem26 struct{ Mtx grokking.Matrix[int] }

func (p Problem26) Solve() {
	rows := make(map[int]struct{})
	cols := make(map[int]struct{})

	for i := 0; i < len(p.Mtx); i++ {
		for j := 0; j < len((p.Mtx)[0]); j++ {
			if (p.Mtx)[i][j] == 0 {
				rows[i] = struct{}{}
				cols[j] = struct{}{}
				continue
			}
		}
	}

	for i := 0; i < len(p.Mtx); i++ {
		for j := 0; j < len((p.Mtx)[0]); j++ {
			if _, ok := rows[i]; ok {
				(p.Mtx)[i][j] = 0
				continue
			}

			if _, ok := cols[j]; ok {
				(p.Mtx)[i][j] = 0
				continue
			}
		}
	}
}
