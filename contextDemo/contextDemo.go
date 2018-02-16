// Run this application.
// Try to change behavior and look what it will give you.
// Also you may catch deadlock in this application. Try to fix it by your own.
// Valid code without bug is in presentation.
package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var GlobalResource chan int

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	GlobalResource = NewResource(ctx, wg)

	wg.Add(1)
	go Worker(ctx, wg)

	time.Sleep(time.Second * 1)
	cancel()
	wg.Wait()
	fmt.Println("Main goroutine stoping")
}

func Worker(ctx context.Context, wg *sync.WaitGroup) {
	defer func() {
		fmt.Println("Worker was stoped")
		wg.Done()
	}()
	fmt.Println("Worker was started")

	childWG := &sync.WaitGroup{}

	for i := 0; i < 3; i++ {
		childWG.Add(1)
		go SubWorker(ctx, childWG, i)
	}

	// Will wait until main would not call cancel()
	<-ctx.Done()
	fmt.Println("Worker receiver ctx.Done signal. Waiting for SubWorkers")
	childWG.Wait()
}

func SubWorker(ctx context.Context, wg *sync.WaitGroup, index int) {
	defer func() {
		fmt.Printf("SubWorker %d was stoped\n", index)
		wg.Done()
	}()
	fmt.Printf("SubWorker #%d was started\n", index)

	for {
		select {
		case <-ctx.Done():
			sleepTime := rand.Intn(200)
			fmt.Printf("SubWorker #%d Received ctx.Done signal. Stoping in %d Milliseconds\n", index, sleepTime)
			time.Sleep(time.Millisecond * time.Duration(sleepTime))
			return
		case <-time.After(time.Millisecond * 500):
			random, ok := <-GlobalResource
			if !ok {
				fmt.Printf("GlobalResource is closed. SubWorker #%d\n", index)
				return
			}
			fmt.Printf("Tick from SubWorker #%d with value %d\n", index, random)
		}
	}
}

func NewResource(ctx context.Context, wg *sync.WaitGroup) chan int {
	out := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				close(out)
				fmt.Println("Resource received ctx.Done signal")
				return
			case out <- rand.Intn(200):
			}
		}
	}()
	return out
}
