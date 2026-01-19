package unionfind

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// UnionFind 是用來處理交集的資料結構
// 他可以用來判斷兩個元素是否在同一個集合中
// 也可以用來合併兩個集合
// 透過 path compression + union by size，
// 所有操作的攤銷時間複雜度接近 O(1)（嚴格為 O(α(n))）
type UnionFind struct {
	// key 與 value 相同的話，表示是root
	fa   map[int]int // key: node, value: parent
	size map[int]int // 只有 root 的 size 有意義
}

func NewUnionFind() *UnionFind {
	return &UnionFind{
		fa:   make(map[int]int),
		size: make(map[int]int),
	}
}

// TestUnionFind 範例
func TestUnionFind(t *testing.T) {
	uf := NewUnionFind()
	uf.Union(1, 2)
	uf.Union(2, 3)

	assert.Equal(t, uf.Find(1), uf.Find(2))

	uf.Union(4, 5)
	uf.Union(5, 6)
	assert.Equal(t, uf.Find(4), uf.Find(6))

	uf.Union(3, 6)
	assert.Equal(t, uf.Find(1), uf.Find(6))
	assert.Equal(t, uf.Find(1), uf.Find(3))
	assert.Equal(t, uf.Find(1), uf.Find(5))
}

func (uf *UnionFind) Find(node int) (root int) {
    // 新節點
    if _, ok := uf.fa[node]; !ok {
        uf.fa[node] = node // 獨立節點自己就是parent
		uf.size[node] = 1
        return node
    }

	if uf.fa[node] != node {
	    uf.fa[node] = uf.Find(uf.fa[node])
	}
    return uf.fa[node]
}

func (uf *UnionFind) Union(n1, n2 int) {
    // 把兩個根合在一起
	r1, r2 := uf.Find(n1), uf.Find(n2)
	if r1 == r2 {
		return
	}
	uf.fa[r2] = r1 // r2 併進 r1
	uf.size[r1] += uf.size[r2] // r1 + size
}

func (uf *UnionFind) Size(node int) int {
	root := uf.Find(node)
	return uf.size[root]
}
