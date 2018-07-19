package main

import (
	"fmt"
	"io"
	"os"
)

type uds struct {
	a string
	b int
	c interface{}
}

func main() {
	testChan := make(chan string)

	fmt.Println("nil:", test(nil))
	fmt.Println("string:", test("string"))
	fmt.Println("123:", test(123))
	fmt.Println("0.123:", test(0.123))
	fmt.Println("chan string:", test(testChan))
	fmt.Println("custom struct:", test(uds{}))
	fmt.Println("[]int:", test([]int{1, 2, 3, 4}))
	fmt.Println("io.ReadWriteCloser:", test(os.File{}))
}

func test(v interface{}) string {
	rwc, ok := v.(io.ReadWriteCloser)
	if !ok {
		return "Oops!"
	}
	_ = rwc
	return "Asserted!"
}

type User struct {
	name string
	age  int
	legs []int
}

func (u User) Name() string {
	u.name = "AAAA"
	return
}

func (u *User) Age() int {

}
