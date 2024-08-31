package lca_test

import "testing"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// getSampleTree 回傳一個測試用的樹
//
//	   1
//	  / \
//	 2   3
//	/ \ / \
//
// 4  5 6  7
func getSampleTree() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}
	//     1
	//    / \
	//   2   3
	//  / \ / \
	// 4  5 6  7
	return root
}

// TestLowestCommonAncestor 範例
func TestLowestCommonAncestor(t *testing.T) {
	root := getSampleTree()
	p := root.Left.Left
	q := root.Right.Right

	t.Log("p: ", p.Val)
	t.Log("q: ", q.Val)

	node := lowestCommonAncestor(root, p, q)
	if node == nil {
		t.Fatal("not found")
	}

	t.Log(node.Val)
}

// lowestCommonAncestor 是找出p, q的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var LCA func(root *TreeNode) *TreeNode
	LCA = func(root *TreeNode) *TreeNode {
		if root == nil { // 空樹
			return nil
		}

		if root == p || root == q { // 根節點是p或q, 另一個節點定在子樹中
			return root
		}

		// 如果不是根節點，那就一定在左子樹或右子樹中
		l := LCA(root.Left)
		r := LCA(root.Right)

		// 如果左右子樹都有找到，那就是root
		if l != nil && r != nil {
			return root
		}

		// 在左邊
		if l != nil {
			return l
		}

		// 在右邊
		if r != nil {
			return r
		}

		return nil
	}

	return LCA(root)
}
