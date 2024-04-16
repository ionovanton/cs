## Progress
### 1 / 4

## Slice
### https://go.dev/blog/slices
A slice is a data structure describing a contiguous section of an array stored separately from the slice variable itself. A slice is not an array. A slice describes a piece of an array.

```go
func main() {
	a := make([]int, 3)
	a[2] = 3

	b := a[2:3]

	fmt.Println(a)
	fmt.Println(b)

	a[2] = 42

	fmt.Println(a)
	fmt.Println(b)
}
```
```
[0 0 3]
[3]
[0 0 42]
[42]
```

You’ll often hear experienced Go programmers talk about the “slice header” because that really is what’s stored in a slice variable. For instance, when you call a function that takes a slice as an argument, such as `foo`, that header is what gets passed to the function. In this call,
the slice argument that is passed to the `foo` function is, in fact, a “slice header”.
```go
func foo(slice []int) {
	slice[0] = 42
}

func main() {
	slice := make([]int, 1)
	fmt.Println(slice)
	foo(slice)
	fmt.Println(slice)
}
```
```
[0]
[42]
```


Even though the slice header is passed by value, the header includes a pointer to elements of an array, so both the original slice header and the copy of the header passed to the function describe the same array. Therefore, when the function returns, the modified elements can be seen through the original slice variable.
```go
func foo(slice []int) {
	for i := range slice {
		slice[i] += 2
	}
}

func main() {
	var buffer [10]int
	slice := buffer[5:10]
	for i := 0; i < len(slice); i++ {
		slice[i] = i
	}
	fmt.Println("before", slice, buffer)
	foo(slice)
	fmt.Println("after", slice, buffer)
}
```
```
before [0 1 2 3 4] [0 0 0 0 0 0 1 2 3 4]
after [2 3 4 5 6] [0 0 0 0 0 2 3 4 5 6]
```

Here we see that the contents of a slice argument can be modified by a function, but its *header cannot*. The length stored in the slice variable is not modified by the call to the function, since the function is passed a copy of the slice header, not the original.
```go
func foo(slice []int) []int {
	slice = slice[0 : len(slice)-1]
	slice[0] = 42
	return slice
}

func main() {
	slice := make([]int, 5)

	fmt.Println("Before: len(slice) =", len(slice))
	newSlice := foo(slice)
	fmt.Println("After:  len(slice) =", len(slice))
	fmt.Println("After:  len(newSlice) =", len(newSlice))
}
```
```
Before: len(slice) = 5
After:  len(slice) = 5
After:  len(newSlice) = 4
```

```go
func PtrSubtractOneFromLength(slicePtr *[]int) {
	*slicePtr = (*slicePtr)[0 : len(*slicePtr)-1]

	// same as above
	// slice := *slicePtr
	// *slicePtr = slice[0 : len(slice)-1]
}

func main() {
	slice := make([]int, 5)
	fmt.Println("Before: len(slice) =", len(slice))
	PtrSubtractOneFromLength(&slice)
	fmt.Println("After:  len(slice) =", len(slice))
}
```
```
Before: len(slice) = 5
After:  len(slice) = 4
```

```go
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
```



https://go.dev/wiki/SliceTricks

https://go.dev/blog/slices-intro

https://research.swtch.com/godata
---
```go

```
```

```