package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func prepareStrings(n int, strLen int) []string {
	ss := make([]string, n)
	for i := 0; i < n; i++ {
		ss[i] = randStringRunes(strLen)
	}
	return ss
}

func BenchmarkMapReadBeforeWrite(b *testing.B) {
	bb := []struct {
		keys   int
		keyLen int
	}{
		{10, 10},
		{10, 100},
		{100, 10},
		{100, 100},
		{1000, 10},
		{1000, 100},
		{10000, 10},
		{10000, 100},
		{100000, 10},
		{100000, 100},
		{1000000, 10},
		{1000000, 100},
	}
	for _, bm := range bb {
		b.Run(fmt.Sprint(bm.keys, "x", bm.keyLen), func(b *testing.B) {
			strings := prepareStrings(bm.keys, bm.keyLen)

			m := make(map[string]struct{}, len(strings))

			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				key := strings[n%len(strings)]
				if _, ok := m[key]; !ok {
					m[key] = struct{}{}
				}
			}
		})
	}
}

func BenchmarkMapWrite(b *testing.B) {
	bb := []struct {
		keys   int
		keyLen int
	}{
		{10, 10},
		{10, 100},
		{100, 10},
		{100, 100},
		{1000, 10},
		{1000, 100},
		{10000, 10},
		{10000, 100},
		{100000, 10},
		{100000, 100},
		{1000000, 10},
		{1000000, 100},
	}
	for _, bm := range bb {
		b.Run(fmt.Sprint(bm.keys, "x", bm.keyLen), func(b *testing.B) {
			strings := prepareStrings(bm.keys, bm.keyLen)

			m := make(map[string]struct{}, len(strings))

			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				key := strings[n%len(strings)]
				m[key] = struct{}{}
			}
		})
	}
}
