package main

import (
	"log"
	"testing"

	"context"
	"sync"

	"github.com/Shopify/sarama"
	"github.com/bsm/sarama-cluster"
)

func BenchmarkConsumeMsgs(b *testing.B) {
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumer, err := cluster.NewConsumer([]string{"kafka:9092"}, "g7", []string{*topic}, config)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()
	ctx := context.Background()
	wg := &sync.WaitGroup{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		consumeMsgs(ctx, wg, 1000000, consumer)
	}
}
