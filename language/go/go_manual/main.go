package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Printf("A->")
	}()

	go func() {
		fmt.Printf("B")
	}()
	time.Sleep(time.Second)
}
