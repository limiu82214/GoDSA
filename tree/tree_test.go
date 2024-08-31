package tree_test

import (
	"container/list"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 二元樹
type BNode struct {
	val   int
	left  *BNode
	right *BNode
}

// 多元樹
type Node struct {
	val      int
	children []*Node
}

// getSampleTree 回傳一個測試用的樹
//
//	   1
//	  / \
//	 2   3
//	/ \ / \
// 4  5 6  7
func getSampleTree() *BNode {
	root := &BNode{val: 1}
	root.left = &BNode{val: 2}
	root.right = &BNode{val: 3}
	root.left.left = &BNode{val: 4}
	root.left.right = &BNode{val: 5}
	root.right.left = &BNode{val: 6}
	root.right.right = &BNode{val: 7}
	//     1
	//    / \
	//   2   3
	//  / \ / \
	// 4  5 6  7
	return root
}

// TestDFS DFS範例
func TestDFS(t *testing.T) {
	root := getSampleTree()
	node := DFS(root, 5)
	assert.Equal(t, 5, node.val)
}

// DFS 範例
// 不容易出錯的方式是`一次只處理一個Node`
func DFS(root *BNode, target int) *BNode {
	if root == nil {
		return nil
	}
	log.Print("visit: ", root.val)

	// 先取值
	if root.val == target {
		return root
	}

	// 再走左
	left := DFS(root.left, target)
	if left != nil {
		return left
	}

	// 然後右
	right := DFS(root.right, target)

	return right
}

// TestBFS BFS範例
func TestBFS(t *testing.T) {
	root := getSampleTree()
	node := BFS(root, 5)
	assert.Equal(t, 5, node.val)
}

// BFS 範例
func BFS(root *BNode, target int) *BNode {
	queue := []*BNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		log.Print("visit: ", node.val)

		if node.val == target {
			return node
		}

		if node.left != nil {
			queue = append(queue, node.left)
		}

		if node.right != nil {
			queue = append(queue, node.right)
		}
	}

	return nil
}

// TestBFSEachLevel BFS 每層範例
func TestBFSEachLevel(t *testing.T) {
	root := getSampleTree()
	node := BFSEachLevel(root, 5)
	assert.Equal(t, 5, node.val)
}

// BFSEachLevel BFS 每層範例
// 這個方法可以用來解決一些需要知道深度的問題
func BFSEachLevel(root *BNode, target int) *BNode {
	queue := []*BNode{root}
	level := 0

	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]
			log.Print(" level: ", level, "visit: ", node.val)

			if node.val == target {
				return node
			}

			if node.left != nil {
				queue = append(queue, node.left)
			}

			if node.right != nil {
				queue = append(queue, node.right)
			}
		}

		level++
	}

	return nil
}

// BFSList 用LinkedList實作BFS
func BFSList(root *BNode, target int) *BNode {
	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		node := queue.Front().Value.(*BNode)
		queue.Remove(queue.Front())

		if node.val == target {
			return node
		}

		if node.left != nil {
			queue.PushBack(node.left)
		}

		if node.right != nil {
			queue.PushBack(node.right)
		}
	}

	return nil
}

func BFSListEachLevel(root *BNode, target int) *BNode {
	queue := list.New()
	queue.PushBack(root)

	level := 0 // level與深度

	for queue.Len() > 0 {
		l := queue.Len() // 每次都確定走完一層，這個很重要！！
		for i := 0; i < l; i++ {
			node := queue.Front().Value.(*BNode)
			queue.Remove(queue.Front())

			if node.val == target {
				return node
			}

			if node.left != nil {
				queue.PushBack(node.left)
			}

			if node.right != nil {
				queue.PushBack(node.right)
			}
		}

		level++ // 每走一層深度就+1
	}

	return nil
}
