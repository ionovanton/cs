package main

import (
	"fmt"
	"unsafe"
)

type T struct {
	i int
}

type ValueMethodCaller interface {
	valueMethod()
}

type PointerMethodCaller interface {
	pointerMethod()
}

// Pointer type receiver
func (receiver *T) pointerMethod() {
	fmt.Printf("Pointer method called on \t%#v with address %p\n", *receiver, receiver)
}

// Value type receiver
func (receiver T) valueMethod() {
	fmt.Printf("Value method called on \t\t%#v with address %p\n", receiver, &receiver)
}

func callValueMethodOnInterface(v ValueMethodCaller) {
	v.valueMethod()
}

func callPointerMethodOnInterface(p PointerMethodCaller) {
	p.pointerMethod()
}

func main() {
	var iface interface{} = (int32)(0)
	// This takes address of the value. Unsafe but works. Not guaranteed to work
	// after possible implementation change!
	var px uintptr = (*[2]uintptr)(unsafe.Pointer(&iface))[1]

	iface = (int32)(1)
	var py uintptr = (*[2]uintptr)(unsafe.Pointer(&iface))[1]

	fmt.Printf("First pointer %#v,  second pointer %#v", px, py)
}
