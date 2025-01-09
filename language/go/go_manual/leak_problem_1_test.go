package main

import (
	"fmt"
	"testing"
	"time"

	"go.uber.org/goleak"
)

func receive1(ch chan int) {
	fmt.Println(<-ch) // Goroutine leaks if no value is sent to ch
}

func LeakProblem1() {
	ch := make(chan int)
	go receive1(ch)
	// No value sent to ch, causing receive to block forever
}

func TestLeakProblem1(t *testing.T) {
	defer goleak.VerifyNone(t)

	LeakProblem1()
}

func receive2(ch chan int) {
	select {
	case val := <-ch:
		fmt.Println(val)
	case <-time.After(time.Second * 5):
		return
	}
}

func LeakProblemSolution1() {
	ch := make(chan int)
	go receive2(ch)
	time.Sleep(time.Second)

	close(ch)
}

func TestLeakProblemSolution1(t *testing.T) {
	defer goleak.VerifyNone(t)

	LeakProblemSolution1()
}
