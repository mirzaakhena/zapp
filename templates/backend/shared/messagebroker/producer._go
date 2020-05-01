package messagebroker

import nsq "github.com/nsqio/go-nsq"

// IMessageBrokerProducer is
type IMessageBrokerProducer interface {
	Publish(topic string, body []byte) error
}

// Producer is
type Producer struct {
	producer *nsq.Producer
}

// NewProducer is
func NewProducer(url string) *Producer {

	obj := new(Producer)

	nsqConfig := nsq.NewConfig()
	producer, err := nsq.NewProducer(url, nsqConfig)
	if err != nil {
		panic(err.Error())
	}

	obj.producer = producer

	return obj
}

// Publish is
func (m *Producer) Publish(topic string, body []byte) error {
	return m.producer.Publish(topic, body)
}
