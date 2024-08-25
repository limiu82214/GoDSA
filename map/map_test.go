package map_test

import (
	"testing"
)

// TestMapWithArrayKey 討論Map的Key使用Array，在簡單的描述x, y座標很有用
func TestMapWithArrayKey(t *testing.T) {
	m := map[[2]int]bool{} // map的key只能用array，不能用slice
	t.Log(m)               // map[]

	m2 := map[[2]int]bool{[2]int{8, 7}: true} // 初值
	t.Log(m2)                                 // map[[8 7]:true]

	m2[[2]int{8, 7}] = false
	t.Log(m2) // map[[8 7]:false]

	m2[[2]int{7, 8}] = true // array的值會被當成key，順序是有差的，*不是set的概念*
	t.Log(m2)               // map[[7 8]:true [8 7]:false]
}
