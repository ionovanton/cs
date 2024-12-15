package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 0, 0, 0}
	b := []int{2, 5, 6}
	m, n := 3, 3

	//a := []int{0}
	//b := []int{}
	//n, m := 1, 0

	//a := []int{0}
	//b := []int{1}
	//n := 1
	//m := 1

	//a := []int{1, 2, 3, 0, 0, 0}
	//b := []int{2, 5, 6}
	//m, n := 3, 3

	merge(a, n, b, m)
	// swap(a, 2, 5)
	fmt.Println(a)
}

func swap(a []int, i, j int) {
	tmp := a[j]
	a[j] = a[i]
	a[i] = tmp
}

func merge(a []int, n int, b []int, m int) {
	if m == 0 {
		return
	}
	if n == 1 {
		a[0] = b[0]
		return
	}
	for i := n; i > 0; i-- {
		a[i+m-1] = b[i-1]
	}

	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			swap(a, i, i+1)
		}
	}
}
