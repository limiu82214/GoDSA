package monotonic_test

// monotonic 基本上可以因為input的指定條件，將答案變成FFFFFTTTTT的monotonic形式
// 再量上 binary search 就可以找到答案
// 以下是兩種置換成monotonic的方式
// 1. 將input的array視為要找的答案，並依照題目的要求，將答案變成FFFFFTTTTT的monotonic形式
// 2. 依照題目的要求，設計假定的答案範圍範圍(推論答案一定會落在這個monotonic區間)，然後將輸入的arr與input
//    將`假設的答案範圍`視為要找的答案，並依照題目的要求，將答案變成FFFFFTTTTT的monotonic形式
//    經典的題目是newspapers
// 時間複雜度: 將會是 O(feasible * log n)

func findBoundary1(questData []any) int {
	// 在這個例子中我們假設答案範圍在arr的index中
	l, r := 0, len(questData)-1
	firstTrue := -1
	for l <= r {
		mid := l + (r-l)/2
		if feasible(questData, mid) {
			firstTrue = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return firstTrue
}

func findBoundary2(qestData []any) int {
	// 在這個例子中我們假設答案範圍在0~10000
	l, r := 0, 100000 // 可以透過一些手段先減少這個範圍
	firstTrue := -1
	for l <= r {
		mid := l + (r-l)/2
		if feasible(qestData, mid) {
			firstTrue = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return firstTrue
}

func feasible(arr []any, mid int) bool {
	return arr[mid] == true
}
