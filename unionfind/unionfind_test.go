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
	set map[int]int // key是id，value是parentID
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
func (uf *UnionFind) Find(id int) int {
	if parentID, exists := uf.set[id]; exists {
		if parentID != id { // 若不是根，應該往上找到根，這裡用遞迴
			uf.set[id] = uf.Find(parentID)
		}

		return uf.set[id]
	}

	return id
}

// O(1)
func (uf *UnionFind) Union(id, id2 int) {
	uf.set[uf.Find(id)] = uf.Find(id2) // 把id的根，接到id2的根上
	// 這裡可以優化，把較深的樹接到較淺的樹上，這樣可以減少Find的時間
}
