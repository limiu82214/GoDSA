package loop

func NPlus1() {
	nums := []int{1, 2, 3, 4, 5}
	k := 3
	for i := range nums {
		// 會錯，因為i+1會超出範圍
		// if nums[i] < k && nums[i+1] > k {
		// 	break
		// }

		if nums[i] < k &&
			(i+1 == len(nums) || nums[i+1] > k) {
			// 當i+1超出範圍時，就不需判斷i+1的條件
		}
	}
}
