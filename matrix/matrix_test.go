package matrix_test

import (
	"container/list"
	"fmt"
)

// 與graph相比，只是把走邊獨立出來

func Mix(matrix [][]int) {
	matrixRowLen := len(matrix)
	matrixColLen := len(matrix[0])
	getNeighbors := func(p [2]int) [][2]int { // 可以理解為走邊，整理要走的每個方向
		r, c := p[0], p[1]
		deltaRow := []int{-1, 0, 1, 0}
		deltaCol := []int{0, 1, 0, -1}
		res := [][2]int{}
		for i, _ := range deltaRow {
			neighborRow := r + deltaRow[i]
			neighborCol := c + deltaCol[i]
			if 0 <= neighborRow && neighborRow < matrixRowLen &&
				0 <= neighborCol && neighborCol < matrixColLen {
				res = append(res, [2]int{neighborRow, neighborCol})
			}
		}
		return res
	}

	BFS := func(row, col int) {
		queue := list.New()
		queue.PushBack([2]int{row, col})
		visited := map[[2]int]bool{[2]int{row, col}: true}

		for queue.Len() > 0 {
			node := queue.Front().Value.([2]int)
			queue.Remove(queue.Front())

			for _, n := range getNeighbors(node) { // 原本是走每個點的邊，現在照matrix訂好可以走的路線Mix
				if visited[n] {
					continue
				}

				visited[n] = true
				queue.PushBack(n)
			}
		}
	}

	BFS(0, 0)
}

func DFSMatrix(row, col, n int) {
	if col == n {
		row, col = row+1, 0
	}
	if row == n {
		return // 走完啦
	}

	fmt.Printf("Visiting cell at row %d, col %d\n", row, col)
	// Visiting cell at row 0, col 0
	// Visiting cell at row 0, col 1
	// Visiting cell at row 0, col 2
	// Visiting cell at row 1, col 0
	// Visiting cell at row 1, col 1
	// Visiting cell at row 1, col 2
	// Visiting cell at row 2, col 0
	// Visiting cell at row 2, col 1
	// Visiting cell at row 2, col 2

	DFSMatrix(row, col+1, n)
}
