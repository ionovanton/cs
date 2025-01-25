package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_(t *testing.T) {
	type test struct {
		gas      []int
		cost     []int
		expected int
	}

	tests := []test{
		{
			gas:      []int{1, 2, 3, 4, 5},
			cost:     []int{3, 4, 5, 1, 2},
			expected: 3,
		},
		{
			gas:      []int{2, 3, 4},
			cost:     []int{3, 4, 3},
			expected: -1,
		},
		{
			gas:      []int{5, 1, 2, 3, 4},
			cost:     []int{4, 4, 1, 5, 1},
			expected: 4,
		},
		{
			gas:      []int{3, 1, 1},
			cost:     []int{1, 2, 2},
			expected: 0,
		},
		{
			gas:      []int{6, 1, 4, 3, 5},
			cost:     []int{3, 8, 2, 4, 2},
			expected: 2,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			got := canCompleteCircuit(tt.gas, tt.cost)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func canCompleteCircuit(gas []int, cost []int) int {
	currentGas, startStation := 0, 0
	totalGas, totalCost := 0, 0
	for i := 0; i < len(gas); i++ {
		totalGas += gas[i]
		totalCost += cost[i]
		currentGas += gas[i] - cost[i]

		// Из этой позиции не можем продолжить движение.
		// Пробуем продолжить движение со следующей.
		if currentGas < 0 {
			startStation = i + 1
			currentGas = 0
		}
	}
	// Нет решений, если всего топлива меньше, чем нужно для поездки
	if totalGas < totalCost {
		return -1
	}
	return startStation
}
