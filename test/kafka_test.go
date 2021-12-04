package test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/Shopify/sarama"
)

func TestKafka(t *testing.T) {
	hosts := "myAddress:9092"
	config := sarama.NewConfig()
	client, _ := sarama.NewClient(strings.Split(hosts, ","), config)

	producer, _ := sarama.NewAsyncProducerFromClient(client)
	defer func() {
		if producer != nil {
			producer.Close()
		}
		if client != nil {
			client.Close()
		}
	}()

	var str string = "test"
	var data []byte = []byte(str)
	be := sarama.ByteEncoder(data)
	fmt.Println(be)
	producer.Input() <- &sarama.ProducerMessage{Topic: "test", Key: nil, Value: be}

}

var consumer sarama.Consumer

type ConsumerCallBack func(data []byte)

func TestKafkaConsumer(t *testing.T) {
	hosts := "myAddress:9092"
	config := sarama.NewConfig()
	client, _ := sarama.NewClient(strings.Split(hosts, ","), config)

	consumer, _ = sarama.NewConsumerFromClient(client)

	partitionConsumer, _ := consumer.ConsumePartition("test", 0, sarama.OffsetNewest)
	defer partitionConsumer.Close()
	for {
		msg := <-partitionConsumer.Messages()
		fmt.Println(msg.Value)
		fmt.Println(string(msg.Value))
	}
}
