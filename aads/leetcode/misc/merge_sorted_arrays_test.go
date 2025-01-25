package test

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_MergeSortedArray(t *testing.T) {
	type test struct {
		name string
		a    []int
		b    []int
		m    int
		n    int
		want []int
	}

	tests := []test{
		{
			a:    []int{1, 2, 3, 0, 0, 0},
			m:    3,
			b:    []int{2, 5, 6},
			n:    3,
			want: []int{1, 2, 2, 3, 5, 6},
		},
		{
			a:    []int{0},
			m:    0,
			b:    []int{1},
			n:    1,
			want: []int{1},
		},
		{
			a:    []int{0},
			m:    0,
			b:    []int{},
			n:    0,
			want: []int{0},
		},
		{
			a:    []int{1, 0},
			m:    1,
			b:    []int{2},
			n:    1,
			want: []int{1, 2},
		},
		{
			a:    []int{1},
			m:    1,
			b:    []int{},
			n:    0,
			want: []int{1},
		},
		{
			a:    []int{4, 5, 6, 0, 0, 0},
			m:    3,
			b:    []int{1, 2, 3},
			n:    3,
			want: []int{1, 2, 3, 4, 5, 6},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test_%d", i+1), func(t *testing.T) {
			merge(tt.a, tt.m, tt.b, tt.n)
			require.Equal(t, tt.want, tt.a)
		})

	}
}

func merge(a []int, m int, b []int, n int) {
	last := m + n - 1
	for ; m > 0 && n > 0; last-- {
		if a[m-1] > b[n-1] {
			a[last] = a[m-1]
			m--
		} else {
			a[last] = b[n-1]
			n--
		}
	}
	for n > 0 {
		a[last] = b[n-1]
		n--
		last--
	}
}
