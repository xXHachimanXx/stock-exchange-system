package main

import (
	"encoding/json"
	"fmt"
	"sync"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/xXHachimanXx/stock-exchange-system/go/internal/infra/kafka"
	"github.com/xXHachimanXx/stock-exchange-system/go/internal/market/dto"
	"github.com/xXHachimanXx/stock-exchange-system/go/internal/market/entity"
	"github.com/xXHachimanXx/stock-exchange-system/go/internal/market/transformer"
)

func main() {
	ordersIn := make(chan *entity.Order)
	ordersOut := make(chan *entity.Order)
	wg := &sync.WaitGroup{}
	defer wg.Wait()

	kafkaMsgChan := make(chan *ckafka.Message)

	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": "host.docker.internal:9094",
		"group.id":          "myGroup",
		"auto.offset.reset": "latest",
	}

	kafkaProducer := kafka.NewKafkaProducer(configMap)
	kafkaConsumer := kafka.NewKafkaConsumer(configMap, []string{"input"})

	go kafkaConsumer.Consume(kafkaMsgChan) // Create new thread

	book := entity.NewBook(ordersIn, ordersOut, wg)
	go book.Trade()

	go func() {
		for msg := range kafkaMsgChan {
			wg.Add(1)
			fmt.Println(string(msg.Value))
			tradeInput := dto.TradeInputDTO{}
			err := json.Unmarshal(msg.Value, &tradeInput)
			if err != nil {
				panic(err)
			}
			order := transformer.TransformInput(tradeInput)
			ordersIn <- order
		}
	}()

	for res := range ordersOut {
		output := transformer.TransformOutput(res)
		outputJson, err := json.MarshalIndent(output, "", "  ")
		fmt.Println(string(outputJson))
		if err != nil {
			fmt.Println(err)
		}
		kafkaProducer.Publish(outputJson, []byte("orders"), "output")
	}
}
