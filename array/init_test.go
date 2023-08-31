package array_test

import "fmt"

func init2Map() {
	n := 4
	m := make([][]bool, n)
	for i := 0; i < n; i++ {
		m[i] = make([]bool, n)
	}

	// fmt.Println(m) [[false false false false] [false false false false] [false false false false] [false false false false]]
}

func mapInit() {
	m := make(map[int]bool)

	fmt.Println(m[20]) // false

	v, ok := m[20]
	fmt.Println(v, ok) // false false

	// ------

	m[20] = true
	fmt.Println(m[20]) // true

	v, ok = m[20]
	fmt.Println(v, ok) // true true
}
