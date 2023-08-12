package maxheap

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxHeap(t *testing.T) {
	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)

	values := []int{5, 3, 8, 2, 1, 7, 6, 4}
	for _, val := range values {
		heap.Push(maxHeap, val)
	}

	expected := []int{8, 7, 6, 5, 4, 3, 2, 1}
	var result []int
	for maxHeap.Len() > 0 {
		val := heap.Pop(maxHeap).(int)
		result = append(result, val)
	}

	assert.Equal(t, expected, result, "Heap should be sorted in descending order")
}

// 實作
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] >= h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
