package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"flag"

	"context"

	"sync"

	"github.com/Shopify/sarama"
)

var (
	numToConsume = flag.Int("get", 3, "Num of messages to consume")
	topic        = flag.String("topic", "testtopic4part", "Kafka topic name")
	group        = flag.String("group", "g2", "Kafka group name")
)

func main() {
	flag.Parse()
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumer, err := cluster.NewConsumer([]string{"kafka:9092"}, *group, []string{*topic}, config)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go consumeMsgs(ctx, wg, *numToConsume, consumer)

	<-signals
	cancel()
	wg.Wait()
}

func consumeMsgs(ctx context.Context, wg *sync.WaitGroup, numToConsume int, consumer *cluster.Consumer) {
	consumed := 0
	defer func() {
		wg.Done()
	}()

	for {
		select {
		case msg := <-consumer.Messages():
			consumed++

			if consumed%1000 == 0 {
				consumer.CommitOffsets()
				consumer.MarkPartitionOffset(msg.Topic, msg.Partition, msg.Offset, "")
			}

			if consumed >= numToConsume {
				return
			}
		case e := <-consumer.Errors():
			fmt.Println(e)
		case <-ctx.Done():
			return
		}
	}
}
