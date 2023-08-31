package tree_test

import "container/list"

type BNode struct {
	val   int
	left  *BNode
	right *BNode
}

// 不容易出錯的方式是`一次只處理一個Node`

func DFS(root *BNode, target int) *BNode {
	if root == nil {
		return nil
	}

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

func BFS(root *BNode, target int) *BNode {
	queue := []*BNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

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

type Node struct {
	val      int
	children []*Node
}
