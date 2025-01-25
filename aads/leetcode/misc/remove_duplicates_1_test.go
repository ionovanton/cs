package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_RemoveDuplicates1(t *testing.T) {
	type test struct {
		nums  []int
		wantK int
		want  []int
	}

	tests := []test{
		{
			nums:  []int{1, 1, 2},
			wantK: 2,
			want:  []int{1, 2},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			gotK := removeDuplicates1(tt.nums)
			FirstKthElementsMatch(t, tt.want, tt.nums, tt.wantK)
			assert.Equal(t, tt.wantK, gotK)
		})
	}
}

func removeDuplicates1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dups := make([]int, 0, len(nums))
	for i := 0; i < len(nums); {
		dup := nums[i]
		dups = append(dups, dup)
		for i < len(nums) && nums[i] == dup {
			i++
		}
	}
	copy(nums, dups)
	return len(dups)
}
