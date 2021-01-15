package producer

import (
	"errors"
	"kafka_demo/internal/common/config"
	"kafka_demo/pkg/kafka"
	"strings"
)

const (
	Brokers  = "kafka_demo_producer.brokers"
	Topic    = "kafka_demo_producer.topic"
	Version  = "kafka_demo_producer.version"
	ClientId = "kafka_demo_producer.client_id"
)

const (
	RequireClientId = "kafka配置错误，缺少Client id"
	RequiredBrokers = "kafka配置错误，缺少Brokers"
	RequiredTopic   = "kafka配置错误，缺少Topic"
	RequiredMsg     = "缺少需要发布的消息"
)

type Service struct{}

func (service *Service) ProduceMsg(msg string) (err error) {
	if msg == "" {
		return errors.New(RequiredMsg)
	}
	clientId := config.GetInstance().GetString(ClientId)
	if clientId == "" {
		return errors.New(RequireClientId)
	}
	brokers := strings.Split(config.GetInstance().GetString(Brokers), ",")
	if len(brokers) == 0 {
		return errors.New(RequiredBrokers)
	}
	topic := config.GetInstance().GetString(Topic)
	if topic == "" {
		return errors.New(RequiredTopic)
	}
	version := config.GetInstance().GetString(Version)

	producer, err := kafka.NewProducer(&kafka.KafkaProducerConfig{
		ClientID: clientId,
		Brokers:  brokers,
		Topic:    topic,
		Version:  version,
	})
	if err != nil {
		return err
	}

	if err = producer.Write([]byte(msg)); err != nil {
		return err
	}

	return nil
}
