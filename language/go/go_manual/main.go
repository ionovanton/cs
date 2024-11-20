package main

import (
	"fmt"
)

type Fooer interface {
	Foo()
}

// Now X implements Fooer
type X struct {
	Fooer
}

type KungFoo struct {
}

func (k KungFoo) Foo() {
	fmt.Println("kung foo")
}

func main() {
	x := X{
		Fooer: KungFoo{},
	}
	Bar(x)
}

func Bar(fooer Fooer) {
	fooer.Foo()
}
