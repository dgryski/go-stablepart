package stablepart

import (
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"
)

type element struct {
	val int
	idx int
}

type elements []element

func (e elements) Len() int           { return len(e) }
func (e elements) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }
func (e elements) Less(i, j int) bool { return e[i].val < e[j].val }

func (e elements) Generate(rand *rand.Rand, size int) reflect.Value {
	// generate random slice
	sz := rand.Intn(size)
	elts := make(elements, sz)
	for i := range elts {
		elts[i].val = rand.Intn(size)
		elts[i].idx = i
	}

	return reflect.ValueOf(elts)
}

func verify(t *testing.T, pre, elts elements, f, l int, which string, p func(i int) bool) bool {

	// check the predicate
	for i := f; i < l; i++ {
		if !p(f) {
			t.Logf("%s predicate failed %v\n", which, elts[0])
			t.Logf("pre =%+v\n", pre)
			t.Logf("elts=%+v\n", elts)
			return false
		}

	}

	// check the order
	for i := f + 1; i < l; i++ {
		if elts[i-1].idx > elts[i].idx {
			t.Logf("%s partition mismatch: %v %v\n", which, elts[i-1], elts[i])
			t.Logf("pre =%+v\n", pre)
			t.Logf("elts=%+v\n", elts)
			return false
		}
	}

	return true
}

func TestStableParition(t *testing.T) {

	check := func(elts elements) bool {

		sz := len(elts)

		// save a copy
		pre := make(elements, sz)
		copy(pre, elts)

		predicate := func(i int) bool { return elts[i].val&1 == 0 }

		// partition
		p := StablePartition(elts, 0, elts.Len(), predicate)

		// verify it's correct
		return true &&
			verify(t, pre, elts, 0, p, "pre", predicate) &&
			verify(t, pre, elts, p, sz, "post", func(i int) bool { return !predicate(i) })
	}

	quick.Check(check, nil)
}
