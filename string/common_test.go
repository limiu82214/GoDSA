package string_test

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"testing"
)

func TestToLower(t *testing.T) {
	s := "12984712893 jkD 2398DJKsjhf "
	out := strings.ToLower(s)
	_ = out
}

func TestRegexpToRemoveSomething(t *testing.T) {
	s := "12984712893 jkd 2398djksjhf "
	str := strings.ToLower(s)
	str = regexp.MustCompile(`[^a-z0-9]`).ReplaceAllString(str, "")
	// fmt.Println(str)
}

func TestChangeString(t *testing.T) {
	s := "12984712893 jkd 2398djksjhf "
	runeS := []rune{}
	for _, r := range s {
		if (r >= rune('0') && r <= rune('9')) || (r >= rune('a') && r <= rune('z')) {
			runeS = append(runeS, r)
		}
	}

	outStr := string(runeS)
	// fmt.Println(outStr)
	_ = outStr
}

func TestRuneSort(t *testing.T) {
	runes := []rune{'我', '是', '程', '式', '設', '計', '師'}
	// note rune is unicode and its order is same as ascii
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	fmt.Println(string(runes))
}

func TestCopySlice(t *testing.T) {
	original := []int{3, 1, 4, 1, 5, 9, 2, 6, 5}
	// 複製原始切片
	copied := make([]int, len(original))
	copy(copied, original)
}

// 直接取string的某個字元是byte
// forrange string是rune
