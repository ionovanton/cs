## Progress
### 2 / 25

# Slice
## https://go.dev/blog/slices
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

## Slice tricks https://go.dev/wiki/SliceTricks
```go
func main() {
	a := []int{1, 2, 3}
	b := []int{7, 8, 9, 10}

	// append vector
	// a: [1 2 3]
	// b: [7 8 9 10 1 2 3]
	b = append(b, a...)

	// copy
	// a: [1, 2, 3]
	// c: [1, 2, 3]
	c := make([]int, len(a))
	copy(c, a)

	// cut
	// b: [7 8 1 2 3]
	b = append(b[:2], b[4:]...)

	// delete
	// b: [7 8 2 3]
	i := 2
	b = append(b[:i], b[i+1:]...)
}
```

Leak-aware cut mitigates memory leaks as values are still referenced in previous slice.
```go
func main() {
	a := []*int64{
		PointerToInt64(1),
		PointerToInt64(2),
		PointerToInt64(3),
		PointerToInt64(4),
		PointerToInt64(5),
	}

	// leak-aware cut
	// [1, 2, 5]
	i, j := 2, 4
	copy(a[i:], a[j:])
	for k, n := len(a)-j+i, len(a); k < n; k++ {
		a[k] = nil
	}
	a = a[:len(a)-j+i]

}

func PointerToInt64(i int64) *int64 {
	return &i
}
```

Other tricks
```go
func main() {
	a := []int{1, 2, 3, 4, 5}
	b := []int{42, 21}

	// Expand
	// Insert n elements at position i
	// a: [1 2 3 4 5 0]
	i, n := len(a), 1
	a = append(a[:i], append(make([]int, n), a[i:]...)...)

	// Extend
	// Append n elements
	// a: [1 2 3 4 5 0 0]
	a = append(a, make([]int, n)...)

	// Extend capacity
	// Make sure there is space for next n elements
	// len, cap = 7, 10 --> len, cap = 7, 12
	n = 5
	a = append(make([]int, 0, len(a)+n), a...)

	// Insert
	// b: [97 42 21]
	i = 0
	b = append(b, 0)
	copy(b[i+1:], b[i:])
	b[i] = 97

	// In-place filtering
	// This tricks uses the fact that a slice shares array and capacity as the original,
	// so the storage is reused for filtered slice.
	c := b[:0]
	for _, x := range b {
		if x < 50 {
			c = append(c, x)
		}
	}
}
```

# Interface
### Interface internals https://golang.design/go-questions/interface/iface-eface/
#### `iface` vs `eface`

`iface` и `eface` are basic structures which describe interfaces in Golang. `iface` includes methods and `eface` describes empty interface which does not include any methods.

`iface` includes two pointers: `tab` pointing to `itab` structure which describes interface type and data pointing to concrete value of interface (usually stored on heap).
```go
type itab struct { // 40 bytes on a 64bit arch
    inter *interfacetype // wrapper around _type with some extra information
    _type *_type         // internal representation of go type within runtime
    hash  uint32         // copy of _type.hash. Used for type switches.
    _     [4]byte
    fun   [1]uintptr     // multiple method addresses are stored here; use pointer arithmetics to iterate
}
```

```go
type interfaceType = abi.InterfaceType

type InterfaceType struct {
	Type
	PkgPath Name      // import path
	Methods []Imethod // sorted by hash
}
```

Interface with methods
```go
type iface struct {
	tab  *itab
	data unsafe.Pointer
}
```

Empty interface
```go
type eface struct {
	_type *_type
	data  unsafe.Pointer
}
````

#### `Type` structure
```go
type _type = abi.Type

type Type struct {
    Size_       uintptr
    PtrBytes    uintptr // number of (prefix) bytes in the type that can contain pointers
    Hash        uint32  // hash of type; avoids computation in hash tables
    TFlag       TFlag   // extra type information flags
    Align_      uint8   // alignment of variable with this type
    FieldAlign_ uint8   // alignment of struct field with this type
    Kind_       uint8   // enumeration for C
    // function for comparing objects of this type
    // (ptr to object A, ptr to object B) -> ==?
    Equal func(unsafe.Pointer, unsafe.Pointer) bool
    // GCData stores the GC type data for the garbage collector.
    // If the KindGCProg bit is set in kind, GCData is a GC program.
    // Otherwise it is a ptrmask bitmap. See mbitmap.go for details.
    GCData    *byte
    Str       NameOff // string form
    PtrToThis TypeOff // type for pointer to this type, may be zero
}
```

Some of Golang first class citizens are described by embedded `Type` structure
```go
type ChanType struct {
	Type
	Elem *Type
	Dir  ChanDir
}

