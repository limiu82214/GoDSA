package array_test

import (
	"testing"
)

// TestMapInit 這是一個簡單的Map初始化
func TestMapInit(t *testing.T) {
	m := make(map[int]bool)

	t.Log(m[20]) // false

	v, ok := m[20]
	t.Log(v, ok) // false false

	// ------

	m[20] = true
	t.Log(m[20]) // true

	v, ok = m[20]
	t.Log(v, ok) // true true
}

// TestMapInit2 map初始化是0值，bool的零值的是false
// map 在在golang是無序的
func TestMapInit2(t *testing.T) {
	n := 4
	m := make([][]bool, n)

	for i := 0; i < n; i++ {
		m[i] = make([]bool, n)
	}

	t.Log(m)
	// fmt.Println(m) [[false false false false][false false false false] [false false false false] [false false false false]]
}
