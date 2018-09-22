package sorting

import (
	"math/rand"
)

func Gen(n uint, max int) []int {
	if n < 1 {
		return nil
	}
	ni := int(n)

	b := &boolgen{}
	s := make([]int, ni)
	for i := 0; i < ni; i++ {
		s[i] = rand.Intn(max)
		if b.Bool() {
			s[i] *= -1
		}
	}
	return s
}

type boolgen struct {
	cache     int64
	remaining int
}

func (b *boolgen) Bool() bool {
	if b.remaining == 0 {
		b.cache, b.remaining = rand.Int63(), 63
	}

	result := b.cache&0x01 == 1
	b.cache >>= 1
	b.remaining--

	return result
}
