package encode

func TwoDimensional2OneDimensional(arr map[int]map[int]interface{}) {
	// arr is a two dimensional array
	// 假設 arr[i][j], i < 10^5, j < 10^5
	i, j := 0, 0
	oneDimensional := make(map[int]interface{}, 0)
	oneDimensional[i*100000+j] = arr[i][j]
	// 因為 i, j 都小於 10^5，所以可以用 i*10^5+j 來表示一維陣列的 index
	// 這樣就有i 與j 的組合，不會重複，用一維陣列來表示二維陣列
}
