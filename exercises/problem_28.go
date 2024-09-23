package exercises

import (
	"github.com/amir-mln/algorithm/datastructure"
)

/*
Spiral Matrix
#Matrices

Given an n×n matrix, return an array containing the matrix
elements in spiral order, starting from the top-left cell.

	a) len(matrix) == len(mtx[i])
	b) 1 ≤ len(matrix) ≤ 10
	c) -100 ≤ mtx[i][j] ≤ 100
*/
var Problem28 = map[string]func(mtx datastructure.Matrix[int]) []int{
	"linear-time": func(mtx datastructure.Matrix[int]) []int {
		slc := make([]int, 0)
		top, left, right, bottom := 0, 0, len(mtx[0]), len(mtx)

		for top < bottom && left < right {
			/*
				top row
				TODO: add documentation like the last loop
			*/
			for i := left; i < right; i++ { // top row
				slc = append(slc, mtx[top][i])
			}
			/*
				right column
				TODO: add documentation like the last loop
			*/
			for i := top + 1; i < bottom && bottom-top > 1; i++ {
				slc = append(slc, mtx[i][right-1])
			}
			/*
				bottom row
				TODO: add documentation like the last loop

			*/
			for i := right - 2; i >= left && bottom-top > 1; i-- {
				slc = append(slc, mtx[bottom-1][i])
			}
			/*
				left column

				"bottom - 2" is due to the fact that bottom and right values are not zero based
				like top and left values are. Also, we have already appended the first element of the
				left column as we were appending the bottom row values.
				"top + 1" is because of a similar reason; we appended the last element of left column (
				or fir element of top row) in the first loop.
				"bottom-top > 1" is because of the edge case where we only have the top row
				and no right or left columns or bottom rows

			*/
			for i := bottom - 2; i >= top+1 && bottom-top > 1; i-- { // left column
				slc = append(slc, mtx[i][left])
			}

			top, left = top+1, left+1
			bottom, right = bottom-1, right-1
		}

		return slc
	},
}
