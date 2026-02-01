[300. Longest Increasing Subsequence](https://leetcode.com/problems/longest-increasing-subsequence)  
爆力 $$O(n^2)$$
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

維護長度各不同的stack，走向greedy $$O(n^2)$$ 但快於爆力
```go
// focus on stacks length
func lengthOfLIS(nums []int) int {
    // 過渡，維護長度逐建增加的stack，只替換第一個符合的
    stacks := make([][]int, 0)
    for _, num := range nums {
        isFound := false
        for _, stack := range stacks {
            top := stack[len(stack)-1]
            if top >= num {
                stack[len(stack)-1] = num
                isFound = true
                break
            }
        }

        if !isFound { // 有找到 -> 不需當起點
            if len(stacks)>0 {
                last := stacks[len(stacks)-1]
                tmp := make([]int, len(last))
                copy(tmp, last)
                tmp = append(tmp, num)
                stacks = append(stacks, tmp)
            } else {
                stacks = append(stacks, []int{num})
            }
        }
    }

    ans := 0
    for _, stack := range stacks {
        ans = max(ans, len(stack))
    }

    return ans
}
```

其實只需要注最後一個數字 (greedy $$O(n^2)$$)
```go
func lengthOfLIS(nums []int) int {
    // 過渡，不需要stacks，只需要維護長度是1, 2, 3 ...
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
