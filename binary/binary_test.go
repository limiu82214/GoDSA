package binary_test

import "fmt"

// CountOnes 算出 n 的二進位表示法中有幾個 1
func CountOnes(n int) int {
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

func Basic() {
	a := 1
	b := 2
	c := 3 // note: c = a ^ b

	fmt.Println(a^b == b^a) // true
	fmt.Println(a^b^b == a) // true
	fmt.Println(a^b^a == b) // true
	fmt.Println(c == a^b)   // true
	fmt.Println(c^a == b)   // true
	fmt.Println(c^b == a)   // true
}
