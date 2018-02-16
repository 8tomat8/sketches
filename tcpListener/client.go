package main

import (
	"log"
	"net"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})

	wg := &sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go client(wg, i)
	}

	wg.Wait()
}

func client(wg *sync.WaitGroup, i int) {
	defer func() {
		wg.Done()
		logrus.Info("Client stoped", i)
	}()

	logrus.Info("Client started", i)
	conn, err := net.Dial("tcp", "127.0.0.1:5000")
	if err != nil {
		log.Fatal(err)
		return
	}

	conn.Write([]byte{uint8(i)})

	data := make([]byte, 2<<10)
	for {
		n, err := conn.Read(data)
		if err != nil {
			logrus.Info("Cli error", i, ":", err)
			break
		}
		time.Sleep(time.Millisecond * 1500)
		logrus.Info(string(data[:n]))
	}
}
