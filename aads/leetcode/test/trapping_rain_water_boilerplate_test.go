package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Trap(t *testing.T) {
	type test struct {
		input    []int
		expected int
	}

	tests := []test{
		{
			input:    []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			expected: 9,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			got := trap(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

/*

- проход считается пройденным если нашли лужу
- новый проход начинается с места прошлой лужи
- цикл идет пока оба указателя не откнутся в конец

*/

func trap(height []int) int {
	for i := 0; i < len(height); i++ {
		for j := i; j < len(height) && height[j] >= height[i]; j++ {
			fmt.Println(height[i], height[j])
		}
	}
	return 0
}
