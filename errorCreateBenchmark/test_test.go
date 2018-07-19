package main

import (
	"testing"
)

type ErrTest string

func (e ErrTest) Error() string {
	return string(e)
}

func (e ErrTest) String() string {
	return string(e)
}

var E error

func Benchmark_f1(b *testing.B) {
	err := ErrTest("Some text")
	var e error
	b.Run("formatS", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			e = formatS(err)
		}
		E = e
	})
	b.Run("formatV", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			e = formatV(err)
		}
		E = e
	})
	b.Run("formatSWithFunc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			e = formatSWithFunc(err)
		}
		E = e
	})
	b.Run("formatVWithFunc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			e = formatVWithFunc(err)
		}
		E = e
	})
	b.Run("concat", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			e = concat(err)
		}
		E = e
	})
}
