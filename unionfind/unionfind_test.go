package unionfind

type UnionFind struct {
	set map[int]int // key是id，value是parentID
	// key 與 value 相同的話，表示是root
}

func NewUnionFind() *UnionFind {
	return &UnionFind{set: make(map[int]int)}
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
