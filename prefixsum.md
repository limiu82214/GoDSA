# Prefix Sum
* 一個關於陣例的概念
* 可以將一般陣列轉換成Prefix Sum陣列
* 其中每一個元素表示原本陣到0-i的總合
    * 例如[1,2,3,4] => [1,3,6,10]
* 通常會在Prefix Sum最前面再加一個0元素，這表示空 => [0,1,3,6,10]
* 有用的特性
    * 某個子陣列的合會等於PrefixSum陣列相減
    * 例如：[1,2,3,4] => [0,1,3,6,10]
        子陣列 [2,3,4] 之合 => 9
        = 原index 1~3 -> PrefixSum[4]-PrefixSum[1] // [) 所以是1和4
    * 可以整理成 arr[i][j] = ps[j] - ps[i-1]
    * 這可以讓原本必需走遍所有組合的情況變成只需要走過一次，也就是O(n^2)=>O(n)
    * 注意通常這會用在與Sum有關、然後需要取出或是解析什麼東西的時候。

Golang
```go
    // preSum // 務實上都用這個
    // pre[i] => sum(num[:i]) // [0 ~ i)
    pre := make([]int, len(nums)+1) // +1 因為要讓第一個元素是0
    for i := range nums {pre[i+1] = pre[i] + nums[i]}
    for i:=0; i<n; i++ { pre[i+1] = pre[i] * nums[i] }

    // sufSum // 務實上只拿來看右邊的總合是多少，很少拿來計算
    // suf[i] => sum(num[i:]) // [i ~ n]
    suf := make([]int, len(nums)+1)
    for i := range nums {
        // i 0 1 2 3  // 4
        // x 3 2 1 0  // 4 
        x := len(nums)-i-1
        suf[x] = suf[x+1] + nums[x]
    }
    for i:=n-1; i>=0; i-- { suf[i] = pre[i+1] + nums[i] }

      0, 1, 2, 3  // idx
    [ 1, 2, 3, 4] // nums
    [ 0, 1, 3, 6,10] // pre // pre[i] = sum(nums[0 : i]) [) // 不含 i
    [10, 9, 7, 4, 0] // suf // suf[i] = sum(nums[i : n]) []  // 含 i
    若要拿 sum(idx 1, idx 2) = nums[1] + nums[2] = 2+3 // []
    pre 是 [) 所以 => pre[3] sum(nums[:3]) - pre[1] sum(nums[:1]) = 6 - 1 = 5
    suf 是 [] 所以 => suf[1] sum(nums[1:]) - suf[3] sum(nums[3:]) = 9 - 4 = 5
```
