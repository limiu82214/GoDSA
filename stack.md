stack 
- 通常()之類的對應容易用到
- 除了單純的值之外，也可以存idx，也可以在stack中放復合類型的值，來存更多資訊
- 如果看起來有Recursion，很可能要把pop後處理完的東西再放回stack，這樣後序處理就會處理這些字串
- 從stack拿出的字符會與原本的順序相反，所以有時會想反轉回來，可以對拿出來的資料做 slices.Reverse()
```go
                // take [ ... ]
                tmp := make([]byte, 0)
                for stack[len(stack)-1] != '[' {
                    tmp = append(tmp, stack[len(stack)-1])
                    stack = stack[:len(stack)-1]
                }
                stack = stack[:len(stack)-1] // pop '['
                slices.Reverse(tmp)
                str := string(tmp)

                // take number
                tmp = tmp[:0]
                for len(stack)>0 && (stack[len(stack)-1]>='0' && stack[len(stack)-1] <= '9') {
                    tmp = append(tmp, stack[len(stack)-1])
                    stack = stack[:len(stack)-1]
                }
                slices.Reverse(tmp)
                number, _ := strconv.Atoi(string(tmp))

```