type StructType struct {
    Type
    PkgPath Name
    Fields  []StructField
}

type SliceType struct {
    Type
    Elem *Type // slice element type
}
```

#### Comparing interface and `nil`
From previous section we found out that interface consists of `tab` holding type related information and `data` holding data of implemented struct

Interface considered `nil` if and only if `tab` **and** `data` are both `nil`
```go
type Inter interface {
	foo()
}

type Strct struct {
	name string
}

func (g Strct) foo() {
	print(g.name)
}

func main() {
	var inter Inter
	x := inter == nil
	fmt.Printf("inter: type is %v; value is %+v\n", reflect.TypeOf(inter), inter)
	println(x)

	var strct *Strct

	inter = strct
	y := inter == nil
	fmt.Printf("inter: type is %v; value is %+v\n", reflect.TypeOf(inter), inter)
	println(y)

	strct1 := &Strct{name: "some struct"}
	inter = strct1
	z := inter == nil
	fmt.Printf("inter: type is %v; value is %+v\n", reflect.TypeOf(inter), inter)
	println(z)
}
```
```
inter: type is <nil>; value is <nil>
true
inter: type is *main.Strct; value is <nil>
false
inter: type is *main.Strct; value is &{name:some struct}
false
```

Another example regarding comparing interface and `nil`
```go
type MyError struct{}

func (i MyError) Error() string {
	return "MyError"
}

func main() {
	err1 := Process1()
	fmt.Printf("err1: type is %v; value is %+v\n", reflect.TypeOf(err1), err1)
	println(err1 == nil) // false because err1 is interface and its tab is pointer to struct

	err2 := Process2()
	fmt.Printf("err2: type is %v; value is %+v\n", reflect.TypeOf(err2), err2)
	println(err2 == nil) // true because err2 is pointer to struct which is nil
}

func Process1() error {
	var err *MyError = nil
	return err
}

func Process2() *MyError {
	var err *MyError = nil
	return err
}
```
```
err1: type is *main.MyError; value is <nil>
false
err2: type is *main.MyError; value is <nil>
true
```

There is also a way to print`iface` addresses of tab and data
```go
// we know that underlying interface type consists of two pointers,
// so we can imitate underlying structure with custom one
type iface struct {
    itab uintptr
    data uintptr
}

func main() {
    var a interface{} = nil
    
    var b interface{} = (*int)(nil)
    
    x := 5
    var c interface{} = (*int)(&x)
    
    ia := *(*iface)(unsafe.Pointer(&a))
    ib := *(*iface)(unsafe.Pointer(&b))
    ic := *(*iface)(unsafe.Pointer(&c))
    
    fmt.Printf("%+v\n%+v\n%+v\n", ia, ib, ic)
}
```
```
{itab:0 data:0}
{itab:4363494528 data:0}
{itab:4363494528 data:1374390603392}
```

We also can see what `data` pointer (uintptr) holds. Spoiler: it's interface implementing structure.
```go
package main

type Strct struct {
	i int
	s string
}

type Inter interface {
	Foo()
}

func (s Strct) Foo() {
}

type iface struct {
	itab uintptr
	data uintptr
}

func main() {

	x := Strct{
		i: 42,
		s: "some random text",
	}
	var a Inter = (*Strct)(&x)

	ia := *(*iface)(unsafe.Pointer(&a))
	da := *(*Strct)(unsafe.Pointer(ia.data))
	fmt.Printf("%+v\n", ia)
	fmt.Printf("%+v\n", da)
}
```
```
{itab:4363752104 data:1374390628048}
{i:42 s:some random text}
```

#### Compiler determines whether struct implements interface
In some libraries you might have seen the following kind of statement
```go
var _ io.Writer = (*myWriter)(nil)
```
It serves as static check for compiler to decide whether struct implements interface
```go
package main

// Check whether Strct implements Inter
var _ Inter = (*Strct)(nil)
var _ Inter = Strct{}

type Inter interface {
	Foo()
}

type Strct struct {
}

//func (w Strct) Foo() {
//}

func main() {
}
```
```
main.go:4:15: cannot use (*Strct)(nil) (value of type *Strct) as Inter value in variable declaration: *Strct does not implement Inter (missing method Foo)
main.go:5:15: cannot use Strct{} (value of type Strct) as Inter value in variable declaration: Strct does not implement Inter (missing method Foo)
```
When compiler performs type checking phase it reports unimplemented interface

#### How interface is being built https://studygolang.com/articles/28873
Let's take a look at the following code
```go
package main

