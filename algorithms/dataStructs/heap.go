package dataStructs

type Heap interface {
	Push(int)
	Pop() int
}

func NewHeap(initialData []int) Heap {
	h := &heap{
		data: make([]int, len(initialData)),
	}

	parent := 0
	h.data[parent] = initialData[0] // putting first slice element on the root position to have value to compare in loop

	for i := 1; i < len(initialData); i += 2 {
		for leaf := 1; leaf <= 2 && i+leaf < len(initialData); leaf++ {
			h.data[parent*2+leaf] = initialData[i+leaf-1]

			h.fixUp(parent*2 + leaf)
		}
		parent++
	}
	return h
}

type heap struct {
	data []int
}

func (h *heap) Push(v int) {
	h.data = append(h.data, v)
	h.fixUp(len(h.data) - 1)
}

func (h *heap) Pop() int {
	res := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]

	h.fixDown(0)
	return res
}

func (h *heap) fixUp(from int) {
	for ; from > 0; from /= 2 {
		if h.data[from] < h.data[from/2] {
			h.data[from], h.data[from/2] = h.data[from/2], h.data[from]
		}
	}
}

func (h *heap) fixDown(from int) {
	for parent := from; parent <= len(h.data)/2; parent++ {
		for leaf := 1; leaf <= 2; leaf++ {
			child := parent*2 + leaf
			if child > len(h.data)-1 {
				break
			}

			if h.data[parent] > h.data[child] {
				h.data[parent], h.data[child] = h.data[child], h.data[parent]
			}
		}
	}
}
