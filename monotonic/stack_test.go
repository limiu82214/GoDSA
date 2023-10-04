package monotonic_test

func InsertIntoStack(maxH []int) {
	mstk := []int{} // 大->小

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

}