type InterFoo interface {
	Foo() int32
}

type Strct1 struct {
	StrctValue int32
}

//go:noinline
func (s Strct1) Foo() int32 {
	return s.StrctValue
}

func main() {
	// It is the only way to make compiler not optimize out interface creation,
	// so we initialize `m` and only then assign struct to it
	var m InterFoo
	m = Strct1{StrctValue: 6742}

	// This call just makes sure that the interface is actually used.
	// Without this call, the linker would see that the interface defined above
	// is in fact never used, and thus would optimize it out of the final
	// executable.
	// As of Go 1.23 I see tendency for compiler optimizing it out regardless of call.
	// So you need to disable compiler optimizations in order to see what's going on under the hood.
	someFooValue := m.Foo()

	println(someFooValue)
}
```

Execute `go tool compile -S -N -l main.go > compiler.log`
It produces the following compiler output. Here we can see `itab` dump.
```
	0x001c 00028 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:20)	MOVW	ZR, main..autotmp_2-28(SP)
	0x0020 00032 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:20)	MOVD	$6742, R1
	0x0024 00036 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:20)	MOVW	R1, main..autotmp_2-28(SP)
	0x0028 00040 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:20)	MOVW	R1, main..autotmp_3-32(SP)
	0x002c 00044 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:20)	MOVD	$main..autotmp_3-32(SP), R1
	0x0030 00048 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:20)	PCDATA	$0, $-2
	0x0030 00048 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:20)	MOVB	(R1), R27
	0x0034 00052 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:20)	PCDATA	$0, $-1
	0x0034 00052 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:20)	MOVWU	(R1), R0
	0x0038 00056 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:20)	MOVW	R0, main..autotmp_4-36(SP)
	0x003c 00060 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:20)	PCDATA	$1, $0
	0x003c 00060 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:20)	CALL	runtime.convT32(SB)
	0x0040 00064 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:20)	MOVD	R0, main..autotmp_5-8(SP)
	0x0044 00068 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:20)	MOVD	$go:itab.<unlinkable>.Strct1,<unlinkable>.InterFoo(SB), R1
	0x004c 00076 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:20)	MOVD	R1, main.m-24(SP)
	0x0050 00080 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:20)	MOVD	R0, main.m-16(SP)
```

It will be used by linker later
```
go:itab.<unlinkable>.Strct1,<unlinkable>.InterFoo SRODATA dupok size=32
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 ed 1c 7f 68 00 00 00 00 00 00 00 00 00 00 00 00  ...h............
	rel 0+8 t=R_ADDR type:<unlinkable>.InterFoo+0
	rel 8+8 t=R_ADDR type:<unlinkable>.Strct1+0
	rel 24+8 t=RelocType(-32767) <unlinkable>.(*Strct1).Foo+0
```

Only hash is set at compile time
```go
type itab struct {
    inter *interfacetype // offset 0x00 ($00) | 00 00 00 00 00 00 00 00
    _type *_type         // offset 0x08 ($08) | 00 00 00 00 00 00 00 00
    hash  uint32         // offset 0x16 ($16) | ed 1c 7f 68
    _     [4]byte        // offset 0x18 ($20) | 00 00 00 00
    fun   [1]uintptr     // offset 0x20 ($24) | 00 00 00 00 00 00 00 00
}
```

We also see that `itab` structure corresponds to linker offset
```
	rel 0+8 t=R_ADDR type:<unlinkable>.InterFoo+0             offset to `inter *interfacetype`
	rel 8+8 t=R_ADDR type:<unlinkable>.Strct1+0               offset to `_type *_type`
	rel 24+8 t=RelocType(-32767) <unlinkable>.(*Strct1).Foo+0 offset to `fun   [1]uintptr`
