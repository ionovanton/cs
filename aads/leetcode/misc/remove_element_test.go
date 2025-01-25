package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_RemoveElement(t *testing.T) {
	type test struct {
		a     []int
		val   int
		want  []int
		wantK int
	}

	tests := []test{
		{
			a:     []int{3, 2, 2, 3},
			val:   3,
			want:  []int{2, 2},
			wantK: 2,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			k := removeElement(tt.a, tt.val)
			assert.Equal(t, tt.want, tt.a)
			assert.Equal(t, tt.wantK, k)
		})
	}
}

func removeElement(nums []int, val int) int {
	k := 0
	result := make([]int, 0, k)
	for _, vv := range nums {
		if vv != val {
			result = append(result, vv)
		}
	}

	copy(nums, result)

	return len(result)
}

//b = append(b[:i], b[i+1:]...)
