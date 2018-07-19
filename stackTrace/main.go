package main

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
)

func main() {
	fmt.Println(testing.AllocsPerRun(20, func() { _ = f1() }))
}

func f1() error {
	return f2()
}

func f2() error {
	return errors.WithStack(errors.New("AAA"))
}