```
When linker will finish its job, the `itab` will be complete at this point

Generate an ELF file using `GOOS=linux GOARCH=arm64 go build -o main.bin -gcflags='-N -l' main.go `
I'll use online tool to read ELF file in order to find virtual address of complete `InterFoo`'s `itab` (http://www.sunshine2k.de/coding/javascript/onlineelfviewer/onlineelfviewer.html)
```
Nr             Value               Size                Info (Binding|Type) Other          Shndx          Name
1562           0x00000000000AF298  0x0000000000000020  GLOBAL | OBJECT     DEFAULT        0x0002         go:itab.main.Strct1,main.InterFoo
```

The only thing is needed is to find .rodata offset. We can find it in section header tables
```
Nr             Name                Type           Address             Offset              Size                Link           Info           AddrAlign           EntSize             Flags
2              .rodata             SHT_PROGBITS   0x0000000000080000  0x0000000000070000  0x0000000000030262  0x00000000     0x00000000     0x0000000000000020  0x0000000000000000  Alloc
```

We have everything we need:
Section offset   is 0x0000000000070000 = 458752
Section vma      is 0x0000000000080000 = 524288
Itab symbol vma  is 0x00000000000AF298 = 717464
Itab symbol size is 0x0000000000000020 = 32

In order to find complete `itab`, its symbol offset we'll use the formula
`symbol offset = symbol vma - section vma + section offset`
Transforms to
`717464 - 524288 + 458752 = 651928`

Upon execution `dd if=main.bin of=/dev/stdout bs=1 count=32 skip=651928 2>/dev/null | hexdump -C` we'll see complete itab
```
00000000  e0 79 08 00 00 00 00 00  60 92 08 00 00 00 00 00  |.y......`.......|
00000010  ed 1c 7f 68 00 00 00 00  40 71 07 00 00 00 00 00  |...h....@q......|
00000020
```

```go
type itab struct {
    inter *interfacetype // offset 0x00 ($00) | e0 79 08 00 00 00 00 00
    _type *_type         // offset 0x08 ($08) | 60 92 08 00 00 00 00 00
    hash  uint32         // offset 0x16 ($16) | ed 1c 7f 68
    _     [4]byte        // offset 0x18 ($20) | 00 00 00 00
    fun   [1]uintptr     // offset 0x20 ($24) | 40 71 07 00 00 00 00 00
}
```

#### Polymorphism and how much it costs
```go
package main

type InterFoo interface {
	Foo(int32) int32
	Bar(int64) int64
	Meow(int16) int16
}

type Strct1 struct {
	StrctValue int32
}

//go:noinline
func (s Strct1) Foo(fooValue int32) int32 {
	return s.StrctValue + fooValue
}

//go:noinline
func (s Strct1) Bar(barValue int64) int64 {
	return int64(s.StrctValue) + barValue
}

//go:noinline
func (s Strct1) Meow(meowValue int16) int16 {
	return int16(s.StrctValue) + meowValue
}

func main() {
	var m InterFoo
	m = Strct1{StrctValue: 6742}

	someFooValue := m.Foo(42)
	someBarValue := m.Bar(788)
	someMeowValue := m.Meow(9128)

	println(someFooValue, someBarValue, someMeowValue)
}

```

