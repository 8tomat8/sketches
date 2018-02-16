package main

import (
	"context"
	"net"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"math/rand"

	"fmt"

	"github.com/sirupsen/logrus"
)

var counter int64

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})

	l, err := net.Listen("tcp", ":5000")
	if err != nil {
		logrus.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)

	stopChan := make(chan struct{})
	go server(ctx, stopChan, l)

	logrus.Info("Server started")
	<-ctx.Done()
	logrus.Info("Timeout")
	l.Close()
	logrus.Info("Listener closed")
	<-stopChan
	logrus.Info("Server stoped")
}

func server(ctx context.Context, stopChan chan struct{}, l net.Listener) {
	defer close(stopChan)

	taskChan, closeFunc := newWorkerPool(3)

	for {
		select {
		case <-ctx.Done():
			logrus.Info("Server stop, waiting for handlers")
			closeFunc()
			return
		default:
			c, err := l.Accept()
			if err != nil {
				continue
			}
			taskChan <- c
		}
	}
}

func newWorkerPool(num int) (chan net.Conn, func()) {
	workerWG := &sync.WaitGroup{}
	taskChan := make(chan net.Conn)
	for i := 0; i < num; i++ {
		workerWG.Add(1)
		go worker(i, workerWG, taskChan)
	}
	return taskChan, func() { close(taskChan); workerWG.Wait() }
}

func worker(i int, workerWG *sync.WaitGroup, connects chan net.Conn) error {
	defer func() {
		fmt.Println("Worker", i, "received stop")
		workerWG.Done()
	}()

	connIndex := make([]byte, 1)
	for connect := range connects {
		_, err := connect.Read(connIndex)
		if err != nil {
			logrus.Warn(err)
			continue
		}
		logrus.Infof("worker %d picked up connect %d", i, uint8(connIndex[0]))

		counter = atomic.AddInt64(&counter, 1)
		fmt.Fprint(connect, "ok"+strconv.Itoa(int(counter)))
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))

		err = connect.Close()
		if err != nil {
			logrus.Error(err)
		}
	}
	return nil
}
