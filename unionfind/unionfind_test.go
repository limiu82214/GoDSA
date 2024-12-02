package unionfind

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// UnionFind 是用來處理交集的資料結構
// 他可以用來判斷兩個元素是否在同一個集合中
// 也可以用來合併兩個集合
// 時間複雜度 Union O(1), Find O(log(n))
// 可以快速的知道兩個元素是否在同一個集合中
type UnionFind struct {
	fa map[int]int // key是id，value是parentID
	// key 與 value 相同的話，表示是root
}

func NewUnionFind() *UnionFind {
	return &UnionFind{set: make(map[int]int)}
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

// O(log(n)) n 是圖中的節點數
func (uf *UnionFind) Find(node int) (root int) {
    // 新節點
    if _, ok := uf.fa[node]; !ok {
        uf.fa[node] = node // 獨立節點自己就是parent
        return node
    }

    // 自己是parent表示已經是根了
    if node == uf.fa[node] {
        return node
    }

    // 把自己指到整個set的根 // 這只是優化
    uf.fa[node] = uf.Find(uf.fa[node])

    return uf.fa[node]
}

func (uf *UnionFind) Union(n1, n2 int) {
    // 把兩個根合在一起
    uf.fa[uf.Find(n1)] = uf.Find(n2)
}