Compiler output for main function
```
main.main STEXT size=224 args=0x0 locals=0x48 funcid=0x0 align=0x0
	0x0000 00000 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:28)	TEXT	main.main(SB), ABIInternal, $96-0
	0x0000 00000 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:28)	MOVD	16(g), R16
	0x0004 00004 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:28)	PCDATA	$0, $-2
	0x0004 00004 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:28)	CMP	R16, RSP
	0x0008 00008 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:28)	BLS	224
	0x000c 00012 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:28)	PCDATA	$0, $-1
	0x000c 00012 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:28)	MOVD.W	R30, -96(RSP)
	0x0010 00016 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:28)	MOVD	R29, -8(RSP)
	0x0014 00020 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:28)	SUB	$8, RSP, R29
	0x0018 00024 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:28)	FUNCDATA	$0, gclocals·J5F+7Qw7O7ve2QcWC7DpeQ==(SB)
	0x0018 00024 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:28)	FUNCDATA	$1, gclocals·3TP5whGZWqE6ZxU0iS+iBA==(SB)
	0x0018 00024 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:29)	STP	(ZR, ZR), main.m-24(SP)
	0x001c 00028 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:30)	MOVW	ZR, main..autotmp_4-36(SP)
	0x0020 00032 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:30)	MOVD	$6742, R1
	0x0024 00036 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:30)	MOVW	R1, main..autotmp_4-36(SP)
	0x0028 00040 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:30)	MOVW	R1, main..autotmp_5-40(SP)
	0x002c 00044 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:30)	MOVD	$main..autotmp_5-40(SP), R1
	0x0030 00048 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:30)	PCDATA	$0, $-2
	0x0030 00048 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:30)	MOVB	(R1), R27
	0x0034 00052 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:30)	PCDATA	$0, $-1
	0x0034 00052 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:30)	MOVWU	(R1), R0
	0x0038 00056 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:30)	MOVW	R0, main..autotmp_6-44(SP)
	0x003c 00060 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:30)	PCDATA	$1, $0
	0x003c 00060 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:30)	CALL	runtime.convT32(SB)
	0x0040 00064 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:30)	MOVD	R0, main..autotmp_7-8(SP)
	0x0044 00068 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:30)	MOVD	$go:itab.<unlinkable>.Strct1,<unlinkable>.InterFoo(SB), R1
	0x004c 00076 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:30)	MOVD	R1, main.m-24(SP)
	0x0050 00080 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:30)	MOVD	R0, main.m-16(SP)
	0x0054 00084 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:32)	MOVD	$42, R1
	0x0058 00088 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:32)	PCDATA	$1, $1
	0x0058 00088 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:32)	CALL	<unlinkable>.(*Strct1).Foo(SB)
	0x005c 00092 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:32)	MOVW	R0, main.someFooValue-48(SP)
	0x0060 00096 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:33)	MOVD	main.m-24(SP), R1
	0x0064 00100 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:33)	PCDATA	$0, $-2
	0x0064 00100 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:33)	MOVB	(R1), R27
	0x0068 00104 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:33)	PCDATA	$0, $-1
	0x0068 00104 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:33)	MOVD	24(R1), R1
	0x006c 00108 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:33)	MOVD	main.m-16(SP), R0
	0x0070 00112 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:33)	MOVD	R1, R2
	0x0074 00116 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:33)	MOVD	$788, R1
	0x0078 00120 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:33)	CALL	(R2)
	0x007c 00124 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:33)	MOVD	R0, main.someBarValue-32(SP)
	0x0080 00128 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:34)	MOVD	main.m-24(SP), R1
	0x0084 00132 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:34)	PCDATA	$0, $-2
	0x0084 00132 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:34)	MOVB	(R1), R27
	0x0088 00136 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:34)	PCDATA	$0, $-1
	0x0088 00136 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:34)	MOVD	40(R1), R1
	0x008c 00140 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:34)	MOVD	main.m-16(SP), R0
	0x0090 00144 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:34)	MOVD	R1, R2
	0x0094 00148 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:34)	MOVD	$9128, R1
	0x0098 00152 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:34)	PCDATA	$1, $0
	0x0098 00152 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:34)	CALL	(R2)
	0x009c 00156 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:34)	MOVH	R0, main.someMeowValue-50(SP)
	0x00a0 00160 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:36)	CALL	runtime.printlock(SB)
	0x00a4 00164 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:36)	MOVW	main.someFooValue-48(SP), R1
	0x00a8 00168 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:36)	MOVD	R1, R0
	0x00ac 00172 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:36)	CALL	runtime.printint(SB)
	0x00b0 00176 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:36)	CALL	runtime.printsp(SB)
	0x00b4 00180 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:36)	MOVD	main.someBarValue-32(SP), R0
	0x00b8 00184 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:36)	CALL	runtime.printint(SB)
	0x00bc 00188 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:36)	CALL	runtime.printsp(SB)
	0x00c0 00192 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:36)	MOVH	main.someMeowValue-50(SP), R1
	0x00c4 00196 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:36)	MOVD	R1, R0
	0x00c8 00200 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:36)	CALL	runtime.printint(SB)
	0x00cc 00204 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:36)	CALL	runtime.printnl(SB)
	0x00d0 00208 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:36)	CALL	runtime.printunlock(SB)
	0x00d4 00212 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:37)	LDP	-8(RSP), (R29, R30)
	0x00d8 00216 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:37)	ADD	$96, RSP
	0x00dc 00220 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:37)	RET	(R30)
	0x00e0 00224 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:37)	NOP
	0x00e0 00224 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:28)	PCDATA	$1, $-1
	0x00e0 00224 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:28)	PCDATA	$0, $-2
	0x00e0 00224 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:28)	MOVD	R30, R3
	0x00e4 00228 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:28)	CALL	runtime.morestack_noctxt(SB)
	0x00e8 00232 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:28)	PCDATA	$0, $-1
	0x00e8 00232 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:28)	JMP	0
	0x0000 90 0b 40 f9 ff 63 30 eb c9 06 00 54 fe 0f 1a f8  ..@..c0....T....
	0x0010 fd 83 1f f8 fd 23 00 d1 ff 7f 04 a9 ff 37 00 b9  .....#.......7..
	0x0020 c1 4a 83 d2 e1 37 00 b9 e1 33 00 b9 e1 c3 00 91  .J...7...3......
	0x0030 3b 00 80 39 20 00 40 b9 e0 2f 00 b9 00 00 00 94  ;..9 .@../......
	0x0040 e0 2b 00 f9 01 00 00 90 21 00 00 91 e1 23 00 f9  .+......!....#..
	0x0050 e0 27 00 f9 41 05 80 d2 00 00 00 94 e0 2b 00 b9  .'..A........+..
	0x0060 e1 23 40 f9 3b 00 80 39 21 0c 40 f9 e0 27 40 f9  .#@.;..9!.@..'@.
	0x0070 e2 03 01 aa 81 62 80 d2 40 00 3f d6 e0 1f 00 f9  .....b..@.?.....
	0x0080 e1 23 40 f9 3b 00 80 39 21 14 40 f9 e0 27 40 f9  .#@.;..9!.@..'@.
	0x0090 e2 03 01 aa 01 75 84 d2 40 00 3f d6 e0 4f 00 79  .....u..@.?..O.y
	0x00a0 00 00 00 94 e1 2b 80 b9 e0 03 01 aa 00 00 00 94  .....+..........
	0x00b0 00 00 00 94 e0 1f 40 f9 00 00 00 94 00 00 00 94  ......@.........
	0x00c0 e1 4f 80 79 e0 03 01 aa 00 00 00 94 00 00 00 94  .O.y............
	0x00d0 00 00 00 94 fd fb 7f a9 ff 83 01 91 c0 03 5f d6  .............._.
	0x00e0 e3 03 1e aa 00 00 00 94 c6 ff ff 17 00 00 00 00  ................
	rel 0+0 t=R_USEIFACE type:main.Strct1+0
	rel 0+0 t=R_USEIFACEMETHOD type:main.InterFoo+104
	rel 0+0 t=R_USEIFACEMETHOD type:main.InterFoo+96
	rel 0+0 t=R_USEIFACEMETHOD type:main.InterFoo+112
	rel 60+4 t=R_CALLARM64 runtime.convT32+0
	rel 68+8 t=R_ADDRARM64 go:itab.<unlinkable>.Strct1,<unlinkable>.InterFoo+0
	rel 88+4 t=R_CALLARM64 <unlinkable>.(*Strct1).Foo+0
	rel 120+0 t=R_CALLIND +0
	rel 152+0 t=R_CALLIND +0
	rel 160+4 t=R_CALLARM64 runtime.printlock+0
	rel 172+4 t=R_CALLARM64 runtime.printint+0
	rel 176+4 t=R_CALLARM64 runtime.printsp+0
	rel 184+4 t=R_CALLARM64 runtime.printint+0
	rel 188+4 t=R_CALLARM64 runtime.printsp+0
	rel 200+4 t=R_CALLARM64 runtime.printint+0
	rel 204+4 t=R_CALLARM64 runtime.printnl+0
	rel 208+4 t=R_CALLARM64 runtime.printunlock+0
	rel 228+4 t=R_CALLARM64 runtime.morestack_noctxt+0
```

