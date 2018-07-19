package main

import (
	"math/rand"

	"sync"

	"fmt"

	"time"

	"sync/atomic"

	"github.com/gocql/gocql"
)

const (
	runs      = 1000000
	goNumRead = 100
)

func main() {
	var c = gocql.NewCluster("127.0.0.1:9042")
	c.Keyspace = "db"
	c.ProtoVersion = 4
	c.Consistency = gocql.One

	var s, err = c.CreateSession()

	if err != nil {
		panic(err)
	}

	query := s.Query(`SELECT "ID" FROM table WHERE "ID" = ? PER PARTITION LIMIT 1`)

	var IDs []gocql.UUID

	for i := 0; i < 1000; i++ {
		IDs = append(IDs, gocql.TimeUUID())
	}

	sem := make(semaphore, goNumRead)
	for i := 0; i < goNumRead; i++ {
		sem <- empty{}
	}

	wg := sync.WaitGroup{}
	wg.Add(runs)

	var (
		foundCounter    int64
		notFoundCounter int64
		i               int
	)

	go func() {
		ticker := time.NewTicker(time.Second)
		for {
			select {
			case <-ticker.C:
				fmt.Println("Sum:", foundCounter+notFoundCounter, "Found:", foundCounter, "Not found:", notFoundCounter, "#:", i)
				atomic.StoreInt64(&foundCounter, 0)
				atomic.StoreInt64(&notFoundCounter, 0)
			}
		}
	}()

	for i = 0; i < runs; i++ {
		i := i
		<-sem
		go func() {
			defer func() {
				sem <- empty{}
				wg.Done()
			}()

			var (
				ID gocql.UUID
			)
			t := IDs[rand.Intn(1000)]
			err := query.Bind(t, IDs[i]).Scan(&ID)
			if err == gocql.ErrNotFound {
				notFoundCounter = atomic.AddInt64(&notFoundCounter, 1)
			} else if err != nil {
				panic(err)
			}
			if err != gocql.ErrNotFound && (IDs[i] != ID) {
				fmt.Println("WRONG!!! ", ID)
			}
			foundCounter = atomic.AddInt64(&foundCounter, 1)
		}()
	}
	wg.Wait()

	fmt.Println("The End...")
}
