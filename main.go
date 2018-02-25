package main

import (
	"runtime"
	"strconv"
)

var a string
var done bool

func setup() {
	a = "hello, world"
	done = true
}

func main() {
	runtime.GOMAXPROCS(1)
	go setup()
	var i int
	for ; !done; i++ {
		//runtime.Gosched() // Uncomment to yield main goroutine
	}
	print(strconv.Itoa(i))
	print(a)
}
