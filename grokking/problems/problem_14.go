package problems

var FirstBadVersion = map[string]func(version int, isBadVersion func(int) bool) [2]int{
	"binary-search": func(version int, isBadVersion func(int) bool) [2]int {
		var output [2]int
		l, r := 1, version
		calls := 0

		for l != r {
			mid := l + (r-l)/2
			calls++
			if isBadVersion(mid) {
				r = mid
			} else {
				l = mid + 1
			}
		}

		output[0] = l
		output[1] = calls
		return output
	},
}
