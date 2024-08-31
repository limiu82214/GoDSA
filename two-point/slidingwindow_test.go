package towpoint

import (
	"log"
	"testing"
)

func TestSlidingWindow(t *testing.T) {
	// 基本上同TwoPointer

	nums := []int{1, 2, 3, 4, 5, 6}
	// 不固定大小的
	l, r := 0, 0
	for r < len(nums) {
		_ = nums[r] // 對 R 做點事情，例如加到window或處理ans
		// 處理window
		if true { // 達成目標
			break
		}

		for { // 在window還是符合你設定的規則時，對 L 做點事情，例如移除window或處理ans
			_ = nums[l]
			// 處理window // 這裡不要忘了
			l++
		}
		r++ // 每次只走R一步
	}

	log.Println("固定window大小的")
	windowSize := 3
	window := nums[0:windowSize]
	// 先依題目算出 window 的內容 => ans
	// do something with window

	for i := windowSize; i < len(nums); i++ {
		l, r := i-windowSize, i

		// 在移除左邊的值時做些事情
		window = window[1:]
		// 在增加右邊的值時做些事情
		window = append(window, nums[r])

		// 更新 ans 如果有必要

		// 因為index 0-2 會在for 之前先算完，所以這裡只會印出後面的
		log.Println("l: ", l, "r: ", r)
		log.Println("fix window: ", window)
	}
}