Call `m.Foo()` with assigning to interface
```
	0x0054 00084 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:32)	MOVD	$42, R1
	0x0058 00088 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:32)	PCDATA	$1, $1
	0x0058 00088 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:32)	CALL	<unlinkable>.(*Strct1).Foo(SB)
```
Call `m.Foo()` without assigning to interface – `m := Strct1{StrctValue: 6742}`
```
    0x0024 00036 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:32)   MOVD    $42, R1
    0x0028 00040 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:32)   PCDATA  $1, $0
    0x0028 00040 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:32)   CALL    main.Strct1.Foo(SB)
    0x002c 00044 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:32)   MOVW    R0, main.someFooValue-16(SP)
```

Call `m.Bar()` with interface
```
	0x0060 00096 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:33)	MOVD	main.m-24(SP), R1
	0x0064 00100 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:33)	PCDATA	$0, $-2
	0x0064 00100 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:33)	MOVB	(R1), R27
	0x0068 00104 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:33)	PCDATA	$0, $-1
	0x0068 00104 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:33)	MOVD	24(R1), R1        ;; dereference R1 and move 24 bytes further, store the result in R1; this will be our function m.Bar()
	0x006c 00108 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:33)	MOVD	main.m-16(SP), R0 ;; catch the result of m.Bar()
	0x0070 00112 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:33)	MOVD	R1, R2            ;; move function address to R2
	0x0074 00116 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:33)	MOVD	$788, R1          ;; prepare function `barValue int64` argument
	0x0078 00120 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:33)	CALL	(R2)              ;; call m.Bar()
	0x007c 00124 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:33)	MOVD	R0, main.someBarValue-32(SP)
```
Call `m.Bar()` without interface
```
    0x0030 00048 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:33)   MOVW    main.m-12(SP), R0
    0x0034 00052 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:33)   MOVD    $788, R1
    0x0038 00056 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:33)   CALL    main.Strct1.Bar(SB)
    0x003c 00060 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:33)   MOVD    R0, main.someBarValue-8(SP)
```

