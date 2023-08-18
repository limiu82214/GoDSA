package sort_test

import (
	"sort"
	"testing"
)

func TestSortInt(t *testing.T) {
	numbers := []int{4, 2, 7, 1, 9, 3}
	sort.Ints(numbers) // 使用 sort.Ints 函數進行排序
	// fmt.Println(numbers)  // 輸出: [1 2 3 4 7 9]
}
