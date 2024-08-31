package monotonic_test

import (
	"log"
	"testing"
)

// TestMonotonicStackSmall2Big 測試單調堆疊
func TestMonotonicStackSmall2Big(t *testing.T) {
	maxH := []int{2, 1, 5, 6, 2, 3}
	mstk := InsertIntoStack(maxH)
	log.Println(maxH) // [2 1 5 6 2 3]
	log.Println(mstk) // [1 4 5]
}

// InsertIntoStack 將maxH的值依序放入stack
// MonotonicStack 的一個重要的觀點是，在stack變化的過程。
// 例如長條圖的值，在代入MonotonicStack時，若遇到一個很小的值，在Pop的時候，可以得知這個值的左右邊界
func InsertIntoStack(maxH []int) (mstk []int) {
	mstk = []int{} // (底)小->大(頂) , 存的是 maxH 的 index

	for i := 0; i < len(maxH); i++ {
		for len(mstk) > 0 && maxH[i] < maxH[mstk[len(mstk)-1]] {
			// 在stack變小前做點事情
			mstk = mstk[:len(mstk)-1]
			// 在stack變小後做點事情
		}

		// 整理好stack之後，還沒放入maxH前
		mstk = append(mstk, i)
		// 整理好stack之後，還沒放入maxH後
	}

	return mstk
}
