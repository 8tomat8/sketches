package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	go handleSignals(cancel)

Loop:
	for i := 0; ; i++ {
		select {
		case <-time.After(time.Second / 2):
			do(i)
		case <-ctx.Done():
			fmt.Println()
			break Loop
		}
	}

	fmt.Println("doing something before exit")
}

func handleSignals(cancel context.CancelFunc) {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Waiting for first signal
	<-sigs

	cancel()
}

func do(i int) {
	time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
	fmt.Println(i)
}
