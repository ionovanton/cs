package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_BestTimeToBuyAndSellStock2(t *testing.T) {
	type test struct {
		prices []int
		want   int
	}

	tests := []test{
		{
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   7,
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
			want:   7,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			got := maxProfit2(tt.prices)
			assert.Equal(t, tt.want, got)
		})
	}
}

func maxProfit2(prices []int) int {
	sums := make([]int, 0, len(prices))
	for i := 0; i < len(prices)-1; i++ {
		sums = append(sums, prices[i+1]-prices[i])
	}
	total := 0
	for _, sum := range sums {
		if sum > 0 {
			total += sum
		}
	}
	return total
}
