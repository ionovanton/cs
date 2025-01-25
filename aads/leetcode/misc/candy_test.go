package test

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Candy(t *testing.T) {
	type test struct {
		raring []int
		want   int
	}

	tests := []test{
		{
			raring: []int{1, 0, 2},
			want:   5,
		},
		{
			raring: []int{1, 2, 2},
			want:   4,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			got := candy(tt.raring)
			require.Equal(t, tt.want, got)
		})
	}
}

func candy(ratings []int) int {
	candies := make([]int, len(ratings))
	for i := 0; i < len(candies); i++ {
		candies[i] = 1
	}
	sum := 0
	for _, c := range candies {
		sum += c
	}
	return sum
}
