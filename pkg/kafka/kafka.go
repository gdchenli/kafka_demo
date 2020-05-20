package kafka

import (
	"github.com/Shopify/sarama"
)

type Kafka struct {
	Brokers      []string
	Topic        string
	ClientID     string
	Version      string
	RequiredAcks int
	MaxRetry     int
	producer     sarama.SyncProducer
}

type KafkaProducerConfig struct {
	ClientID string   `mapstructure:"client_id"`
	Brokers  []string `mapstructure:"kafka_brokers"`
	Topic    string   `mapstructure:"kafka_topic"`
	Version  string   `mapstructure:"version"`
}

var sampleConfig = `
  ## URLs of kafka brokers
  brokers = ["localhost:9092"]
  ## Kafka topic for producer messages
  topic = "inventory_update"

  ## Optional Client id
  # client_id = "inventory_service"

  ## Set the minimal supported Kafka version.  Setting this enables the use of new
  ## Kafka features and APIs.  Of particular interest, lz4 compression
  ## requires at least version 0.10.0.0.
  ##   ex: version = "1.1.0"
  # version = ""
  ## CompressionCodec represents the various compression codecs recognized by
  ## Kafka in messages.
  ##  0 : No compression
  ##  1 : Gzip compression
  ##  2 : Snappy compression
  ##  3 : LZ4 compression
  # compression_codec = 0

  ##  RequiredAcks is used in Produce Requests to tell the broker how many
  ##  replica acknowledgements it must see before responding
  ##   0 : the producer never waits for an acknowledgement from the broker.
  ##       This option provides the lowest latency but the weakest durability
  ##       guarantees (some data will be lost when a server fails).
  ##   1 : the producer gets an acknowledgement after the leader replica has
  ##       received the data. This option provides better durability as the
  ##       client waits until the server acknowledges the request as successful
  ##       (only messages that were written to the now-dead leader but not yet
  ##       replicated will be lost).
  ##   -1: the producer gets an acknowledgement after all in-sync replicas have
  ##       received the data. This option provides the best durability, we
  ##       guarantee that no messages will be lost as long as at least one in
  ##       sync replica remains.
  # required_acks = -1

  ## The maximum number of times to retry sending a metric before failing
  ## until the next flush.
  # max_retry = 3

  ## The maximum permitted size of a message. Should be set equal to or
  ## smaller than the broker's 'message.max.bytes'.
  # max_message_bytes = 1000000
`

func NewProducer(config *KafkaProducerConfig) (*Kafka, error) {

	kafka := &Kafka{
		Brokers:  config.Brokers,
		Topic:    config.Topic,
		ClientID: config.ClientID,
		Version:  config.Version,
	}

	err := kafka.Connect()
	if err != nil {
		return nil, err
	}

	return kafka, nil

}
func (k *Kafka) GetTopicName() string {
	var topicName string

	topicName = k.Topic

	return topicName
}

func (k *Kafka) SampleConfig() string {
	return sampleConfig
}
func (k *Kafka) Connect() error {

	config := sarama.NewConfig()

	if k.Version != "" {
		version, err := sarama.ParseKafkaVersion(k.Version)
		if err != nil {
			return err
		}
		config.Version = version
	}

	if k.ClientID != "" {
		config.ClientID = k.ClientID
	} else {
		config.ClientID = "inventory_service"
	}

	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(k.Brokers, config)
	if err != nil {
		return err
	}
	k.producer = producer
	return nil
}

func (k *Kafka) Close() error {
	return k.producer.Close()
}

func (k *Kafka) Description() string {
	return "Configuration for the Kafka server to send inventory update to "
}

func (k *Kafka) Write(output []byte) error {
	msg := &sarama.ProducerMessage{
		Topic: k.Topic,
		Value: sarama.StringEncoder(output),
	}
	_, _, err := k.producer.SendMessage(msg)

	if err != nil {
		return err
	}
	return nil
}
