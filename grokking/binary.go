package grokking

import (
	"cmp"
)

func BinarySearch[T cmp.Ordered](s []T, target T) (idx int, found bool) {
	return BinarySearchIdx(s, target, 0, len(s)-1)
}

func BinarySearchIdx[T cmp.Ordered](s []T, target T, l, r int) (idx int, found bool) {
	for l != r {
		mid := l + (r-l)/2

		if s[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return l, s[l] == target
}
