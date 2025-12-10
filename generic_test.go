package godsa_test

import (
	"slices"
	"testing"
)

// TestMapInit 這是一個簡單的Map初始化
func TestGeneric(t *testing.T) {
    // MOD with int
    const MOD = 1_000_000_007 // 10^9 + 7
    t.Log("MOD: ", MOD)

    // Slice
    s := []int{1, 2, 3, 4, 5}
    t.Log("Min of slice:", slices.Min(s)) // 1
}
