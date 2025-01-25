package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ProductOfArrayExceptSelf(t *testing.T) {
	type test struct {
		nums     []int
		expected []int
	}

	tests := []test{
		{
			nums:     []int{1, 2, 3, 4},
			expected: []int{24, 12, 8, 6},
		},
		{
			nums:     []int{-1, 1, 0, -3, 3},
			expected: []int{0, 0, 9, 0, 0},
		},
		{
			nums:     []int{0, 0, -3, 3},
			expected: []int{0, 0, 0, 0},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			got := productExceptSelf(tt.nums)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func productExceptSelf(nums []int) []int {
	product := 1
	zeroCounts := 0
	productWithoutZero := 1
	for _, num := range nums {
		if num == 0 {
			zeroCounts++
		} else {
			productWithoutZero *= num
		}
		product *= num
	}

	if zeroCounts != 1 {
		for i := range nums {
			if nums[i] == 0 {
				nums[i] = 0
			} else {
				nums[i] = product / nums[i]
			}
		}
	} else {
		for i := range nums {
			if nums[i] == 0 {
				nums[i] = productWithoutZero
			} else {
				nums[i] = product / nums[i]
			}
		}
	}

	return nums
}
