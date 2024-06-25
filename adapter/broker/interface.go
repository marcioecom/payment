package broker

type ProducerInterface interface {
	Publish(msg any, key []byte, topic string) error
}
