package math_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPost(t *testing.T) {
	expected := 8.0
	value := math.Pow(2, 3) // 這樣使用平方
	assert.Equal(t, expected, value, "2^3 should be 8")
}
