package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
	"strconv"

	"github.com/Shopify/sarama"
)

func setKafkaPayload() []byte  {
 metadata := map[string]interface{} {
                "serverAddr":"http://192.168.204.128:8200",
                "secretPath":"secret/data/cncp/orchestration/01EYVC0B46VHF8TGGJNNZF4KF6",
        }
        event := map[string]interface{}{
           "correlationId":"01EY97B2HDJ4E065NWSX092S98",
           "messageId":"01EYVC0B46S11S8V1Y54XNGFMK",
           "workloadId":"01EY97B2HDJ4E065NWSX092S98",
           "clusterId":"01EYVC0B46VHF8TGGJNNZF4KF6",
           "platform":"OnPrem",
           "credentialsStore":"HashiCorp Vault",
           "credentialsMetadata": metadata,
           "status":"Complete",
           "statusMessage":"ghjghjgjhg",
        }
        message, err := json.Marshal(event)
        if err != nil {
                panic(err)  
        }
	return message
}

func main() {

	// Setup configuration
	topic := os.Args[1]
	config := sarama.NewConfig()
	// Return specifies what channels will be populated.
	// If they are set to true, you must read from
	// config.Producer.Return.Successes = true
	// The total number of times to retry sending a message (default 3).
	config.Producer.Retry.Max = 5
	// The level of acknowledgement reliability needed from the broker.
	config.Producer.RequiredAcks = sarama.WaitForAll
	brokers := []string{"192.168.204.128:9092"}
	producer, err := sarama.NewAsyncProducer(brokers, config)
	if err != nil {
		// Should not reach here
		panic(err)
	}

	defer func() {
		if err := producer.Close(); err != nil {
			// Should not reach here
			panic(err)
		}
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	var enqueued, errors int
	doneCh := make(chan struct{})
	go func() {
		//for {

			time.Sleep(500 * time.Millisecond)

			strTime := strconv.Itoa(int(time.Now().Unix()))
			payload := setKafkaPayload()
			//fmt.Println(payload)
			msg := &sarama.ProducerMessage{
				Topic: topic,
				Key:   sarama.StringEncoder(strTime),
				Value: sarama.StringEncoder(payload),
			}
			select {
			case producer.Input() <- msg:
				enqueued++
				fmt.Println("Produce message")
			case err := <-producer.Errors():
				errors++
				fmt.Println("Failed to produce message:", err)
			case <-signals:
				doneCh <- struct{}{}
			}
		//}
	}()

	<-doneCh
	log.Printf("Enqueued: %d; errors: %d\n", enqueued, errors)
}
