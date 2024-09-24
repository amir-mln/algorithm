package exercises

import (
	"github.com/amir-mln/algorithm/datastructure"
)

/*
# Where Will the Ball Fall

You have `n` balls and a `2D` grid of size `m × n` representing a box.
The box is open on the top and bottom sides. Each cell in the box has
a diagonal that can redirect a ball to the right or the left. You must
drop `n` balls at each column’s top. The goal is to determine whether
each ball will fall out of the bottom or become stuck in the box. Each
cell in the grid has a value of `1` or `−1`.

- `1` represents that the grid will redirect the ball to the right.
- `−1` represents that the grid will redirect the ball to the left.

A ball gets stuck if it hits a V-shaped pattern between two grids or if
a grid redirects the ball into either wall of the box. The solution should
return an array of size `n`, with the `i`th element indicating the column
that the ball falls out of, or it becomes −1 if it’s stuck. If the ball
drops from column `x` and falls out from column `y`, then in the result
array, index `x` contains value `y`.

Constraints:

 1. 1 ≤ len(matrix) ≤ 100
 2. 1 ≤ len(mtx[i]) ≤ 100
 3. mtx[i][j] is either 1 or -1
*/
var Problem29 = map[string]func(mtx datastructure.Matrix[int]) []int{
	"linear-time": func(mtx datastructure.Matrix[int]) []int {
		res := make([]int, len(mtx[0]))

		for i := 0; i < len(mtx[0]); i++ {
			r, c := 0, i

			for {
				val := mtx[r][c]
				if val+c == -1 || val+c == len(mtx[0]) { // hit the left or right wall
					res[i] = -1
					break
				}

				if val+mtx[r][c+val] == 0 { // stuck between V-like cells
					res[i] = -1
					break
				}

				c += val
				r += 1
				if r == len(mtx) { // fell out of bottom
					res[i] = c
					break
				}
			}
		}

		return res
	},
}
