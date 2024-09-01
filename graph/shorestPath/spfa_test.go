package shorest_path_test

import (
	"math"
	"testing"
)

// The Shortest-Path Faster Algorithm (SPFA)
// SPFA 是一種用來解決最短路徑問題的演算法

// getSimpleGraph 回傳一個簡單的圖
// n 個節點, 0 ~ n-1
// [節點a, 節點b, 權重]
// [ [0, 1, 1], [0, 2, 2], [1, 3, 3], [2, 3, 4], [3, 4, 5], [4, 5, 6] ]
//
//	0 -(1)- 1 -(3)- 3 -(5)- 4 -(6)- 5
//	|              /
//
// (2)           /
//
//	|          /
//	2  -(4)- /
//	return
//
// map[[2]int{a, b}] int // a -> b 的權重
func getSimpleGraph() map[[2]int]int {
	return map[[2]int]int{
		{0, 1}: 1, {1, 0}: 1,
		{0, 2}: 2, {2, 0}: 2,
		{1, 3}: 3, {3, 1}: 3,
		{2, 3}: 4, {3, 2}: 4, // 切換這個注解，可以改變 0 -> 5最短的路徑
		// {2, 3}: 1, {3, 2}: 1, // 切換這個注解，可以改變 0 -> 5最短的路徑
		{3, 4}: 5, {4, 3}: 5,
		{4, 5}: 6, {5, 4}: 6,
	}
}

// TestSPFA 範例
func TestSPFA(t *testing.T) {
	graph := getSimpleGraph()
	start := 0
	end := 5
	nodes := SPFA(graph, start, end)
	t.Log(nodes)
}

// SPFA 範例
// SPFA 是一種用來解決最短路徑問題的演算法
// 因為與原點的距離會不斷的增加，所以不會循環
// 但是當節點的數量過多的時候，會有很多重覆的訪問 (因為一但更新了，就要加入queue)
// 可以參考dijsktra，這個會利用priority queue來解決這個問題
func SPFA(graph map[[2]int]int, start, end int) int {
	path := map[int][]int{} // 表示 某個點可以到其他哪些點
	for k, _ := range graph {
		path[k[0]] = append(path[k[0]], k[1])
	}

	getNeighbor := func(n int) []int {
		return path[n]
	}

	BFS := func(root int, target int) int {
		q := []int{root}

		distance := make([]int, len(graph)) // 定義從原點到某個點的距離
		for i := 0; i < len(graph); i++ {   // 初始化 先把距離都填最大值
			distance[i] = math.MaxInt32
		}

		distance[root] = 0 // 起點肯定是 0的

		for len(q) > 0 {
			n := q[0]
			q = q[1:]

			for _, neighbor := range getNeighbor(n) {
				// 目前到這個鄰居的距離  < 到現在個點的距離 + 這個點到鄰居的距離
				if distance[neighbor] < distance[n]+graph[[2]int{n, neighbor}] {
					// 不更新，因為現走現在這個點沒有比較快
					continue
				}

				// 更新
				distance[neighbor] = distance[n] + graph[[2]int{n, neighbor}]
				// 因為路徑更短了，所以要把這個點加入queue
				q = append(q, neighbor)
			}
		}

		return distance[target]
	}

	return BFS(start, end)
}