Call `m.Meow()` with interface
```
	0x0080 00128 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:34)	MOVD	main.m-24(SP), R1
	0x0084 00132 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:34)	PCDATA	$0, $-2
	0x0084 00132 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:34)	MOVB	(R1), R27
	0x0088 00136 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:34)	PCDATA	$0, $-1
	0x0088 00136 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:34)	MOVD	40(R1), R1
	0x008c 00140 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:34)	MOVD	main.m-16(SP), R0
	0x0090 00144 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:34)	MOVD	R1, R2
	0x0094 00148 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:34)	MOVD	$9128, R1
	0x0098 00152 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:34)	PCDATA	$1, $0
	0x0098 00152 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:34)	CALL	(R2)
	0x009c 00156 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:34)	MOVH	R0, main.someMeowValue-50(SP)
```

Call `m.Meow()` without interface
```
    0x0040 00064 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:34)   MOVW    main.m-12(SP), R0
    0x0044 00068 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:34)   MOVD    $9128, R1
    0x0048 00072 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:34)   CALL    main.Strct1.Meow(SB)
    0x004c 00076 (/Users/ayionov/Desktop/cs/language/go/go_manual/main.go:34)   MOVH    R0, main.someMeowValue-18(SP)
```

It must be obvious: the more assembler instructions, greater the cost of a single call

Interface symbol
```
go:itab.<unlinkable>.Strct1,<unlinkable>.InterFoo SRODATA dupok size=48
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 1b 5c ca a7 00 00 00 00 00 00 00 00 00 00 00 00  .\..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 0+8 t=R_ADDR type:<unlinkable>.InterFoo+0
	rel 8+8 t=R_ADDR type:<unlinkable>.Strct1+0
	rel 24+8 t=RelocType(-32767) <unlinkable>.(*Strct1).Bar+0
	rel 32+8 t=RelocType(-32767) <unlinkable>.(*Strct1).Foo+0
	rel 40+8 t=RelocType(-32767) <unlinkable>.(*Strct1).Meow+0
```
```go
type itab struct {
    inter *interfacetype // offset 0x00 ($00) | 00 00 00 00 00 00 00 00
    _type *_type         // offset 0x08 ($08) | 00 00 00 00 00 00 00 00
    hash  uint32         // offset 0x16 ($16) | 1b 5c ca a7
    _     [4]byte        // offset 0x18 ($20) | 00 00 00 00
    fun   [1]uintptr     // offset 0x20 ($28) | 00 00 00 00 00 00 00 00 // place for the first function – Foo()
	                     // offset 0x28 ($36) | 00 00 00 00 00 00 00 00 // Bar()
						 // offset 0x36 ($42) | 00 00 00 00 00 00 00 00 // Meow()
}
```

Look at complete `itab`

Section offset   is 0x0000000000070000 = 458752
Section vma      is 0x0000000000080000 = 524288
Itab symbol vma  is 0x00000000000AF440 = 717888
Itab symbol size is 0x0000000000000030 = 48

`symbol offset = 717888 - 524288 + 458752 = 652352`
`dd if=main.bin of=/dev/stdout bs=1 count=48 skip=652352 2>/dev/null | hexdump -C`
```
00000000  e0 85 08 00 00 00 00 00  40 ae 08 00 00 00 00 00  |........@.......|
00000010  ed 1c 7f 68 00 00 00 00  00 72 07 00 00 00 00 00  |...h.....r......|
00000020  a0 72 07 00 00 00 00 00  40 73 07 00 00 00 00 00  |.r......@s......|
00000030
```

`itab` grows in size as interface gets more functions
```go
type itab struct {
    inter *interfacetype // offset 0x00 ($00) | e0 85 08 00 00 00 00 00
    _type *_type         // offset 0x08 ($08) | 40 ae 08 00 00 00 00 00
    hash  uint32         // offset 0x16 ($16) | 1b 5c ca a7
    _     [4]byte        // offset 0x18 ($20) | 00 00 00 00
    fun   [1]uintptr     // offset 0x20 ($28) | 00 72 07 00 00 00 00 00 // place for the first function – Foo()
	                     // offset 0x28 ($36) | a0 72 07 00 00 00 00 00 // Bar()
						 // offset 0x36 ($42) | 40 73 07 00 00 00 00 00 // Meow()
}
```

