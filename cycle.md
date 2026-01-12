
給一個array，頭尾行成cycle，右移k位直到回到原點
```go
        // i 表示第一個
        carry := nums[i] // 先把第一個拿起來
        next := i
        for {
            next = (next+k)%n // 算出要放到哪
            nums[next], carry = carry, nums[next] // 放進去，並拿出他的值

            if next == i { break } // 最後放回一開始的位置時結事
        }
```

- 如果你想要每個元素都右移k位，有個關鍵的因素要考慮
  - 如果 gcd(len(nums), k) != 1 那就會有提前在某個cycle裡loop的情況(因為有gcd(len(nums), k)個cycle)，不會每一個元素都走到 // 因為 0 -> 0+k -> 0+2k ... -> 0，有最大公因素的話有些元素會被跳過
    - 所以就變成前gcd(len(nums), k) 個要都要cycle一次
  - 如果 gcd(len(nums), k) == 1 那就不會有問題，，走一遍就可以全部走完 
