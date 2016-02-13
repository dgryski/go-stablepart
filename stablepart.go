package stablepart

import "sort"

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

func Rotate(d sort.Interface, first, k, last int) int {
	Reverse(d, first, k)
	Reverse(d, k, last)
	Reverse(d, first, last)
	return first + last - k
}

func Reverse(d sort.Interface, first, last int) {
	lend := (last - first)
	for i, j := 0, lend-1; i < lend/2; i, j = i+1, j-1 {
		d.Swap(first+i, first+j)
	}
}

func Slide(d sort.Interface, first, last, pos int) (int, int) {
	if pos < first {
		return pos, Rotate(d, pos, first, last)
	}
	if last < pos {
		return Rotate(d, first, last, pos), pos
	}
	return first, last
}

func Gather(d sort.Interface, first, last, pos int, pred func(i int) bool) (int, int) {
	notp := func(i int) bool { return !pred(i) }
	return StablePartition(d, first, pos, notp), StablePartition(d, pos, last, pred)
}
