package shorest_path_test

// Dijkstra's Algorith
// 基本上是 是 SPFA 但利用了priority queue 來加速

import (
	"container/heap"
	"math"
	"testing"
)

// TestDijkstra2_1 範例
func TestDijkstra2_1(t *testing.T) {
	graph := getSimpleGraph()
	start := 0
	end := 5
	nodes := Dijkstra2_1(graph, start, end)
	t.Log(nodes)
}

// Dijkstra2_1 這個版本改善了 Dijkstra1 存了兩分distance的問題
// 只使用 distance
func Dijkstra2_1(graph map[[2]int]int, start, end int) int {
	path := map[int][]int{} // 表示 某個點可以到其他哪些點
	for k, _ := range graph {
		path[k[0]] = append(path[k[0]], k[1])
	}

	getNeighbor := func(n int) []int {
		return path[n]
	}

	BFS := func(root int, target int) int {
		// 不使用直接使用minQueue 代替原本的queue
		minQueue := &minHeap{} // 記錄目前最短路徑的節點，這樣可以快速找到最短路徑
		heap.Init(minQueue)

		distance := make([]int, len(graph)) // 定義從原點到某個點的距離
		for i := 0; i < len(graph); i++ {   // 初始化 先把距離都填最大值
			distance[i] = math.MaxInt32
		}

		distance[root] = 0 // 起點肯定是 0的
		heap.Push(minQueue, &nodeDistance{idx: root, distance: 0})

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

				// 因為路徑更短了，所以要把這個點加入queue
				heap.Push(minQueue, &nodeDistance{idx: neighbor, distance: distance[neighbor]})
			}
		}

		return distance[target]
	}

	return BFS(start, end)
}
