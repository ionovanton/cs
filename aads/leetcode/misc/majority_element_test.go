package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MajorityElement(t *testing.T) {
	type test struct {
		nums []int
		want int
	}

	tests := []test{
		{
			nums: []int{3, 2, 3},
			want: 3,
		},
		{
			nums: []int{2, 2, 1, 1, 1, 2, 2},
			want: 2,
		},
		{
			nums: []int{6, 6, 6, 7, 7},
			want: 6,
		},
		{
			nums: []int{8, 8, 7, 7, 7},
			want: 7,
		},
		{
			nums: []int{6, 5, 5},
			want: 5,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			got := majorityElement(tt.nums)
			assert.Equal(t, tt.want, got)
		})
	}
}

func majorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	m := make(map[int]int, len(nums))
	for _, num := range nums {
		m[num]++
	}
	maxK, maxMajority := 0, len(nums)/2
	for k, appeared := range m {
		if appeared > maxMajority {
			maxMajority = appeared
			maxK = k
		}
	}
	return maxK
}
