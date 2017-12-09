package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"

	"hash"

	"github.com/OneOfOne/xxhash"
	"github.com/Shopify/sarama"
)

var body = []byte(`msg2`)

var (
	numberOfMessages = flag.Int("msgs", 1000000, "Number of messages to produce.")
	topic            = flag.String("topic", "testtopic4part", "Kafka topic name")
)

func main() {
	flag.Parse()
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewCustomHashPartitioner(func() hash.Hash32 {
		return xxhash.New32()
	})

	producer, err := sarama.NewAsyncProducer([]string{"kafka:9092"}, config)
	if err != nil {
		panic(err)
	}

	// Trap SIGINT to trigger a graceful shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	var (
		wg                          sync.WaitGroup
		enqueued, successes, errors int
	)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for range producer.Successes() {
			successes++
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for err := range producer.Errors() {
			log.Println(err)
			errors++
		}
	}()

	start := time.Now()
	for i := 0; i < *numberOfMessages; i++ {
		message := &sarama.ProducerMessage{Topic: *topic, Value: sarama.ByteEncoder(body)}
		producer.Input() <- message
		enqueued++
	}
	producer.AsyncClose()
	wg.Wait()

	log.Printf("Successfully produced: %d; errors: %d\n", successes, errors)
	log.Println(time.Since(start))
}
