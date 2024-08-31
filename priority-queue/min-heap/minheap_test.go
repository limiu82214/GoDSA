package minheap

import (
	"container/heap" // 背
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestMinHeap 範例
func TestMinHeap(t *testing.T) {
	minHeap := &MinHeap{} // 背
	// t := make(MinHeap, 0)
	// minHeap = &t
	heap.Init(minHeap) // 背

	values := []int{5, 3, 8, 2, 1, 7, 6, 4}
	for _, val := range values {
		heap.Push(minHeap, val)
	}

	expected := []int{1, 2, 3, 4, 5, 6, 7, 8}
	var result []int
	for minHeap.Len() > 0 {
		val := heap.Pop(minHeap).(int) // 記的要斷言
		result = append(result, val)
	}

	// (*minHeap)[0] = 1 // 取Top()的
	assert.Equal(t, expected, result, "Heap should be sorted in ascending order")
}

// O(log n)
// 如果想要做出不同定義的queue，通常觀注在Less()
// 實作
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push O(log n)
// golang 會在你Push之後做上浮的算法
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop O(1) + O(log n)
// golang 會先將根元素與最後一個元素交換，然後呼叫Pop()，最後做下沉的算法
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]

	return x
}
