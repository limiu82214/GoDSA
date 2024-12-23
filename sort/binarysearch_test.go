package sort_test

import "testing"

// TestBinarySearch 測試二分搜尋
func TestBinarySearch(t *testing.T) {
	t.Log(binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5))  // 4
	t.Log(binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 10)) // -1
	t.Log(binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 0))  // -1
}

func binarySearch(arr []int, target int) int {
	l, r := 0, len(arr)-1
	for l <= r { // 相同index的話也不能跳過
		mid := l + (r-l)/2 // 防止 l + r 溢位
		if arr[mid] == target {
			return mid
		}
		if arr[mid] > target {
			r = mid - 1 // -1 表示不保留 mid，應該依照題目決定要不要保留 mid，也可以保留在某個地方再捨棄
		} else {
			l = mid + 1 // +1 表示不保留 mid，應該依照題目決定要不要保留 mid，也可以保留在某個地方再捨棄
		}
	}

	return -1
}

/**
* 如何不陷入無限循環
	* 改動前先要改動的部份有幾處
	* 每一步都取得進步
	* 每一次改動都要確認: 退出策略
*/

// TestMonotonicTrueWithBinayrSearch 測試monotonic 拿第一個True
func TestMonotonicTrueWithBinayrSearch(t *testing.T) {
	t.Log(findFirstTrue([]bool{false, false, false, false, false, false, true, true, true, true}))     // 6
	t.Log(findFirstTrue([]bool{false, false, false, false, false, false, false, false, false, false})) // -1
}

// findFirstTrue 找到第一個True 適用於 F F F F F T T T T
func findFirstTrue(arr []bool) int {
	l, r := 0, len(arr)-1
	firstIdx := -1

	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] == true {
			firstIdx = mid // first True
			r = mid - 1
		} else {
			// firstIdx = mid  // last False
			l = mid + 1
		}
	}

	return firstIdx
}

// TestMonotonicFalseWithBinayrSearch 測試monotonic 拿第一個False
func TestMonotonicFalseWithBinayrSearch(t *testing.T) {
	t.Log(findFirstFalse([]bool{true, true, true, true, true, false, false, false, false}))      // 5
	t.Log(findFirstFalse([]bool{true, true, true, true, true, true, true, true, true}))          // -1
	t.Log(findFirstFalse([]bool{false, false, false, false, false, false, false, false, false})) // 0
}

// findFirstFalse 找到第一個False 適用於 T T T T T F F F F
func findFirstFalse(arr []bool) int {
	l, r := 0, len(arr)-1
	firstIdx := -1

	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] == false {
			firstIdx = mid // first False
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return firstIdx
}

// 觀察 L 與 R 的變化
func Binary(arr []bool) int {
	l, r := 0, len(arr)-1

	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] == false {
			r = mid - 1 // arr[mid] 是 false 移動R，所以R在找true
		} else {
			l = mid + 1 // arr[mid] 是  true 移動L，所以L在找false
		}
	}
	// 一般情況
	//     [true, true, true, true, true, false, false, false, false]
	//       ↑L                                                  ↑R
	//     [true, true, true, true, true, false, false, false, false]
	//                               ↑R    ↑L

	// 全true
	//    [true, true, true, true, true, true, true, true, true]
	//      ↑L                                              ↑R
	//    [true, true, true, true, true, true, true, true, true]
	//                                                      ↑R   ↑L

	// 全false
	//     [false, false, false, false, false, false, false, false, false]
	//       ↑L                                                       ↑R
	//     [false, false, false, false, false, false, false, false, false]
	//  ↑R   ↑L

	fmt.Println(arr)
	fmt.Print("[")
	for i := range arr {
		lc := " "
		rc := " "
		if i == l {
			lc = "L"
		}
		if i == r {
			lc = "R"
		}
		if arr[i] == true {
			fmt.Print(" ", lc, rc, i, " ")
		} else {
			fmt.Print("  ", lc, rc, i, " ")
		}
	}
	fmt.Print("]")
	fmt.Println()
	fmt.Println("L(找false):", l, ", R(找true):", r)

	return l
}


