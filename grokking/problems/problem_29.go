package problems

import "github.com/amir-mln/algorithm/grokking"

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
type Problem29 struct{ Mtx grokking.Matrix[int] }

func (p Problem29) Solve() []int {
	res := make([]int, len(p.Mtx[0]))

	for i := 0; i < len(p.Mtx[0]); i++ {
		r, c := 0, i

		for {
			val := p.Mtx[r][c]
			if val+c == -1 || val+c == len(p.Mtx[0]) { // hit the left or right wall
				res[i] = -1
				break
			}

			if val+p.Mtx[r][c+val] == 0 { // stuck between V-like cells
				res[i] = -1
				break
			}

			c += val
			r += 1
			if r == len(p.Mtx) { // fell out of bottom
				res[i] = c
				break
			}
		}
	}

	return res
}
