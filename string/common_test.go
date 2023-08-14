package string_test

import (
	"regexp"
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
