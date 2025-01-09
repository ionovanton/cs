package main

import (
	"testing"
	"time"

	"go.uber.org/goleak"
)

func leak() error {
	go func() {
		time.Sleep(time.Minute)
	}()

	return nil
}

func TestLeakFunction(t *testing.T) {
	defer goleak.VerifyNone(t)

	if err := leak(); err != nil {
		t.Fatal("error not expected")
	}
}
