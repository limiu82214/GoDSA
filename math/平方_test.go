package math_test

import (
	"log"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestPow 測試平方
func TestPow(t *testing.T) {
	expected := 8.0
	value := math.Pow(2, 3) // 這樣使用平方
	assert.Equal(t, expected, value, "2^3 should be 8")
	log.Println(value)
}
