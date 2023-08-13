package keyvalueheap

import "container/heap"

// 實作比較key，儲存value的minheap (golang的原生是比較value，儲存value)
// 實作
type KV struct {
	key   int // 這個會被用來比較heap
	value int // 這個是實際代表的值/也可以放結構
}

type MinHeap []KV

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].key < h[j].key }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(KV))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// 使用額外的方法，讓事情變的簡單
func (h *MinHeap) PushKV(key, value int) {
	heap.Push(h, KV{key: key, value: value}) // 熟練可以直接在外面寫這行
}

func (h *MinHeap) PopValue() int {
	item := heap.Pop(h).(KV) // 熟練可以直接在外面寫這行
	return item.value
}
