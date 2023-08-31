package sort_test

func binarySearch(arr []int, target int) int {
	l, r := 0, len(arr)-1
	for l <= r { // 相同index的話也不能跳過
		mid := l + (r-l)/2 // 防止 l + r 溢位
		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			l = mid + 1 // +1 表示不保留 mid，應該依照題目決定要不要保留 mid，也可以保留在某個地方再捨棄
		} else {
			r = mid - 1 // -1 表示不保留 mid，應該依照題目決定要不要保留 mid，也可以保留在某個地方再捨棄
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

func findBoundary(arr []bool) int {
	l, r := 0, len(arr)-1
	firstTrue := -1
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] == true {
			firstTrue = mid
			r = mid - 1
		} else {
			// firstFalse = mid
			l = mid + 1
		}
	}

	return firstTrue
}

