package graph_test

import (
	"container/list"
	"log"
	"testing"
)

// getSampleGraph 回傳一個測試用的圖
//
// 5 -> 2 -> 3
// ↓         |
// 0        /
// ↑      ↙
// 4 -> 1
func getSampleGraph() map[int][]int {
	return map[int][]int{
		5: {2, 0}, // 5 -> 2, 5 -> 0
		4: {0, 1}, // 4 -> 0, 4 -> 1
		2: {3},    // 2 -> 3
		3: {1},    // 3 -> 1
		1: {},
		0: {},
	}
}

// TestKahn 範例
func TestKahn(t *testing.T) {
	graph := getSampleGraph()
	nodes := Kahn(graph)
	t.Log(nodes)

	if len(nodes) == len(graph) {
		// 無環
	} else {
		// 有環
	}
}

// Kahn 範例
// Kahn's algorithm 是一種尋找有向無環圖(DAG)的拓撲排序的演算法
// 可以用來判斷圖是否有環 -> 常用於依賴關係的判斷
func Kahn(graph map[int][]int) []int {
	// 先算出每個node的入度
	findIndegree := func(graph map[int][]int) map[int]int {
		indegree := make(map[int]int)
		for node := range graph { // 初值0
			indegree[node] = 0
		}

		for _, nodes := range graph {
			for _, node := range nodes {
				indegree[node]++
			}
		}

		return indegree
	}
	log.Print(findIndegree(graph))

	topoSort := func() []int {
		res := make([]int, 0)
		indegree := findIndegree(graph)
		queue := list.New()

		// 先把入度為0的node放進queue
		for node, degree := range indegree {
			if degree == 0 {
				queue.PushBack(node)
			}
		}

		for queue.Len() > 0 {
			node := queue.Front().Value.(int)
			queue.Remove(queue.Front())
			log.Print("visit: ", node)

			res = append(res, node) // 因為入度是0，表示沒有其他node指向這個node，所以可以直接放進res

			for _, neighbor := range graph[node] {
				indegree[neighbor]--
				if indegree[neighbor] == 0 {
					queue.PushBack(neighbor)
				}
			}
		}

		return res
	}

	return topoSort()
}
