package stablepart

import "sort"

// StablePartition partitions the data d[first:last] according to the predicate
// It returns the index i such that the elements d[:i] satisfy the predicate
// and d[i:] do not.
func StablePartition(d sort.Interface, first, last int, pred func(i int) bool) int {

	n := last - first
	if n == 0 {
		return first
	}

	if n == 1 {
		r := first
		if pred(first) {
			r++
		}
		return r
	}

	mid := first + n/2

	return Rotate(d, StablePartition(d, first, mid, pred), mid, StablePartition(d, mid, last, pred))
}

// Rotate rotates d[first:last] by k positions so that d[k] is now at d[first].
// It returns the new index of d[first].
func Rotate(d sort.Interface, first, k, last int) int {
	Reverse(d, first, k)
	Reverse(d, k, last)
	Reverse(d, first, last)
	return first + last - k
}

// Reverse reverses d[first:last]
func Reverse(d sort.Interface, first, last int) {
	lend := (last - first)
	for i, j := 0, lend-1; i < lend/2; i, j = i+1, j-1 {
		d.Swap(first+i, first+j)
	}
}

// Slide shifts d[first:last] to target position pos and returns the new
// indices of the sublist.  If pos < first, the list slides so that the sublist
// is at d[pos:X]; if p > last, the resulting list is at d[X:pos].
func Slide(d sort.Interface, first, last, pos int) (int, int) {
	if pos < first {
		return pos, Rotate(d, pos, first, last)
	}
	if last < pos {
		return Rotate(d, first, last, pos), pos
	}
	return first, last
}

// Gather groups together elements satisfying pred around the target position pos.
func Gather(d sort.Interface, first, last, pos int, pred func(i int) bool) (int, int) {
	notp := func(i int) bool { return !pred(i) }
	return StablePartition(d, first, pos, notp), StablePartition(d, pos, last, pred)
}
