package main

import (
	"fmt"
	"time"
)

func main() {
	ch1, ch2 := make(chan string), make(chan string)
	done := make(chan struct{})

	go x()
	go f(ch1, ch2, done)
	go f(ch2, ch1, done)
	<-done
	<-done
}

func f(send chan<- string, receive <-chan string, done chan<- struct{}) {
	send <- "Some msg"
	<-receive
	done <- struct{}{}
}

func x() {
	for {
		time.Sleep(time.Second)
		fmt.Println("alive")
	}
}
