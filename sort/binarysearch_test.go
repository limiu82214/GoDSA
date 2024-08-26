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
			// firstIdx = mid  // first False
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
