package trie_test

import (
	"encoding/json"
	"testing"
)

type Trie struct {
	childSet map[byte]*Trie
	isEnd    bool
}

// TestTrie 範例
func TestTrie(t *testing.T) {
	trie := NewTrie()
	trie.Insert("apple")
	if !trie.Search("apple") {
		t.Fatal("search apple failed")
	}
	PrintTrie(&trie)

	if trie.Search("app") {
		t.Fatal("search app failed, app is not in trie")
	}

	if !trie.StartsWith("app") {
		t.Fatal("start with app failed, start with app is in trie")
	}

	trie.Insert("app")
	if !trie.Search("app") {
		t.Fatal("search app failed")
	}
	println("------------")

	trie.Insert("appends")
	PrintTrie(&trie)

}

func PrintTrie(t *Trie) {
	b, _ := json.Marshal(t.childSet)
	println("childSet: ", string(b))
	for k, v := range t.childSet {
		println("key: ", string(k))
		PrintTrie(v)
	}
}

func NewTrie() Trie {
	return Trie{
		childSet: make(map[byte]*Trie),
		isEnd:    false,
	}
}

func (this *Trie) Insert(word string) {
	if word == "" {
		this.isEnd = true
		return
	}
	if _, ok := this.childSet[word[0]]; !ok {
		t := NewTrie()
		this.childSet[word[0]] = &t
	}

	this.childSet[word[0]].Insert(string(word[1:]))
}

func (this *Trie) Search(word string) bool {
	if word == "" {
		if this.isEnd {
			return true
		}
		return false
	}
	if _, ok := this.childSet[word[0]]; ok {
		return this.childSet[word[0]].Search(string(word[1:]))
	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	if prefix == "" {
		return true
	}
	if _, ok := this.childSet[prefix[0]]; ok {
		return this.childSet[prefix[0]].StartsWith(string(prefix[1:]))
	}
	return false
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
