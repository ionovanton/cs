package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_JumpGame(t *testing.T) {
	type test struct {
		nums     []int
		expected bool
	}

	tests := []test{
		{
			nums:     []int{2, 3, 3, 1, 4},
			expected: true,
		},
		{
			nums:     []int{3, 2, 1, 0, 4},
			expected: false,
		},
		{
			nums:     []int{0},
			expected: true,
		},
		{
			nums:     []int{1},
			expected: true,
		},
		{
			nums:     []int{0, 0},
			expected: false,
		},
		{
			nums:     []int{0, 20},
			expected: false,
		},
		{
			nums:     []int{0, 2, 3},
			expected: false,
		},
		{
			nums:     []int{100, 0},
			expected: true,
		},
		{
			nums:     []int{2, 0, 0},
			expected: true,
		},
		{
			nums:     []int{1, 0, 1, 0},
			expected: false,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			got := canJump(tt.nums)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func canJump(nums []int) bool {
	goal := len(nums) - 1

	for i := len(nums) - 1; i > -1; i-- {
		if i+nums[i] >= goal {
			goal = i
		}
	}
	return goal == 0
}
