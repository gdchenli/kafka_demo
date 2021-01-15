package kafka

var MessagingClient MessagingProducer

type MessagingProducer interface {
	Connect() error
	Close() error
	Description() string

	SampleConfig() string

	Write(output Message) error
}

// Message contains the fields of the kafka message
type Message interface {
	Value() []byte
}
