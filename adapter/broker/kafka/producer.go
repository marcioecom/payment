package kafka

import (
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/marcioecom/payment/adapter/presenter"
)

type Producer struct {
	configMap *ckafka.ConfigMap
	presenter presenter.Presenter
}

func NewProducer(configMap *ckafka.ConfigMap, presenter presenter.Presenter) *Producer {
	return &Producer{
		configMap: configMap,
		presenter: presenter,
	}
}

func (p *Producer) Publish(msg any, key []byte, topic string) error {
	producer, err := ckafka.NewProducer(p.configMap)
	if err != nil {
		return err
	}

	if err = p.presenter.Bind(msg); err != nil {
		return err
	}

	presenterMsg, err := p.presenter.Show()
	if err != nil {
		return err
	}

	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{
			Topic: &topic, Partition: ckafka.PartitionAny,
		},
		Value: presenterMsg,
		Key:   key,
	}

	if err = producer.Produce(message, nil); err != nil {
		return err
	}

	return nil
}
