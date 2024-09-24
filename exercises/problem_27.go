package exercises

import (
	"github.com/amir-mln/algorithm/datastructure"
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
var Problem27 = map[string]func(mtx datastructure.Matrix[int]){
	"loops": func(mtx datastructure.Matrix[int]) {
		top, left, right, bottom := 0, 0, len(mtx)-1, len(mtx)-1

		for top < bottom && left < right {
			for i := 0; i < right-left; i++ {
				temp := mtx[top+i][right]                  // store right column
				mtx[top+i][right] = mtx[top][left+i]       // swap right column with top row
				mtx[top][left+i] = mtx[bottom-i][left]     // swap top row with left column
				mtx[bottom-i][left] = mtx[bottom][right-i] // swap left column with bottom row
				mtx[bottom][right-i] = temp
			}

			top, left = top+1, left+1
			bottom, right = bottom-1, right-1
		}
	},
}
