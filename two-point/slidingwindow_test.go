package towpoint

import "testing"

func TestSlidingWindow(t *testing.T) {
	// 基本上同TwoPointer

	// // 先移動R，直到滿足window條件
	// long := 3 // 這也可以是某個條件的判斷

	// // 如果L與R的移動方式相同，可以只算其中一個，然後推算出另一個
	// for i := 5; i < 200; i++ {
	// 	// 如果能在這個for中一次L或R只動一個step，另一個再依照對方更動，會更容易
	// 	l, r := i-long, i
	// 	// ...
	// 	_, _ = l, r
	// }

	nums := []int{1, 2, 3, 4, 5, 6}
	// 不固定大小的
	l, r := 0, 0
	for {
		_ = nums[r] // 對 R 做點推事情
		// 處理window
		if true { // 達成目標
			break
		}

		for { // 對 L 做點事情
			_ = nums[l]
			// 處理window // 這裡不要忘了
			l++
		}
		r++ // 每次只走R一步
	}

	// 這是固定window大小的
windowSize := 3
for i := windowSize; i < len(nums); i++ {
	l, r := i-windowSize, i
	window := nums[l:r]
	// k=3
	// r =   0   1   2   3   4
	// l =  -3  -2  -1   0   1

	// remove[l]
	window = window[1:]
	// append[r]
	window = append(window, nums[r])

	// here is complete window
	if r+1 >= windowSize {
		// 此刻 l = -1, r = 2 . window = [1, 2, 3]
	}
}
}
