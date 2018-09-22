package heap

import (
	"github.com/8tomat8/sketches/algorithms/dataStructs"
)

func Sort(s []int) {
	if len(s) < 1 {
		return
	}

	// Algorithm uses standard behavior of Heap data structure,
	// so look for the algorithm implementation in dataStructs package

	// First we need to push all data to heap.
	// For simplicity, data was added on heap init
	heap := dataStructs.NewHeap(s)

	// Second step is to pop all elements from the heap.
	// Heap will return them in ASC/DESC (depends on heap implementation) order
	for i := 0; i < len(s); i++ {
		s[i] = heap.Pop()
	}
}