Running `go test -bench=. -benchmem`
```go
type pair struct {
	first  int
	second int
}

type InterFoo interface {
	Foo(*pair) int
}

type Strct1 struct {
	StrctValue int
}

func (s Strct1) Foo(p *pair) int {
	return s.StrctValue + p.first + p.second
}

type Strct2 struct {
	StrctValue int
}

func (s Strct2) Foo(p *pair) int {
	return s.StrctValue + p.first + p.second
}

func BenchmarkIface(b *testing.B) {
	var resultIface int
	b.Run("InterFoo", func(b *testing.B) {
		var m InterFoo
		m = Strct1{StrctValue: 6742}
		for i := 0; i < b.N; i++ {
			resultIface = m.Foo(&pair{i, i})
		}
	})
	println(resultIface)
}

func BenchmarkStrct2(b *testing.B) {
	var resultStrct2 int
	b.Run("Strct2", func(b *testing.B) {
		m := Strct2{6742}
		for i := 0; i < b.N; i++ {
			resultStrct2 = m.Foo(&pair{i, i})
		}
	})
	println(resultStrct2)
}
```

will produce the following result
```
goos: darwin
goarch: arm64
pkg: go_manual/bench
cpu: Apple M2 Pro
BenchmarkIface/InterFoo-12              97989724                12.21 ns/op           16 B/op          1 allocs/op
195986188
BenchmarkStrct2/Strct2-12               1000000000               0.4206 ns/op          0 B/op          0 allocs/op
2000006740
PASS
ok      go_manual/bench 2.544s
```
Pure structure almost 30 times faster than interface analogue. Interface also results in 1 heap allocation per operation.

Changing `Foo`'s `p` argument argument to a nonpointer `p pair` will lay different result
```
goos: darwin
goarch: arm64
pkg: go_manual/bench
cpu: Apple M2 Pro
BenchmarkIface/InterFoo-12              561276046                1.971 ns/op           0 B/op          0 allocs/op
1122558832
BenchmarkStrct2/Strct2-12               1000000000               0.4067 ns/op          0 B/op          0 allocs/op
2000006740
PASS
ok      go_manual/bench 2.053s
```

Almost 5 times faster now and no heap allocations this time. Perhaps compiler optimizes a call somehow.





# Escape analysis

#### Interface usage and escaping to heap using `go tool compile -m main.go`

First let's try declaring structs `X` and `Y`. `X` will implement interface `I` whilst Y will not.
```go
type I interface {
	Foo()
}

type X struct {
}

type Y struct {
}

func (x X) Foo() {}

func main() {
	var x I

	x = X{}
	x.Foo()

	y := Y{}
	_ = y
}
```
```
...
manual/main.go:18:7: X{} escapes to heap
```

Now `Y` will implement `I` as well but without calling `Foo()`
```go
type I interface {
	Foo()
}

type X struct {
}

type Y struct {
}

func (x X) Foo() {}

func (y Y) Foo() {}

func main() {
	var x, y I

	x = X{}
	x.Foo()

	y = Y{}
	_ = y
}
```
```
manual/main.go:20:7: X{} escapes to heap
manual/main.go:23:7: Y{} does not escape
```

1. Go's escape analysis determines whether a variable needs to be heap-allocated, depending on whether the variable's lifetime extends beyond the function's scope.
2. Method calls on interfaces that refer to stack-allocated variables could lead to invalid memory references, so the Go compiler forces these variables to escape to the heap.

Only when we make `y` call `Foo()` then `y` will escape to heap.
```go
type I interface {
	Foo()
}

type X struct {
}

type Y struct {
}

func (x X) Foo() {}

func (y Y) Foo() {}

func main() {
	var x, y I

	x = X{}
	x.Foo()

	y = Y{}
	y.Foo()
}
```
```
manual/main.go:20:7: X{} escapes to heap
manual/main.go:23:7: Y{} escapes to heap
```



---

TODO:
1. interface
  - internals
  - type assertion
  - placing: near implementation or in one separate file (probably should move to architecture related part)
2. for range loop
3. errors, panics and os.exit
  - erros.As, errors.Is
4. address semantics
5. gc
6. defer
7. map
8. strings
9. closure
10. marshalling internals
  - custom marshalling
11. goroutines and scheduler
  - net poller
  - context switch (including internals)
  - internals
  - GOMAXPROCS
12. channels
13. race condition and data race
14. context
15. select
16. sync.Map
17. rwmutex, mutex
18. memory layout
  - escape analysis
19. memory leaks
20. pprof
21. benchmarking
22. effective go
23. uber go code guideline
24. avito go code guideline
25. patterns
26. go mistakes
27. go assembler

https://github.com/emluque/golang-internals-resources?tab=readme-ov-file



---
```go

```
```

```
