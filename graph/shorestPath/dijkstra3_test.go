package shorest_path_test

// Dijkstra's Algorith
// 基本上是 是 SPFA 但利用了priority queue 來加速

import (
	"container/heap"
	"math"
	"testing"
)

// TestDijkstra 範例
func TestDijkstra3(t *testing.T) {
	graph := getSimpleGraph()
	start := 0
	end := 5
	nodes := Dijkstra3(graph, start, end)
	t.Log(nodes)
}

// Dijkstra3_1
func Dijkstra3(graph map[[2]int]int, start, end int) int {
	neighbors := map[int][][2]int{} // 表示 某個點可以到其他哪些點 ex: 0 -> [[1, 2], [2, 8] ...] 表示 0 -> 可以到 1 , 2 且距離分別是 2, 8
	// [2]int 也可以另外宣告一個struct來表示，會清楚
	for k, v := range graph {
		neighbors[k[0]] = append(neighbors[k[0]], [2]int{k[1], v})
	}

	getNeighbor := func(n int) [][2]int {
		return neighbors[n]
	}

	dijk := func(root int, target int) int {
		q := &minHeap3{} // 用minHeap取代queue，這樣每個點都只要走一次，因為每次算出來的一定都是最短的
		heap.Init(q)

		dis := make([]node, len(neighbors)) // 追踨從起點到某點的距離
		for i, _ := range neighbors {
			dis[i] = node{idx: i, distance: math.MaxInt32}
		}

		dis[root].distance = 0
		heap.Push(q, dis[root])

		for len(*q) > 0 {
			n := heap.Pop(q).(node)

			for _, neighbor := range getNeighbor(n.idx) {
				neighborIdx := neighbor[0]
				neighborDis := neighbor[1]

				newDistance := dis[n.idx].distance + neighborDis
				if dis[neighborIdx].distance < newDistance { // 這裡要特別注意題目，以免loop造成TLE
					continue
				}

				// 更新
				dis[neighborIdx].distance = newDistance
				heap.Push(q, node{idx: neighborIdx, distance: newDistance})
			}
		}

		return dis[target].distance
	}

	return dijk(start, end)
}

type node struct {
	idx      int
	distance int
}

type minHeap3 []node

func (h minHeap3) Len() int            { return len(h) }
func (h minHeap3) Less(i, j int) bool  { return h[i].distance < h[j].distance }
func (h minHeap3) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap3) Push(x interface{}) { *h = append(*h, x.(node)) }
func (h *minHeap3) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
