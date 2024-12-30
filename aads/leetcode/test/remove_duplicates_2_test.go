package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_RemoveDuplicates2(t *testing.T) {
	type test struct {
		nums  []int
		wantK int
		want  []int
	}

	tests := []test{
		{
			nums:  []int{1, 1, 2},
			wantK: 3,
			want:  []int{1, 1, 2},
		},
		{
			nums:  []int{1, 1, 1, 2, 2, 3},
			wantK: 5,
			want:  []int{1, 1, 2, 2, 3},
		},
		{
			nums:  []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			wantK: 7,
			want:  []int{0, 0, 1, 1, 2, 3, 3},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			gotK := removeDuplicates2(tt.nums)
			FirstKthElementsMatch(t, tt.want, tt.nums, tt.wantK)
			assert.Equal(t, tt.wantK, gotK)
		})
	}
}

func removeDuplicates2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dups := make([]int, 0, len(nums))
	for i := 0; i < len(nums); {
		dup := nums[i]
		dups = append(dups, dup)
		k := 0
		for i < len(nums) && nums[i] == dup {
			k++
			i++
		}
		if k >= 2 {
			dups = append(dups, dup)
		}
	}
	copy(nums, dups)
	return len(dups)
}
