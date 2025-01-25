package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Rotate(t *testing.T) {
	type test struct {
		nums []int
		k    int
		want []int
	}

	tests := []test{
		{
			nums: []int{1, 2, 3, 4, 5, 6, 7},
			k:    3,
			want: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			nums: []int{-1, -100, 3, 99},
			k:    2,
			want: []int{3, 99, -1, -100},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			rotate(tt.nums, tt.k)
			assert.Equal(t, tt.want, tt.nums)
		})
	}
}

func rotate(nums []int, k int) {
	if len(nums) == k || len(nums) == 1 {
		return
	}
	a := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		a[(i+k)%len(nums)] = nums[i]
	}
	copy(nums, a)
}
