package shorest_path_test

// Dijkstra's Algorith
// 基本上是 是 SPFA 但利用了priority queue 來加速

import (
	"container/heap"
	"math"
	"testing"
)

// TestDijkstra1 範例
func TestDijkstra1(t *testing.T) {
	graph := getSimpleGraph()
	start := 0
	end := 5
	nodes := Dijkstra1(graph, start, end)
	t.Log(nodes)
}

// Dijkstra1 這個版本改善了SPFA 沒有利用 權重的問題
// 並且因為都是從最短路徑的點開始，所以每一個點只會走一次
func Dijkstra1(graph map[[2]int]int, start, end int) int {
	path := map[int][]int{} // 表示 某個點可以到其他哪些點
	for k, _ := range graph {
		path[k[0]] = append(path[k[0]], k[1])
	}

	getNeighbor := func(n int) []int {
		return path[n]
	}

	// 所有的節點 // 這樣可以直接改值
	allNodes := []*nodeDistance{}
	for i := 0; i < len(path); i++ { // 如果你有孤立的點，請不要用path
		allNodes = append(allNodes, &nodeDistance{idx: i, distance: math.MaxInt32})
	}

	BFS := func(root int, target int) int {
		// 不使用直接使用minQueue 代替原本的queue
		minQueue := &minHeap{} // 記錄目前最短路徑的節點，這樣可以快速找到最短路徑
		heap.Init(minQueue)
		for _, node := range allNodes {
			heap.Push(minQueue, node)
		}

		distance := make([]int, len(graph)) // 定義從原點到某個點的距離
		for i := 0; i < len(graph); i++ {   // 初始化 先把距離都填最大值
			distance[i] = math.MaxInt32
		}

		distance[allNodes[root].idx] = 0 // 起點肯定是 0的
		allNodes[root].distance = 0
		heap.Fix(minQueue, allNodes[root].heapIdx)

		for len(*minQueue) > 0 {
			n := heap.Pop(minQueue).(*nodeDistance)

			for _, neighbor := range getNeighbor(n.idx) {
				// 目前到這個鄰居的距離  < 到現在個點的距離 + 這個點到鄰居的距離
				if distance[neighbor] < distance[n.idx]+graph[[2]int{n.idx, neighbor}] {
					// 不更新，因為現走現在這個點沒有比較快
					continue
				}

				// 更新
				distance[neighbor] = distance[n.idx] + graph[[2]int{n.idx, neighbor}]
				// 不加到queue了，直接加到minDistance
				// q = append(q, neighbor)
				// 現在更新距離，所以heap要修正排序
				allNodes[neighbor].distance = distance[neighbor]
				heap.Fix(minQueue, allNodes[neighbor].heapIdx) // O(log n)
			}
		}

		return distance[target]
	}

	return BFS(start, end)
}

type nodeDistance struct {
	idx      int
	distance int
	heapIdx  int
}

type minHeap []*nodeDistance

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].distance < h[j].distance }
func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].heapIdx, h[j].heapIdx = h[j].heapIdx, h[i].heapIdx
}
func (h *minHeap) Push(x interface{}) {
	n := len(*h)
	e := x.(*nodeDistance)
	e.heapIdx = n
	*h = append(*h, e)
}
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	x.heapIdx = -1
	*h = old[:n-1]
	return x
}
