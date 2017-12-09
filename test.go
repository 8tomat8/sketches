package main

import (
	"fmt"
	"time"
)

type Person struct {
	Name string
}

func main() {
	t := time.Now()
	fmt.Println(t)

	loc, err := time.LoadLocation("GMT")
	if err != nil {
		panic(err)
	}
	t = t.In(loc)

	fmt.Println(t)
	t1 := time.Now().In(loc)
	fmt.Println(t1)
	fmt.Println(t.Before(t1))
}
