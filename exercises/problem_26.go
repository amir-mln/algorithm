package exercises

import (
	"github.com/amir-mln/algorithm/datastructure"
)

/*
Set Matrix Zeros

Given a matrix, `mtx`, if any element within the matrix is zero, set that row and column to zero.

	a) 1 ≤ len(mtx), len(mtx[i]) ≤ 20
	b) -2**31 ≤ mtx[i][j] ≤ 2**31 - 1
*/
var Problem26 = map[string]func(mtx datastructure.Matrix[int]){
	"hash-map": func(mtx datastructure.Matrix[int]) {
		rows := make(map[int]struct{})
		cols := make(map[int]struct{})

		for i := 0; i < len(mtx); i++ {
			for j := 0; j < len((mtx)[0]); j++ {
				if (mtx)[i][j] == 0 {
					rows[i] = struct{}{}
					cols[j] = struct{}{}
					continue
				}
			}
		}

		for i := 0; i < len(mtx); i++ {
			for j := 0; j < len((mtx)[0]); j++ {
				if _, ok := rows[i]; ok {
					(mtx)[i][j] = 0
					continue
				}

				if _, ok := cols[j]; ok {
					(mtx)[i][j] = 0
					continue
				}
			}
		}
	},
}
