package trie_test

type Trie struct {
	childSet map[byte]*Trie
	isEnd    bool
}

func Constructor() Trie {
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
		t := Constructor()
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
