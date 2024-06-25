package kafka

import (
	"testing"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/marcioecom/payment/adapter/presenter/transaction"
	"github.com/marcioecom/payment/domain/entity"
	"github.com/marcioecom/payment/usecase/processtx"
	"github.com/stretchr/testify/assert"
)

func TestProducerPublish(t *testing.T) {
	expectedOutput := processtx.TransactionDtoOutput{
		ID:           "1",
		Status:       entity.REJECTED,
		ErrorMessage: "you don't have limit for this transaction",
	}

	configMap := &ckafka.ConfigMap{
		"test.mock.num.brokers": 3,
	}
	producer := NewProducer(configMap, transaction.NewKafkaPresenter())
	err := producer.Publish(expectedOutput, []byte("1"), "test")
	assert.Nil(t, err)
}
