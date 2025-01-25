package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func FirstKthElementsMatch[T any](t *testing.T, expect, got []T, k int) {
	assert.Equal(t, expect, got[0:k])
}
