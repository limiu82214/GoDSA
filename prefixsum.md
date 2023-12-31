# Prefix Sum
* 一個關於陣例的概念
* 可以將一般陣列轉換成Prefix Sum陣列
* 其中每一個元素表示原本陣到0-i的總合
    * 例如[1,2,3,4] => [1,3,6,10]
* 通常會在Prefix Sum最前面再加一個0元素，這表示空 => [0,1,3,6,10]
* 有用的特性
    * 某個子陣列的合會等於PrefixSum陣列相減
    * 例如：[1,2,3,4] => [0,1,3,6,10]
        某子陣列 [2,3,4](arr[1][3]) 之合 => 9
        = PrefixSum[0][3]
    * 可以整理成 arr[i][j] = ps[j] - ps[i-1]
    * 這可以讓原本必需走遍所有組合的情況變成只需要走過一次，也就是O(n^2)=>O(n)
    * 注意通常這會用在與Sum有關、然後需要取出或是解析什麼東西的時候。