func slicesBinarySearchFunWithBool() {
	// slices.BinarySearch是 [) 意味著每次處理的時候 L 是包含那個i，所以最後的時候 L>=i :L命中時L會等於i，L沒命中時L會等於i+1 (每次處理的時候 L 是包含那個i)
	//                                             R 是不包含那i，              R< i :L命中時R會等於i，L沒命中時R會等於i+1 (每次處理的時候 R 是不包含那i) (因為就算命中了，也不會提前跳出for)

	// --
	m := map[bool]int{true: 0, false: 1}
	arr := []bool{true, true, true, true, true, false, false, false, false}
	idx, isFound := slices.BinarySearchFunc(arr, false, func(a, b bool) int {
		return m[a] - m[b]
	})
	fmt.Println("(true->false) find false: []bool{true, true, true, true, true, false, false, false, false}", "isFound:", isFound, "idx:", idx)
	idx, isFound = slices.BinarySearchFunc(arr, true, func(a, b bool) int {
		return m[a] - m[b]
	})
	fmt.Println("(true->false) find true: []bool{true, true, true, true, true, false, false, false, false}", "isFound:", isFound, "idx:", idx)

	// --
	arr = []bool{true, true, true, true, true, true, true, true, true}
	idx, isFound = slices.BinarySearchFunc(arr, false, func(a, b bool) int {
		return m[a] - m[b]
	})
	fmt.Println("(true->false) find false: []bool{true, true, true, true, true, true, true, true, true}", "isFound:", isFound, "idx:", idx)
	idx, isFound = slices.BinarySearchFunc(arr, true, func(a, b bool) int {
		return m[a] - m[b]
	})
	fmt.Println("(true->false) find true: []bool{true, true, true, true, true, true, true, true, true}", "isFound:", isFound, "idx:", idx)

	// --
	arr = []bool{false, false, false, false, false, false, false, false, false}
	idx, isFound = slices.BinarySearchFunc(arr, false, func(a, b bool) int {
		return m[a] - m[b]
	})
	fmt.Println("(true->false) find false: []bool{false, false, false, false, false, false, false, false, false}", "isFound:", isFound, "idx:", idx)
	idx, isFound = slices.BinarySearchFunc(arr, true, func(a, b bool) int {
		return m[a] - m[b]
	})
	fmt.Println("(true->false) find true: []bool{false, false, false, false, false, false, false, false, false}", "isFound:", isFound, "idx:", idx)

	// --
	m = map[bool]int{false: 0, true: 1}
	arr = []bool{false, false, false, false, false, true, true, true, true}
	idx, isFound = slices.BinarySearchFunc(arr, false, func(a, b bool) int {
		return m[a] - m[b]
	})
	fmt.Println("(false->true) find false: []bool{false, false, false, false, false, true, true, true, true}", "isFound:", isFound, "idx:", idx)
	idx, isFound = slices.BinarySearchFunc(arr, true, func(a, b bool) int {
		return m[a] - m[b]
	})
	fmt.Println("(false->true) find true: []bool{false, false, false, false, false, true, true, true, true}", "isFound:", isFound, "idx:", idx)

	// --
	arr = []bool{true, true, true, true, true, true, true, true, true}
	idx, isFound = slices.BinarySearchFunc(arr, false, func(a, b bool) int {
		return m[a] - m[b]
	})
	fmt.Println("(false->true) find false: []bool{true, true, true, true, true, true, true, true, true}", "isFound:", isFound, "idx:", idx)
	idx, isFound = slices.BinarySearchFunc(arr, true, func(a, b bool) int {
		return m[a] - m[b]
	})
	fmt.Println("(false->true) find true: []bool{true, true, true, true, true, true, true, true, true}", "isFound:", isFound, "idx:", idx)

	// --
	arr = []bool{false, false, false, false, false, false, false, false, false}
	idx, isFound = slices.BinarySearchFunc(arr, false, func(a, b bool) int {
		return m[a] - m[b]
	})
	fmt.Println("(false->true) find false: []bool{false, false, false, false, false, false, false, false, false}", "isFound:", isFound, "idx:", idx)
	idx, isFound = slices.BinarySearchFunc(arr, true, func(a, b bool) int {
		return m[a] - m[b]
	})
	fmt.Println("(false->true) find true: []bool{false, false, false, false, false, false, false, false, false}", "isFound:", isFound, "idx:", idx)


    // (true->false) find false: []bool{true, true, true, true, true, false, false, false, false} isFound: true idx: 5
    // (true->false) find true: []bool{true, true, true, true, true, false, false, false, false} isFound: true idx: 0
    // (true->false) find false: []bool{true, true, true, true, true, true, true, true, true} isFound: false idx: 9
    // (true->false) find true: []bool{true, true, true, true, true, true, true, true, true} isFound: true idx: 0
    // (true->false) find false: []bool{false, false, false, false, false, false, false, false, false} isFound: true idx: 0
    // (true->false) find true: []bool{false, false, false, false, false, false, false, false, false} isFound: false idx: 0
    // (false->true) find false: []bool{false, false, false, false, false, true, true, true, true} isFound: true idx: 0
    // (false->true) find true: []bool{false, false, false, false, false, true, true, true, true} isFound: true idx: 5
    // (false->true) find false: []bool{true, true, true, true, true, true, true, true, true} isFound: false idx: 0
    // (false->true) find true: []bool{true, true, true, true, true, true, true, true, true} isFound: true idx: 0
    // (false->true) find false: []bool{false, false, false, false, false, false, false, false, false} isFound: true idx: 0
    // (false->true) find true: []bool{false, false, false, false, false, false, false, false, false} isFound: false idx: 9
}
