package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var slice []int //          -- len: 0, cap: 0, slice == nil: true, sliceHeader: &{0 0 0}
	// slice := []int(nil)      -- len: 0, cap: 0, slice == nil: true, sliceHeader: &{0 0 0}
	// slice := []int{}         -- len: 0, cap: 0, slice == nil: false, sliceHeader: &{4302501024 0 0}
	// slice := make([]int, 0)  -- len: 0, cap: 0, slice == nil: false, sliceHeader: &{4302501024 0 0}

	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&slice))

	fmt.Printf(
		"len: %d, cap: %d, slice == nil: %t, sliceHeader: %v\n",
		len(slice), cap(slice), slice == nil, sliceHeader,
	)
}
