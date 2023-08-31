package graph_test

import (
	"container/list"
)

type Node struct {
	val       int
	neighbors []*Node
}

func BFSList(root *Node, target int) *Node {
	queue := list.New()
	queue.PushBack(root)

	visited := map[*Node]bool{root: true} // 圖是有環的，與樹相比需要防止Loop, 記得root要先放

	for queue.Len() > 0 {
		node := queue.Front().Value.(*Node)
		queue.Remove(queue.Front())

		if node.val == target {
			return node
		}

		for _, n := range node.neighbors { // 與樹需要走子樹，圖需要走鄰居
			if visited[n] {
				continue
			}

			visited[n] = true
			queue.PushBack(n)
		}
	}

	return nil
}

func BFSListEachLevel(root *Node, target int) *Node {
	queue := list.New()
	queue.PushBack(root)
	visited := map[*Node]bool{root: true}
	level := 0

	for queue.Len() > 0 {
		l := queue.Len()
		for i := 0; i < l; i++ {
			node := queue.Front().Value.(*Node)
			queue.Remove(queue.Front())

			if node.val == target {
				return node
			}

			for _, n := range node.neighbors {
				if visited[n] {
					continue
				}

				visited[n] = true
				queue.PushBack(n)
			}
		}
		level++
	}

	return nil
}

func BFSSlice(graph [][]int) {
	visited := make(map[int]bool)
	queue := []int{0}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for _, n := range graph[node] {
			if visited[n] {
				continue
			}

			visited[node] = true
			queue = append(queue, n)
		}
	}
}

func DFS(root *Node, visited map[*Node]bool, target int) *Node {
	// 相比於樹，不會有Nil的情況，每個neighbor都會是一個存在的頂點
	// if root == nil {
	// 	return nil
	// }

	if root.val == target {
		return root
	}

	for _, n := range root.neighbors {
		if visited[n] {
			continue
		}

		visited[root] = true
		DFS(n, visited, target) // 這裡可以判斷返回值，但為了簡單的範例，不寫上去
	}

	return nil
	// DFS(root, map[*Node]bool{root: true}, target)
}

func DFSSlice(graph [][]int, visited map[int]bool) {
	for _, n := range graph[0] {
		if visited[n] {
			continue
		}

		visited[n] = true
		DFSSlice(graph, visited)
	}

	// DFS(graph, map[*Node]bool{graph[0]: true}, target)
}
