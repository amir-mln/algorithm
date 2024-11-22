package problems

import (
	"github.com/amir-mln/algorithm/grokking"
)

/*
# Rotate Matrix

Given an n×n matrix, rotate the matrix 90 degrees clockwise.
The performed rotation should be in place, i.e., the given matrix
is modified directly without allocating another matrix.

 1. len(matrix) == len(mtx[i])
 2. 1 ≤ len(matrix) ≤ 20
 3. -2**31 ≤ mtx[i][j] ≤ 2**31 - 1
*/
type Problem27 struct {
	Mtx grokking.Matrix[int]
}

func (p Problem27) Solve() {
	top, left, right, bottom := 0, 0, len(p.Mtx)-1, len(p.Mtx)-1

	for top < bottom && left < right {
		for i := 0; i < right-left; i++ {
			temp := p.Mtx[top+i][right]                    // store right column
			p.Mtx[top+i][right] = p.Mtx[top][left+i]       // swap right column with top row
			p.Mtx[top][left+i] = p.Mtx[bottom-i][left]     // swap top row with left column
			p.Mtx[bottom-i][left] = p.Mtx[bottom][right-i] // swap left column with bottom row
			p.Mtx[bottom][right-i] = temp
		}

		top, left = top+1, left+1
		bottom, right = bottom-1, right-1
	}
}
