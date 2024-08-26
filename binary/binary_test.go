package binary_test

import "testing"

// TestXOR 這是一個XOR的簡單測試，可以知道XOR的特性
func TestXOR(t *testing.T) {
	a := 1
	b := 2
	c := 3 // note: c = a ^ b

	t.Log(a^b == b^a) // true
	t.Log(a^b^b == a) // true
	t.Log(a^b^a == b) // true
	t.Log(c == a^b)   // true
	t.Log(c^a == b)   // true
	t.Log(c^b == a)   // true
}

// TestCountOnes 算出 n 的二進位表示法中有幾個 1
func TestCountOnes(t *testing.T) {
	countOnes := func(n int) int {
		count := 0

		for n > 0 {
			n &= (n - 1)
			count++

			// if 1 & n == 1 {
			// 	count++
			// }
			// n >>= 1
		}
		return count
	}

	t.Log(countOnes(0)) // 0
	t.Log(countOnes(1)) // 1
	t.Log(countOnes(9)) // 2
}

// TestBinaryPlus 二進位加法，不需要+運算子
func TestBinaryPlus(t *testing.T) {
	addFunc := func(a, b int) int {
		for b != 0 {
			carry := a & b
			a = a ^ b
			b = carry << 1
		}

		return a
	}

	t.Log(addFunc(1, 2)) // 3
	t.Log(addFunc(2, 3)) // 5
}
