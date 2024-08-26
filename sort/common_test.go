package sort_test

import (
	"log"
	"sort"
	"testing"
)

// TestSortInt 測試排序整數
func TestSortInt(t *testing.T) {
	numbers := []int{4, 2, 7, 1, 9, 3}
	sort.Ints(numbers)   // 使用 sort.Ints 函數進行排序
	log.Println(numbers) // 輸出: [1 2 3 4 7 9]

	sort.Sort(sort.Reverse(sort.IntSlice(numbers))) // 使用 sort.Sort 函數進行排序
	log.Println(numbers)                            // 輸出: [9 7 4 3 2 1]
}

// TestSortFunc 測試排序自訂函數 (這是不穩定排序)
func TestSortFunc(t *testing.T) {
	numbers := []int{4, 3, 7, 1, 9, 3}
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})
	log.Println(numbers) // 輸出: [1 3 3 4 7 9]

	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] > numbers[j]
	})
	log.Println(numbers) // 輸出: [9 7 4 3 3 1]
}
