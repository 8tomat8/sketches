package main

import (
	"testing"

	"time"

	"github.com/coocood/freecache"
	"github.com/patrickmn/go-cache"
)

func BenchmarkFreeCache(b *testing.B) {
	c := freecache.NewCache(0)
	data := []byte{1}
	for i := 0; i < b.N; i++ {
		c.SetInt(int64(i), data, 60)
	}
}

func BenchmarkGoCache(b *testing.B) {
	c := cache.New(30*time.Second, 10*time.Minute)

	for i := 0; i < b.N; i++ {
		c.SetDefault(letterRunes[i], true)
	}
}

var letterRunes = []string("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
