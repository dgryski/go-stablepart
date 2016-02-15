package stablepart

import (
	"fmt"
)

type ints []int

func (is ints) Len() int { return len(is) }

func (is ints) Less(i int, j int) bool {
	return is[i] < is[j]
}

func (is ints) Swap(i int, j int) {
	is[i], is[j] = is[j], is[i]
}

func ExampleStablePartition() {
	var data = ints{6, 9, 8, 1, 0, 5, 2, 4, 3, 7}
	idx := StablePartition(data, 0, len(data), func(i int) bool { return data[i] < 5 })

	fmt.Println(data)
	fmt.Println(idx)
	// Output: [1 0 2 4 3 6 9 8 5 7]
	// 5
}

func ExampleReverse() {
	var data = ints{1, 2, 3, 4, 5}
	Reverse(data, 1, len(data)-1)

	fmt.Println(data)
	// Output: [1 4 3 2 5]
}

func ExampleRotate() {
	var data = ints{1, 2, 3, 4, 5}
	// rotate data[1:4] by so that d[3] is first
	idx := Rotate(data, 1, 3, 4)

	fmt.Println(data)
	fmt.Println(idx)
	// Output: [1 4 2 3 5]
	// 2
}

func ExampleGather() {
	var data = ints{6, 9, 8, 1, 0, 5, 2, 4, 3, 7}
	// gather data < 5 around position 6
	start, end := Gather(data, 0, len(data), 6, func(i int) bool { return data[i] < 5 })

	fmt.Println(data)
	fmt.Println(start, end)
	// Output: [6 9 8 5 1 0 2 4 3 7]
	// 4 9
}

func ExampleSlide() {
	var data = ints{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// slide elements data[2:4] to target position 8
	start, end := Slide(data, 2, 4, 8)

	fmt.Println(data)
	fmt.Println(start, end)
	// Output: [0 1 4 5 6 7 2 3 8 9]
	// 6 8
}
