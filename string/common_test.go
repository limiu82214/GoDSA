package string_test

import (
	"log"
	"regexp"
	"sort"
	"strings"
	"testing"
)

// TestToLower 測試轉換為小寫
func TestToLower(t *testing.T) {
	s := "12984712893 jkD 2398DJKsjhf "
	out := strings.ToLower(s)
	log.Println(out) // 12984712893 jkd 2398djksjhf
}

// TestRegexpToRemoveSomething 測試regex
func TestRegexpToRemoveSomething(t *testing.T) {
	s := "12984712893 jkd 2398djksjhf "
	str := strings.ToLower(s)
	str = regexp.MustCompile(`[^a-z0-9]`).ReplaceAllString(str, "")
	log.Println(str) // 12984712893jkd2398djksjhf
}


// TestChangeString 過濾字串
// 直接取string的某個字元是byte
// forrange string是rune
// rune和byte若字是ascii，則是一樣的
func TestChangeString(t *testing.T) {
	s := "12984712893 jkd 2398djksjhf "
	runeS := []rune{}
	for _, r := range s {
		if (r >= rune('0') && r <= rune('9')) || (r >= rune('a') && r <= rune('z')) {
			runeS = append(runeS, r)
		}
	}

	outStr := string(runeS)
	log.Println(outStr) // 12984712893jkd2398djksjhf
}

// TestRuneSort 排序rune
func TestRuneSort(t *testing.T) {
	runes := []rune{'我', '是', '程', '式', '設', '計', '師'}
	// note rune is unicode and its order is same as ascii
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	log.Println(string(runes)) // 師式我是程計設
}

// TestCopySlice 測試複製切片
func TestCopySlice(t *testing.T) {
	original := []int{3, 1, 4, 1, 5, 9, 2, 6, 5}
	// 複製原始切片
	copied := make([]int, len(original))
	copy(copied, original)
	log.Println("original:", original) // original: [3 1 4 1 5 9 2 6 5]
	log.Println("copied:", copied)     // copied: [3 1 4 1 5 9 2 6 5]
}

// TestHasPrefix 測試是否有前綴
func TestHasPrefix(t *testing.T) {
	// func HasPrefix(str, pre string) bool
	b := strings.HasPrefix("Gopher", "Go")
	log.Println(b) // true

}
