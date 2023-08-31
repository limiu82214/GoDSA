package graph_test

import "container/list"

func Kahn(graph map[int][]int) []int {
	findIndegree := func(graph map[int][]int) map[int]int {
		indegree := make(map[int]int)
		for node := range graph {
			indegree[node] = 0
		}
		for _, nodes := range graph {
			for _, node := range nodes {
				indegree[node]++
			}
		}
		return indegree
	}

	topoSort := func() []int {
		res := make([]int, 0)
		indegree := findIndegree(graph)
		queue := list.New()

		for node, degree := range indegree {
			if degree == 0 {
				queue.PushBack(node)
			}
		}

		for queue.Len() > 0 {
			node := queue.Front().Value.(int)
			queue.Remove(queue.Front())

			res = append(res, node)

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
