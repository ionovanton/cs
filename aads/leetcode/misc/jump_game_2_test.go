package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_JumpGame2(t *testing.T) {
	type test struct {
		nums     []int
		expected int
	}

	tests := []test{
		{
			nums:     []int{1, 100, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			expected: 2,
		},
		{
			nums:     []int{0},
			expected: 0,
		},
		{
			nums:     []int{1},
			expected: 0,
		},
		{
			nums:     []int{100, 0},
			expected: 1,
		},
		{
			nums:     []int{2, 0, 0},
			expected: 1,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			got := jump(tt.nums)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func jump(nums []int) int {
	result, l, r := 0, 0, 0

	for r < len(nums)-1 {
		farthest := 0
		for i := l; i < r+1; i++ {
			farthest = max(farthest, i+nums[i])
		}
		l = r + 1
		r = farthest
		result++
	}
	return result
}
