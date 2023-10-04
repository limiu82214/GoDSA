package lca_test

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    var LCA func (root *TreeNode) *TreeNode
    LCA = func(root *TreeNode) *TreeNode {
        if root == nil { return nil }
        if root == p  || root == q {
            return root
        }

        l := LCA(root.Left)
        r := LCA(root.Right)

        if l!=nil && r!=nil { return root }
        if l!=nil { return l }
        if r!=nil { return r }

        return nil
    }

    return LCA(root)
}
