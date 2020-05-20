package consumer

import (
	"errors"
	"fmt"
	"kafka_demo/internal/config"
	"kafka_demo/pkg/kafka"
	"strings"
)

const (
	Brokers = "kafka_demo_consumer.brokers"
	Topic   = "kafka_demo_consumer.topic"
	GroupId = "kafka_demo_consumer.group_id"
)

type Service struct{}

func (service *Service) Consumer() {
	kafkaConsumerConfig, err := service.getConsumerConfig()
	if err != nil {
		return
	}

	kafka.Consume(kafkaConsumerConfig, service)
}

func (service *Service) getConsumerConfig() (kafkaConsumerConfig *kafka.ConsumerConfig, err error) {
	brokers := config.GetInstance().GetString(Brokers)
	fmt.Printf("brokers %+v\n", brokers)
	if brokers == "" {
		return kafkaConsumerConfig, errors.New("kafka_demo_consumer brokers error")
	}

	groupId := config.GetInstance().GetString(GroupId)
	fmt.Printf("groupId %+v\n", groupId)
	if groupId == "" {
		return kafkaConsumerConfig, errors.New("kafka_demo_consumer group id error")
	}

	topics := config.GetInstance().GetString(Topic)
	fmt.Printf("topics %+v\n", topics)
	if topics == "" {
		return kafkaConsumerConfig, errors.New("kafka_demo_consumer topic error")
	}

	kafkaConsumerConfig = &kafka.ConsumerConfig{
		Brokers: strings.Split(brokers, ","),
		GroupId: groupId,
		Topic:   strings.Split(topics, ","),
	}

	fmt.Printf("kafkaConsumerConfig %+v\n", kafkaConsumerConfig)

	return kafkaConsumerConfig, nil
}

func (service *Service) HandleMsg(msgBytes []byte) {
	fmt.Println("kafka msg:", string(msgBytes))
}
