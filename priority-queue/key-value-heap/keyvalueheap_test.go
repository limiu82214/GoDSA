package keyvalueheap

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestMinKVHeap 範例
func TestMinKVHeap(t *testing.T) {
	minKVHeap := &MinKVHeap{} // 背
	// t := make(MinHeap, 0)
	// minHeap = &t
	heap.Init(minKVHeap) // 背

	values := []KV{
		KV{key: 5, value: "data:5"},
		KV{key: 3, value: "data:3"},
		KV{key: 8, value: "data:8"},
		KV{key: 2, value: "data:2"},
		KV{key: 1, value: "data:1"},
		KV{key: 7, value: "data:7"},
		KV{key: 6, value: "data:6"},
		KV{key: 4, value: "data:4"},
	}
	t.Log("values:", values)
	t.Log("original minKVHeap:", minKVHeap)

	for _, val := range values {
		heap.Push(minKVHeap, val)
	}
	t.Log("pushed minKVHeap:", minKVHeap)

	// (*minHeap)[0] = 1 // 取Top()的
	t.Log("Top", (*minKVHeap)[0])

	expected := []KV{
		KV{key: 1, value: "data:1"},
		KV{key: 2, value: "data:2"},
		KV{key: 3, value: "data:3"},
		KV{key: 4, value: "data:4"},
		KV{key: 5, value: "data:5"},
		KV{key: 6, value: "data:6"},
		KV{key: 7, value: "data:7"},
		KV{key: 8, value: "data:8"},
	}

	t.Log("expected:", expected)
	var result []KV
	for minKVHeap.Len() > 0 {
		kv := heap.Pop(minKVHeap).(KV) // 記的要斷言
		result = append(result, kv)
	}
	t.Log("result:", result)

	assert.Equal(t, expected, result, "Heap should be sorted in ascending order")
}

// 實作比較key，儲存value的minHeap (golang的原生是比較value，儲存value, 如[]int)
// 實作
type KV struct {
	key   int    // 這個會被用來比較heap
	value string // 這個是實際代表的值/也可以放結構
}
type MinKVHeap []KV

func (h MinKVHeap) Len() int           { return len(h) }
func (h MinKVHeap) Less(i, j int) bool { return h[i].key < h[j].key }
func (h MinKVHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinKVHeap) Push(x interface{}) {
	*h = append(*h, x.(KV))
}

func (h *MinKVHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// 使用額外的方法，讓事情變的簡單
func (h *MinKVHeap) PushKV(key int, value string) {
	heap.Push(h, KV{key: key, value: value}) // 熟練可以直接在外面寫這行
}

func (h *MinKVHeap) PopValue() string {
	item := heap.Pop(h).(KV) // 熟練可以直接在外面寫這行
	return item.value
}
