package main

import (
	"fmt"
	"testing"
)

func TestTest1(t *testing.T) {

	t.Parallel()
	for i := 0; i < 10; i += 2 {
		i := i
		t.Run("Test1", func(t *testing.T) {
			t.Parallel()
			test(i)
			fmt.Println("Test1", i)
		})
	}
}

func TestTest2(t *testing.T) {

	t.Parallel()
	for i := 1; i < 10; i += 2 {
		i := i
		t.Run("Test2", func(t *testing.T) {
			t.Parallel()
			test(i)
			fmt.Println("Test2", i)
		})
	}
}
