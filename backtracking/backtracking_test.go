package backtracking_test

// 先畫出樹狀圖，再實作
// leaf是解
// edge是往解的路線
// 通常會有一個startIdx，當他每次預只深1層的時候就表示樹的深度
//   也可以是處理原始數據的索引，這取決於你的樹是如何建立的
// 可選->additionalStates，代表額外的狀態
// 看一下是要返回什麼值，再決定是用哪種模式
// 1. 返回一個值
// 2. 返回多個值
// 時間複雜度是狀態空間樹中的節點數乘以每個節點的操作

// 適合於需要計算某種累計結果的問題，例如：
// 計算滿足某些條件的路徑數量。
// 找到最短/最長的路徑長度。
// 計算所有可能路徑的總和。
// startIdx 通常意謂著樹的深度/原資料處理的進度索引
func dfs(startIdx int, additionalStates ...interface{}) int {
	if isLeaf(startIdx) {
		return 1
	}
	ans := 0 // Assuming initial value is 0
	for _, edge := range getEdges(startIdx, additionalStates...) {
		// If you have additional states to update
		// update(additionalStates...)
		ans = aggregate(ans, dfs(startIdx+len(edge.([]interface{})), additionalStates...)) // 處理了資料的長度，會加到startIdx，這樣是防止重覆的處理相同的資料
		// If you have additional states to revert
		// revert(additionalStates...)
	}
	return ans
}

var ans []string // Using string as a placeholder type for path
// 不返回任何值，而是直接修改全局變量 ans 來收集結果。這種模式常見於需要搜集所有可能結果的問題，例如尋找所有可能的路徑或組合。
// 適合於需要找到所有滿足某些條件的具體解的問題，例如：
// 找到所有可能的路徑。
// 列出所有可能的組合或排列。
func dfs2(startIdx int, path string, additionalStates ...interface{}) {
	if isLeaf(startIdx) {
		// 如果path是[]string，要特別注意DFS會修改path的值，因此放入答案之前要先拷貝一份 OR 如下面一樣，代入的值直接是newPath也可以
		// tmpPath := make([]string, len(path))
		// copy(tmpPath, path)
		// ans = append(ans, tmpPath)

		ans = append(ans, path)
		return
	}
	for _, edge := range getEdges(startIdx, additionalStates...) {
		// Prune if needed
		// if !isValid(edge) {
		//     continue
		// }
		newPath := path + edge.(string)
		// If you have additional states to update
		// update(additionalStates...)
		dfs2(startIdx+len(edge.([]interface{})), newPath, additionalStates...)
		// If you have additional states to revert
		// revert(additionalStates...)
	}
}

// isLeaf 通常意謂著找到了一個解
func isLeaf(startIdx int) bool {
	// Implement your logic here
	return false
}

// getEdges 通常意謂著怎麼列出前往解的路線
func getEdges(startIdx int, additionalStates ...interface{}) []interface{} {
	// Implement your logic here
	return []interface{}{}
}

// aggregate 通常意謂著怎麼計算累計結果
func aggregate(a, b int) int {
	// Implement your aggregation logic here
	return a + b
}
