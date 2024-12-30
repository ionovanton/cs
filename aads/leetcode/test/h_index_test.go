package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_hIndex(t *testing.T) {
	type test struct {
		citations []int
		expect    int
	}

	tests := []test{
		{
			citations: []int{3, 0, 6, 1, 5},
			expect:    3,
		},
		{
			citations: []int{1, 3, 1},
			expect:    1,
		},
		{
			citations: []int{1},
			expect:    1,
		},
		{
			citations: []int{0},
			expect:    0,
		},
		{
			citations: []int{2, 1},
			expect:    1,
		},
		{
			citations: []int{1, 2},
			expect:    1,
		},
		{
			citations: []int{100, 2},
			expect:    2,
		},
		{
			citations: []int{100},
			expect:    1,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			got := hIndex(tt.citations)
			assert.Equal(t, tt.expect, got)
		})
	}
}

func hIndex(citations []int) int {
	if len(citations) == 1 && citations[0] == 0 {
		return 0
	}
	// испиользуем counting sort
	counting := make([]int, len(citations)+1) // последний элемент является бакетом
	// семантика последнего элемента будет "хотя бы counting[i] документов с таким-то количеством ссылок"
	for _, citation := range citations {
		counting[min(citation, len(citations))]++
	}

	h := len(citations)
	papers := counting[h]
	for papers < h {
		h -= 1
		papers += counting[h]
	}

	return h
}

/*

Предыдущее решение

func hIndex(citations []int) int {
	if len(citations) == 1 && citations[0] == 0 {
		return 0
	}
	// испиользуем counting sort
	counting := make([]int, len(citations)+1) // последний элемент является бакетом
	// семантика последнего элемента будет "хотя бы counting[i] документовОкейОк с таким-то количеством ссылок"
	for _, citation := range citations {
		counting[min(citation, len(citations))]++
	}

	minDiff := math.MaxInt
	h := -1
	for i, j := len(citations), 0; i > 0; i-- {
		j += counting[i]
		currDiff := abs(j - i)
		if minDiff >= currDiff {
			minDiff = currDiff
			h = i
		}
	}

	return h
}



*/
