package main

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := []int{42, 21}

	// Expand
	// Insert n elements at position i
	// a: [1 2 3 4 5 0]
	i, n := len(a), 1
	a = append(a[:i], append(make([]int, n), a[i:]...)...)

	// Extend
	// Append n elements
	// a: [1 2 3 4 5 0 0]
	a = append(a, make([]int, n)...)

	// Extend capacity
	// Make sure there is space for next n elements
	// len, cap = 7, 10 --> len, cap = 7, 12
	n = 5
	a = append(make([]int, 0, len(a)+n), a...)

	// Insert
	// b: [97 42 21]
	i = 0
	b = append(b, 0)
	copy(b[i+1:], b[i:])
	b[i] = 97

	// In-place filtering
	// This tricks uses the fact that a slice shares array and capacity as the original,
	// so the storage is reused for filtered slice.
	c := b[:0]
	for _, x := range b {
		if x < 50 {
			c = append(c, x)
		}
	}
}
