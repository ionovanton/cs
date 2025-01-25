package main

import (
	"fmt"
	"strings"
)

var goroutinesNum = 3

func formatWork(in int, input string) string {
	return fmt.Sprint(strings.Repeat("  ", in), "█",
		strings.Repeat("  ", goroutinesNum-in),
		in, " recieved work ", input)
}

func main() {
}
