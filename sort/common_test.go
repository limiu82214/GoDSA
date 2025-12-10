package sort_test

import (
	"log"
	"slices"
	"sort"
	"testing"
)

// TestSorts 測試排序整數 (old)
func TestSorts(t *testing.T) {
	numbers := []int{4, 2, 7, 1, 9, 3}
	sort.Ints(numbers)   // 使用 sort.Ints 函數進行排序
	log.Println("sort.Ints: ", numbers) // 輸出: [1 2 3 4 7 9]

	sort.Sort(sort.Reverse(sort.IntSlice(numbers))) // 使用 sort.Sort 函數進行排序
	log.Println("sort.Sort,Reverse: ", numbers) // 輸出: [9 7 4 3 2 1]

	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})
	log.Println("CusSort `<`: ", numbers) // 輸出: [1 2 3 4 7 9]

	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] > numbers[j]
	})
	log.Println("CusSort `>`: ", numbers) // 輸出: [9 7 4 3 2 1]
}


// TestSliceSort 測試排序(1.21後推薦)
func TestSortFunc(t *testing.T) {
	numbers := []int{4, 2, 7, 1, 9, 3}

	slices.Sort(numbers)
	log.Println("slices.Sort: ", numbers) // 輸出: [1 2 3 4 7 9]

	slices.Reverse(numbers)
	log.Println("slices.Reverse: ", numbers) // 輸出: [9 7 4 3 2 1 ]

	slices.SortFunc(numbers, func(a, b int) int { // func, -1表示結果a要在b前，1表示結果a要在b後，0表示a與b一致
		return a-b
		// if a < b {return -1}
		// if a > b {return 1}
		// return 0
	})
	log.Println("slices.SortFunc `a-b`:", numbers)

	slices.SortFunc(numbers, func(a, b int) int {
		return b-a
		// if a > b {return -1}
		// if a < b {return 1}
		// return 0
	})
	log.Println("slices.SortFunc `b-a`:", numbers)

}
