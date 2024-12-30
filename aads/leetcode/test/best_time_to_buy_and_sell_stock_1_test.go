package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test_BestTimeToBuyAndSellStock1(t *testing.T) {
	type test struct {
		prices []int
		want   int
	}

	tests := []test{
		{
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   5,
		},
		{
			prices: []int{7, 6, 4, 3, 1},
			want:   0,
		},
		{
			prices: []int{1, 2},
			want:   1,
		},
		{
			prices: []int{2, 4, 1},
			want:   2,
		},
		{
			prices: []int{3, 2, 6, 5, 0, 3},
			want:   4,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			got := maxProfit1(tt.prices)
			assert.Equal(t, tt.want, got)
		})
	}
}

func maxProfit1(prices []int) int {
	min, k := math.MaxInt, 0
	maxDiff := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < min {
			min = prices[i]
			k = i
			for i := k + 1; i < len(prices); i++ {
				currentDiff := prices[i] - prices[k]
				if maxDiff <= currentDiff {
					maxDiff = currentDiff
				}
			}
		}
	}

	return maxDiff
}
