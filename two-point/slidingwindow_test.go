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
    	l, r := 0, 0
    	for r < len(s2) {
		// R 進入窗口
		window = append(window, nums[r])
		r++
		// [l, r)
	
		// 長度足夠
		if r-l >= size {
	
		    // 判斷答案, 更新ans 如果有必要
	
		    // L 離開窗口
		    window = window[1:]
		    l++
		}
	}
}
