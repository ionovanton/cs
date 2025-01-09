package main

import (
	"fmt"
	"go.uber.org/goleak"
	"sync"
	"testing"
)

func LeakProblem3() {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("Goroutine", i)
		}()
	}
}

func TestLeakProblem3(t *testing.T) {
	defer goleak.VerifyNone(t)

	LeakProblem3()
}

func LeakProblemSolution3() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			fmt.Println("Goroutine", j)
		}(i)
	}
	wg.Wait()
}

func TestLeakProblemSolution3(t *testing.T) {
	defer goleak.VerifyNone(t)

	LeakProblemSolution3()
}
