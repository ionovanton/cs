package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Trap(t *testing.T) {
	type test struct {
		input    []int
		expected int
	}

	tests := []test{
		{
			input:    []int{1, 2, 1, 3, 2, 1, 2, 4, 3, 2, 3, 2},
			expected: 6,
		},
		{
			input:    []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			expected: 6,
		},
		{
			input:    []int{4, 2, 0, 3, 2, 5},
			expected: 9,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			got := trap(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func trap(height []int) int {
	level := 0
	absoluteMax := 0
	tower := 0
	for i := 0; i < len(height); i++ {
		if absoluteMax < height[i] {
			absoluteMax = height[i]
			tower = i
		}
	}

	mx := 0
	for i := 0; i < tower; i++ {
		mx = max(mx, height[i])
		level += mx - height[i]
	}

	mx = 0
	for i := len(height) - 1; i > tower; i-- {
		mx = max(mx, height[i])
		level += mx - height[i]
	}

	return level
}
