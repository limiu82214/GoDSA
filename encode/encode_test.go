package encode_test

import "testing"

func TestTwoDimensional2OneDimensional(t *testing.T) {
	twoDimensional2OneDimensional := func(arr map[int]map[int]string) map[int]string {
		oneDimensional := make(map[int]string, 0)
		// arr is a two dimensional array
		// 假設 arr[i][j], i < 10^5, j < 10^5

		// 放一個 0,0 的值
		// i, j := 0, 0
		// oneDimensional := make(map[int]string, 0)
		// oneDimensional[i*100000+j] = arr[i][j]

		// 把整個二維陣列轉成一維陣列
		for i := range arr {
			for j := range arr[i] {
				oneDimensional[i*100000+j] = arr[i][j]
			}
		}
		// 因為 i, j 都小於 10^5，所以可以用 i*10^5+j 來表示一維陣列的 index
		// 這樣就有i 與j 的組合，不會重複，用一維陣列來表示二維陣列
		return oneDimensional
	}

	m := map[int]map[int]string{
		0: {
			0: "01", // m[0][0] = "01"
			1: "01", // m[0][1] = "01"
		},
		1: {
			0: "11", // m[1][0] = "11", .....100000 => m[1][0]
		},
		2: {
			200: "12", // m[2][200] = "12" .....200200 => m[2][200]
		},
	}

	// xxxxxyyyyy => m[i(xxxxx)][j(yyyyy)]
	t.Log(twoDimensional2OneDimensional(m)) //map[0:01 1:01 100000:11 200200:12]
}
