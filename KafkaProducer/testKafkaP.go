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

//var body = []byte(`{"Action":"performanceStorage","Topic":"performance","RegID":"800efff8-1a57-400c-9c28-d63b6d57217c","DcDateTimeUtc":"2017-05-30T09:54:12.087592268Z","Message":{"createTimeUTC":"2017-05-30T09:54:12.049328087Z","createdBy":"/continuum/agent/plugin/performance","name":"storage","type":"performanceStorages","storages":[{"createTimeUTC":"2017-05-30T09:54:12.049328087Z","createdBy":"/continuum/agent/plugin/performance","index":0,"name":"sda","type":"performanceStorage","metric":{"idleTime":0,"writeCompleted":28344315,"writeTimeMs":18081340,"readCompleted":263559,"readTimeMs":82020,"freeSpaceBytes":0,"usedSpaceBytes":0,"totalSpaceBytes":0,"diskTimeTotal":1.705108e+07},"partitions":[{"createTimeUTC":"2017-05-30T09:54:12.049328087Z","createdBy":"/continuum/agent/plugin/performance","index":0,"name":"sda1","type":"performanceStoragePartition","metric":{"idleTime":0,"writeCompleted":28146015,"writeTimeMs":17838076,"readCompleted":263424,"readTimeMs":81996,"freeSpaceBytes":0,"usedSpaceBytes":0,"totalSpaceBytes":0,"diskTimeTotal":1.6823836e+07}},{"createTimeUTC":"2017-05-30T09:54:12.049328087Z","createdBy":"/continuum/agent/plugin/performance","index":1,"name":"sda2","type":"performanceStoragePartition","metric":{"idleTime":0,"writeCompleted":0,"writeTimeMs":0,"readCompleted":2,"readTimeMs":4,"freeSpaceBytes":0,"usedSpaceBytes":0,"totalSpaceBytes":0,"diskTimeTotal":4}},{"createTimeUTC":"2017-05-30T09:54:12.049328087Z","createdBy":"/continuum/agent/plugin/performance","index":2,"name":"sda5","type":"performanceStoragePartition","metric":{"idleTime":0,"writeCompleted":23,"writeTimeMs":4,"readCompleted":90,"readTimeMs":16,"freeSpaceBytes":0,"usedSpaceBytes":0,"totalSpaceBytes":0,"diskTimeTotal":20}}]},{"createTimeUTC":"2017-05-30T09:54:12.049328087Z","createdBy":"/continuum/agent/plugin/performance","index":0,"name":"sdb","type":"performanceStorage","metric":{"idleTime":0,"writeCompleted":0,"writeTimeMs":0,"readCompleted":188,"readTimeMs":2232,"freeSpaceBytes":0,"usedSpaceBytes":0,"totalSpaceBytes":0,"diskTimeTotal":1232},"partitions":[{"createTimeUTC":"2017-05-30T09:54:12.049328087Z","createdBy":"/continuum/agent/plugin/performance","index":0,"name":"sdb2","type":"performanceStoragePartition","metric":{"idleTime":0,"writeCompleted":0,"writeTimeMs":0,"readCompleted":2,"readTimeMs":72,"freeSpaceBytes":0,"usedSpaceBytes":0,"totalSpaceBytes":0,"diskTimeTotal":72}},{"createTimeUTC":"2017-05-30T09:54:12.049328087Z","createdBy":"/continuum/agent/plugin/performance","index":1,"name":"sdb3","type":"performanceStoragePartition","metric":{"idleTime":0,"writeCompleted":0,"writeTimeMs":0,"readCompleted":46,"readTimeMs":808,"freeSpaceBytes":0,"usedSpaceBytes":0,"totalSpaceBytes":0,"diskTimeTotal":716}},{"createTimeUTC":"2017-05-30T09:54:12.049328087Z","createdBy":"/continuum/agent/plugin/performance","index":2,"name":"sdb5","type":"performanceStoragePartition","metric":{"idleTime":0,"writeCompleted":0,"writeTimeMs":0,"readCompleted":48,"readTimeMs":792,"freeSpaceBytes":0,"usedSpaceBytes":0,"totalSpaceBytes":0,"diskTimeTotal":712}}]},{"createTimeUTC":"2017-05-30T09:54:12.049328087Z","createdBy":"/continuum/agent/plugin/performance","index":0,"name":"sdc","type":"performanceStorage","metric":{"idleTime":0,"writeCompleted":0,"writeTimeMs":0,"readCompleted":0,"readTimeMs":0,"freeSpaceBytes":0,"usedSpaceBytes":0,"totalSpaceBytes":0,"diskTimeTotal":0},"partitions":null}]}}`)

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
