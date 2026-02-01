[300. Longest Increasing Subsequence](https://leetcode.com/problems/longest-increasing-subsequence)  
爆力 $$O(n^2)$$
對於每個num，我們都會加上前面最長的字序列的長度。
```go
// normal
func lengthOfLIS(nums []int) int {
    // dp[i] = max(nums[j]+1), where j<i and nums[j]<nums[i]
    // O((n*(n-1))/2) -> (O(n^2)
    n := len(nums)
    dp := make([]int, n)
    for i:= range nums { dp[i]=1 }
    for i:= range nums {
        for j:= range i {
            if nums[j] < nums[i] {
                dp[i] = max(dp[i], dp[j]+1)
            }
        }
    }

    return slices.Max(dp)
}
```
觀察：  
- [5, 7], [5, 6] 當遇到 8 的時候，都可以使用，長度變為3。因此相同長度的子序列我們可以只保留尾數最小的。  
至此，我們可以在每個不同長度下只維護一組子序列  
- 當新來的一個num，他可以優化某組子序列，進而讓我們可以只保存每個長度下的最優解
- [2,5,3,7,10,6,18]  
input num:  2
0 : [2]
input num:  5
0 : [2]
1 : [2 5]
input num:  3
0 : [2]
1 : [2 3]
input num:  7
0 : [2]
1 : [2 3]
2 : [2 3 7]
input num:  10
0 : [2]
1 : [2 3]
2 : [2 3 7]
3 : [2 3 7 10]
input num:  6
0 : [2]
1 : [2 3]
2 : [2 3 6]
3 : [2 3 7 10]
input num:  18
0 : [2]
1 : [2 3]
2 : [2 3 6]
3 : [2 3 7 10]
4 : [2 3 7 10 18]
維護長度各不同的arr，走向greedy $$O(n^2)$$ 但快於爆力  
```go
func lengthOfLIS(nums []int) int {
    // 維護不同長度的最長子序列
    arrs := make([][]int, 0)
    for _, num := range nums {
        isOptimized := false
        for _, arr := range arrs {
            top := arr[len(arr)-1]
            if top >= num { // 可以優化某組子序列
                isOptimized = true
                arr[len(arr)-1] = num
                break
            }
        }

        // 不能優化表示得到了一組更長的子序列
        if !isOptimized {
            if len(arrs)>0 {
                arr := arrs[len(arrs)-1]
                tmp := append([]int{}, arr...)
                tmp = append(tmp, num)
                arrs = append(arrs, tmp)
            } else {
                arrs = append(arrs, []int{num})
            }
        }

        // for演示debug
        fmt.Println("input num: ", num)
        for i, arr := range arrs {
            fmt.Println(i, ":", arr)
        }
    }

    ans := 0
    for _, arr := range arrs {
        ans = max(ans, len(arr))
    }

    return ans
}


```

其實只需要觀注最後一個數字，並因此只需要維護arrs的最後一個數字  
(greedy $$O(n^2)$$)  
```go
func lengthOfLIS(nums []int) int {
    // 過渡，不需要arrs，只需要維護長度是1, 2, 3 ...
    tails := make([]int, 0)
    for _, num := range nums {
        isFound := false
        for i, top := range tails {
            if top >= num {
                tails[i] = num
                isFound = true
                break
            }
        }

        if !isFound { // 有找到 -> 不需當起點
            tails = append(tails, num)
        }
    }

    return len(tails)
}
```

觀察到這個演算法的 tails 只會個遞增的陣列，所以可以用 binary search 來找到直接要替換的 tail。 $$O(nlogn)$$  
最終 greedy + binary search
```go
func lengthOfLIS(nums []int) int {
    // 用二分法快速的定位tail，使n -> logn
    tails := make([]int, 0)
    for _, num := range nums {
        l, r := -1, len(tails)
        for l+1<r {
            mid := l + (r-l)/2
            tail := tails[mid] // top
            // 小 -> 大
            // > > > '<' < <
            if num <= tail {
                r = mid
            } else {
                l = mid
            }
        }

        if r == len(tails) {
            tails = append(tails, num)
        } else {
            tails[r] = num
        }
    }

    return len(tails)
}

```
