package main

import (
	"context"
	"fmt"
	"go.uber.org/goleak"
	"testing"
	"time"
)

func LeakProblem2() {
	go func() {
		for {
			time.Sleep(10 * time.Second)
			fmt.Println("Doing work")
		}
	}()
	// The goroutine runs indefinitely
}

func TestLeakProblem2(t *testing.T) {
	defer goleak.VerifyNone(t)

	LeakProblem2()
}

func LeakProblemSolution2() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			return
		case <-time.After(2 * time.Second):
			// doing work
		}
	}(ctx)

	// Some work
	time.Sleep(time.Second * 2)
}

func TestLeakProblemSolution2(t *testing.T) {
	defer goleak.VerifyNone(t)

	LeakProblemSolution2()
}
