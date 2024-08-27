package math_test

import (
	"log"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestPow 測試平方
func TestPow(t *testing.T) {
	expected := 8.0
	value := math.Pow(2, 3) // 這樣使用平方
	assert.Equal(t, expected, value, "2^3 should be 8")
	log.Println(value)
}

// TestPowMod 測試平方取餘數，可以處理大數而不會溢位
func TestPowMod(t *testing.T) {
	powMod := func(x, y, mod int) int { // x ^ y % 3
		// a*b mod c = (a mod c * b mod c) mod c
		res := 1 // 1 是因為 當 b 是 1時，不影響 a 的計算結果, 所以最適合做初值

		// 以 4^13 為例，把13拆成二進位 1101 => 2^3 + 2^2 + 2^0 = 8 + 4 + 1 = 13
		// 拆二進位的元因是我們想要對一次處理一個 x ，而次方每/2一次剛好就是右移一次
		//
		//       1       1      0       1         => y
		//   2^3*1 + 2^2*1+ 2^1*0 + 2^0*1 =   13  => y
		//  ^                      ^
		// 4                      4               => x
		for y > 0 {
			// 最後一位是1，要表示 x^1，因為要一個個處理掉，所以把x算到res裡，注意定理
			//  x*res % mod = (x % mod * res % mod) % mod
			// 所以可以取 x 的一個次方下來，乘到 res 中 不影響結果
			// 4^13*res   % 3
			// 4^12*res*4 % 3 = (4^12 % 3 * (res*4) % 3) % 3
			if y&1 == 1 {
				res = res * x % mod
			}

			// 接下來 因為變成了 1 1 0 0 了，我們可以透過右移的方式快速的計算 >> 完後是 1 1 0
			// x^y  等於 (x*x)^(y/2)
			//
			//   ^12             ^(6+6)  => y |           ^6
			// 4       可以看成是 4        => x |         ^2            ^6
			//   對y/2，對x做平方，值沒變 --------------> 4        => (4*4)
			y >>= 1
			x = x * x % mod

			// 解釋一下為什麼要多 % mod 主要原因是因為
			// 1. 防止 x 過大
			// 2. 照下面的推導，不會改變結果
			// a*b mod c = ((a mod c) * (b mod c)) mod c
			// 簡單推導，從a*b開始：
			// a = b = x
			// x^2 %z = (x%z * x%z) %z = (x%z)^2 %z
			// a*b 也以看成 x * x * x
			// x^3 %z = (x%z * x%z * x%z) %z = (x%z)^3 %z
			// 所以 x^y %z = (x%z)^y %z
			//
			// 從上面反推：
			// x^y %z  = ((x%z) * (x%z) ... (x^z) )%z  (y個x%z)
			//         = (x%z)^y %z
			// 所以 x^y %z = (x%z)^y %z
		}

		return res
	}

	log.Println("4^13 % 3 = ", powMod(4, 13, 3))                       // 4^13 % 3 = 1
	log.Println("17^375638465834 % 5 = ", powMod(17, 375638465834, 5)) // 17^375638465834 % 5 =  4

}
