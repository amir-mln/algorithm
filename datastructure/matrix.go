package datastructure

import (
	"fmt"
	"slices"
)

type Matrix[T comparable] [][]T

func NewMatrix[T comparable](rows, cols uint) *Matrix[T] {
	mtx := Matrix[T](make([][]T, rows))

	for i := 0; i < int(rows); i++ {
		mtx[i] = make([]T, cols)
	}

	return &mtx
}

func (mtx Matrix[T]) Equals(target Matrix[T]) bool {
	if len(mtx[0]) != len(target[0]) || len(mtx) != len(target) {
		return false
	}

	for i := 0; i < len(mtx); i++ {
		if !slices.Equal(mtx[i], target[i]) {
			return false
		}
	}

	return true
}

func (mtx Matrix[T]) String() string {
	s := "\n[\n"

	for _, slc := range mtx {
		s += fmt.Sprintf("\t%v,\n", slc)
	}

	s += "]"
	return s
}
