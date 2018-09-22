package bubble

import (
	"math"
	"testing"

	"github.com/8tomat8/sketches/algorithms/sorting"
)

func TestSort(t *testing.T) {
	testcases := []struct {
		n   uint
		max int
	}{
		{10, 10},
		{0, math.MaxInt64},
		{1, math.MaxInt64},
		{100, math.MaxInt64},
		{2000, math.MaxInt64},
	}

	for _, tt := range testcases {
		s := sorting.Gen(tt.n, tt.max)
		l := len(s)
		Sort(s)

		if len(s) != l {
			t.Fatalf("Expected %d values in result slice, got %d", l, len(s))
		}

		for i, v := range s {
			if i == 0 {
				continue
			}

			if v < s[i-1] {
				t.Log(s)
				t.Fatalf("value on the left is bigger (left: %d, right: %d)", s[i-1], v)
			}
		}
	}
}
