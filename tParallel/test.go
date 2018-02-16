package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	fmt.Println(test(1))
}

func test(i int) string {
	time.Sleep(time.Millisecond * 50)
	return strconv.Itoa(i)
}
