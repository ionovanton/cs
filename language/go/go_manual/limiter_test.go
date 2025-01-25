package main

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

func Test_Limiter(t *testing.T) {
	type test struct {
		name string
		d    time.Duration
	}

	tests := []test{
		{
			name: "success",
			d:    time.Second,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			limiter()

			select {
			case <-time.After(time.Second):
				if timesEvoke.Load() > 5 {
					fmt.Println("FAIL")
					return
				}
			}
		})
	}
}

var timesEvoke = atomic.Int32{}

func limiter() {
	quota := make(chan struct{}, 10)

	t := time.NewTicker(time.Nanosecond)

	go func() {
		once := time.NewTicker(time.Second)
		for range once.C {
			for i := 0; i < 10; i++ {
				<-quota
			}
		}
	}()

	times := 100

	i := 0
	for tick := range t.C {
		if i == times {
			return
		}
		a := []int{1, 2, 3}
		quota <- struct{}{}
		Foo(i, tick, a)
		i++
	}

}

func Foo(i int, t time.Time, args []int) {
	fmt.Println(i, t)
	timesEvoke.Add(1)
}
