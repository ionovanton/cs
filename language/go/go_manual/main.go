package main

import (
	"bufio"
	"os"
)

//go:noinline
func main() {
	r := bufio.NewReader(os.Stdin) // does not escape
	line, _, _ := r.ReadLine()     // escapes

	s := make([]string, 16) // escapes because end size is unknown
	s[0] = string(line)     // escapes
	println(s)
}
