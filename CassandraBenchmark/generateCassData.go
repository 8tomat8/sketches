package main

import (
	"math/rand"

	"sync"

	"fmt"

	"flag"

	"github.com/gocql/gocql"
)

var goNum = flag.Int("num", 10, "asdasd")

func main() {
	var c = gocql.NewCluster("127.0.0.1:9042")
	c.Keyspace = "db"
	c.ProtoVersion = 4
	c.Consistency = gocql.Any

	var s, err = c.CreateSession()

	if err != nil {
		panic(err)
	}

	query := s.Query(`INSERT INTO table ("ID", "Field1", "Field2", "Field3") VALUES ( ?, ?, ?, ?)`)

	ID := gocql.TimeUUID()

	topics := []string{}

	for i := 0; i < 1000; i++ {
		topics = append(topics, RandStringRunes(8))
	}

	sem := make(semaphore, *goNum)
	for i := 0; i < *goNum; i++ {
		sem <- empty{}
	}

	wg := sync.WaitGroup{}
	wg.Add(1000000)

	for i := 0; i < 1000000; i++ {
		<-sem
		go func() {
			err := query.Bind(ID, "Data1", gocql.TimeUUID(), topics[rand.Intn(1000)]).Exec()
			if err != nil {
				panic(err)
			}
			if i%1000 == 0 {
				fmt.Println(i, query.String())
			}
			sem <- empty{}
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("DATA WAS GENERATED!!!")
}
