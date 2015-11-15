package stablepart

import "sort"

func StablePartition(d sort.Interface, f, l int, p func(i int) bool) int {

	n := l - f
	if n == 0 {
		return f
	}

	if n == 1 {
		r := f
		if p(f) {
			r++
		}
		return r
	}

	m := f + n/2

	return Rotate(d, StablePartition(d, f, m, p), m, StablePartition(d, m, l, p))
}

func Rotate(d sort.Interface, f, k, l int) int {
	Reverse(d, f, k)
	Reverse(d, k, l)
	Reverse(d, f, l)
	return f + l - k
}

func Reverse(d sort.Interface, f, l int) {
	lend := (l - f)
	for i, j := 0, lend-1; i < lend/2; i, j = i+1, j-1 {
		d.Swap(f+i, f+j)
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

func Gather(d sort.Interface, first, last, pos int, p func(i int) bool) (int, int) {
	notp := func(i int) bool { return !p(i) }
	return StablePartition(d, first, pos, notp), StablePartition(d, pos, last, p)
}
