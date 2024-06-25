package main

import (
	"database/sql"
	"encoding/json"
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/marcioecom/payment/adapter/broker/kafka"
	"github.com/marcioecom/payment/adapter/factory"
	"github.com/marcioecom/payment/adapter/presenter/transaction"
	"github.com/marcioecom/payment/usecase/processtx"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatalf("sql open error: %v", err)
	}

	repositoryFactory := factory.NewRepositoryDatabaseFactory(db)
	repository := repositoryFactory.CreateTransactionRepository()

	configMapProducer := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
	}
	kafkaPresenter := transaction.NewKafkaPresenter()
	producer := kafka.NewProducer(configMapProducer, kafkaPresenter)

	msgChan := make(chan *ckafka.Message)

	configMapConsumer := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"client.id":         "goapp",
		"group.id":          "goapp",
	}
	topics := []string{"transactions"}
	consumer := kafka.NewConsumer(configMapConsumer, topics)

	go func() {
		if err := consumer.Consume(msgChan); err != nil {
			log.Default().Printf("consume error: %v", err)
		}
	}()

	usecase := processtx.NewProcessTransaction(repository, producer, "transactions-result")

	log.Default().Println("starting looking for messages")

	for msg := range msgChan {
		var input processtx.TransactionDtoInput
		if err := json.Unmarshal(msg.Value, &input); err != nil {
			log.Default().Printf("unmarshal error: %v", err)
			continue
		}

		output, err := usecase.Execute(input)
		if err != nil {
			log.Default().Printf("usecase error: %v", err)
			continue
		}

		log.Default().Printf("output: %+v", output)
	}
}
