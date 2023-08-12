package minheap

import (
	"container/heap" // 這行是必需的
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinHeap(t *testing.T) {
	minHeap := &MinHeap{}
	heap.Init(minHeap)

	values := []int{5, 3, 8, 2, 1, 7, 6, 4}
	for _, val := range values {
		heap.Push(minHeap, val)
	}

	expected := []int{1, 2, 3, 4, 5, 6, 7, 8}
	var result []int
	for minHeap.Len() > 0 {
		val := heap.Pop(minHeap).(int)
		result = append(result, val)
	}

	assert.Equal(t, expected, result, "Heap should be sorted in ascending order")
}

// 實作
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
