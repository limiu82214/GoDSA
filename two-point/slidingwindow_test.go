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

	log.Println("萬金油版本")
	l, r = 0, 0
	for {
		// X 非法狀態，此時 L 和 R 在同一個位置上
		// # R 前進
		// for r<len(nums) && R前進的條件 { w.Add(R); r++} // 若 R 不是每次走一步，可以依條件往前
		window = append(window, nums[r])
		// [L, R] 閉區間
		r++
		
		// # 判斷答案 & 出界
		// [L, R) 半開區間
		// if r-l >= k { ... 判斷答案 } // 半開區間 R-l = 長度
		if r == len(nums) { break } // 半開區間的R等於nums長度時，就歷遍完了

		// # L 前進
		// for l<len(nums) && L前進的條件 { w.Del(L); l++} // 若 L不是每次走一步，可以依條件往前
		// if  r-l >=k { 等長度到達 k 才開始走L ，這等於變形成滑動窗口
		window = window[1:]
		// (L, R) 開區間
		l++
		// }
		// X 非法狀態，此時 L 和 R 在同一個位置上
	}

	log.Println("固定window大小的")
    	windowSize := 3
    	l, r = 0, 0
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
