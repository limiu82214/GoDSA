package shorest_path_test

// Dijkstra's Algorith
// 基本上是 是 SPFA 但利用了priority queue 來加速

import (
	"container/heap"
	"math"
	"testing"
)

// TestDijkstra 範例
func TestDijkstra3_1(t *testing.T) {
	graph := getSimpleGraph()
	start := 0
	end := 5
	nodes := Dijkstra3_1(graph, start, end)
	t.Log(nodes)
}

// Dijkstra3_1
// 用 dis 來記錄最短距離
// 用 minHeap 來記錄目前最短的點
// 重點是兩者之間的關系
func Dijkstra3_1(graph map[[2]int]int, start, end int) int {
	neighbors := map[int][][2]int{} // 表示 某個點可以到其他哪些點 ex: 0 -> [[1, 2], [2, 8] ...] 表示 0 -> 可以到 1 , 2 且距離分別是 2, 8
	for k, v := range graph {
		neighbors[k[0]] = append(neighbors[k[0]], [2]int{k[1], v})
	}

	getNeighbor := func(n int) [][2]int {
		return neighbors[n]
	}

	dijk := func(root int, target int) int {
		q := &minHeap3{} // 用minHeap取代queue，這樣每個點都只要走一次，因為每次算出來的一定都是最短的
		heap.Init(q)

		dis := make([]int, len(neighbors)) // 追踨從起點到某點的距離
		for i, _ := range neighbors {
			dis[i] = math.MaxInt32
		}

		dis[root] = 0
		heap.Push(q, node{idx: root, distance: 0})

		for len(*q) > 0 {
			n := heap.Pop(q).(node)

			for _, neighbor := range getNeighbor(n.idx) {
				neighborIdx := neighbor[0]
				neighborDis := neighbor[1]

				if dis[neighborIdx] < dis[n.idx]+neighborDis {
					continue
				}

				// 更新
				dis[neighborIdx] = dis[n.idx] + neighborDis
				heap.Push(q, node{idx: neighborIdx, distance: dis[neighborIdx]})
			}
		}

		return dis[target]
	}

	return dijk(start, end)
}
