package common

import (
	"cmp"
)

func BinarySearch[T cmp.Ordered](s []T, target T) (idx int, found bool) {
	l, r := 0, len(s)-1

	for l != r {
		mid := (r - l) / 2

		if s[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return l, s[l] == target
}
