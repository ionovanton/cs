package bench

import (
	"testing"
)

type pair struct {
	first  int
	second int
}

type InterFoo interface {
	Foo(pair) int
}

type Strct1 struct {
	StrctValue int
}

func (s Strct1) Foo(p pair) int {
	return s.StrctValue + p.first + p.second
}

type Strct2 struct {
	StrctValue int
}

func (s Strct2) Foo(p pair) int {
	return s.StrctValue + p.first + p.second
}

func BenchmarkIface(b *testing.B) {
	var resultIface int
	b.Run("InterFoo", func(b *testing.B) {
		var m InterFoo
		m = Strct1{StrctValue: 6742}
		for i := 0; i < b.N; i++ {
			resultIface = m.Foo(pair{i, i})
		}
	})
	println(resultIface)
}

func BenchmarkStrct2(b *testing.B) {
	var resultStrct2 int
	b.Run("Strct2", func(b *testing.B) {
		m := Strct2{6742}
		for i := 0; i < b.N; i++ {
			resultStrct2 = m.Foo(pair{i, i})
		}
	})
	println(resultStrct2)
}
