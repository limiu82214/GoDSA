package shorest_path_test

// Dijkstra's Algorith
// 基本上是 是 SPFA 但利用了priority queue 來加速

import (
	"container/heap"
	"math"
	"testing"
)

// TestDijkstra2 範例
func TestDijkstra2(t *testing.T) {
	graph := getSimpleGraph()
	start := 0
	end := 5
	nodes := Dijkstra2(graph, start, end)
	t.Log(nodes)
}

// Dijkstra2 這個版本改善了 Dijkstra1 存了兩分distance的問題
// 只使用 allNodes
func Dijkstra2(graph map[[2]int]int, start, end int) int {
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

		allNodes[root].distance = 0
		heap.Fix(minQueue, allNodes[root].heapIdx)

		for len(*minQueue) > 0 {
			n := heap.Pop(minQueue).(*nodeDistance)

			for _, neighbor := range getNeighbor(n.idx) {
				// 目前到這個鄰居的距離  < 到現在個點的距離 + 這個點到鄰居的距離
				if allNodes[neighbor].distance < allNodes[n.idx].distance+graph[[2]int{n.idx, neighbor}] {
					// 不更新，因為現走現在這個點沒有比較快
					continue
				}

				// 更新
				allNodes[neighbor].distance = allNodes[n.idx].distance + graph[[2]int{n.idx, neighbor}]
				heap.Fix(minQueue, allNodes[neighbor].heapIdx) // O(log n)
			}
		}

		return allNodes[target].distance
	}

	return BFS(start, end)
}
