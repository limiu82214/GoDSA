package graph_test

import (
	"container/list"
	"log"
	"testing"
)

type Node struct {
	val       int
	neighbors []*Node
}

// 圖/樹上的最長直徑
// 1. 從任一點起點A走最長，走到最遠的點B
// 2. 從B 再走一次最長，得到"直徑"
// 因為最長的直徑的兩端一定是離圖/樹中心最遠的位置，而從任一點出發走最遠，一定會走到離樹中心最遠的位置。
//   再從這個位置出發就可以走到另一點最遠的位置。

// getSampleListGraph 回傳一個測試用的圖
//
// 0 - 1 - 3
// |   |   |
// 2 - 4 - 5
func getSampleListGraph() *Node {
	n0 := &Node{val: 0}
	n1 := &Node{val: 1}
	n2 := &Node{val: 2}
	n3 := &Node{val: 3}
	n4 := &Node{val: 4}
	n5 := &Node{val: 5}

	n0.neighbors = []*Node{n1, n2}
	n1.neighbors = []*Node{n0, n3, n4}
	n2.neighbors = []*Node{n0, n5}
	n3.neighbors = []*Node{n1}
	n4.neighbors = []*Node{n1}
	n5.neighbors = []*Node{n2}

	return n0
}

// TestBFSList BFSList範例
func TestBFSList(t *testing.T) {
	root := getSampleListGraph()
	node := BFSList(root, 5)

	if node == nil {
		t.Fatal("not found")
	}
}

// BFSList 範例
func BFSList(root *Node, target int) *Node {
	queue := list.New()
	queue.PushBack(root)

	visited := map[*Node]bool{root: true} // 圖是有環的，與樹相比需要防止Loop, 記得root要先放

	for queue.Len() > 0 {
		node := queue.Front().Value.(*Node)
		queue.Remove(queue.Front())
		log.Print("visit: ", node.val)

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

// TestBFSListEachLevel BFSListEachLevel範例
func TestBFSListEachLevel(t *testing.T) {
	root := getSampleListGraph()
	node := BFSListEachLevel(root, 5)

	if node == nil {
		t.Fatal("not found")
	}
}

// TestBFSListEachLevel BFSListEachLevel範例
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
			log.Print(" level: ", level, ", visit: ", node.val)

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

// TestBFSSlice BFSSlice範例
// 0 - 1 - 3
// |   |   |
// 2 - 4 - 5
func getSampleSliceGraph() [][]int {
	return [][]int{
		{1, 2},
		{0, 3, 4},
		{0, 5},
		{1},
		{1},
		{2},
	}
}

// TestBFSSlice BFSSlice範例
func TestBFSSlice(t *testing.T) {
	graph := getSampleSliceGraph()
	BFSSlice(graph, 1)
}

// BFSSlice 範例
func BFSSlice(graph [][]int, root int) {
	visited := make(map[int]bool)
	queue := []int{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		log.Print("visit: ", node)

		for _, n := range graph[node] {
			if visited[n] {
				continue
			}

			visited[node] = true
			queue = append(queue, n)
		}
	}
}

// TestBFSSliceEachLevel BFSSliceEachLevel範例
func TestBFSSliceEachLevel(t *testing.T) {
	graph := getSampleSliceGraph()
	BFSSliceEachLevel(graph, 1)
}

// BFSSliceEachLevel 範例
func BFSSliceEachLevel(graph [][]int, root int) {
	visited := make(map[int]bool)
	queue := []int{root}
	level := 0

	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]
			log.Print(" level: ", level, ", visit: ", node)

			for _, n := range graph[node] {
				if visited[n] {
					continue
				}

				visited[node] = true
				queue = append(queue, n)
			}
		}

		level++
	}
}

// TestDFSList DFSList範例
func TestDFSList(t *testing.T) {
	root := getSampleListGraph()
	node := DFSList(root, map[*Node]bool{}, 5)

	if node == nil {
		t.Fatal("not found")
	}
	log.Print("found: ", node.val)
}

// DFSList 範例
func DFSList(root *Node, visited map[*Node]bool, target int) *Node {
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
		rst := DFSList(n, visited, target)
		if rst != nil {
			return rst
		}
	}

	return nil
}

// TestDFSSlice DFSSlice範例
func TestDFSSlice(t *testing.T) {
	graph := getSampleSliceGraph()
	DFSSlice(graph, map[int]bool{0: true}, 0)
}

// DFSSlice 範例
func DFSSlice(graph [][]int, visited map[int]bool, root int) {
	log.Print("visit: ", root, ", visited: ", visited)
	// log.Print("graph: ", graph[root])

	for _, n := range graph[root] {
		if visited[n] {
			continue
		}

		visited[n] = true
		DFSSlice(graph, visited, n)
	}
}
