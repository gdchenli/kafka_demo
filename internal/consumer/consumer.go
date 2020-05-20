package consumer

import (
	"errors"
	"kafka_demo/internal/common/config"
	"kafka_demo/pkg/kafka"
	"strings"

	"github.com/sirupsen/logrus"
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
		logrus.Errorf("kafka消费消息失败，错误：%v", err.Error())
		return
	}

	kafka.Consume(kafkaConsumerConfig, service)
}

func (service *Service) getConsumerConfig() (kafkaConsumerConfig *kafka.ConsumerConfig, err error) {
	brokers := config.GetInstance().GetString(Brokers)
	if brokers == "" {
		return kafkaConsumerConfig, errors.New("kafka_demo_consumer brokers error")
	}
	groupId := config.GetInstance().GetString(GroupId)
	if groupId == "" {
		return kafkaConsumerConfig, errors.New("kafka_demo_consumer group id error")
	}
	topics := config.GetInstance().GetString(Topic)
	if topics == "" {
		return kafkaConsumerConfig, errors.New("kafka_demo_consumer topic error")
	}

	kafkaConsumerConfig = &kafka.ConsumerConfig{
		Brokers: strings.Split(brokers, ","),
		GroupId: groupId,
		Topic:   strings.Split(topics, ","),
	}

	return kafkaConsumerConfig, nil
}

func (service *Service) HandleMsg(msgBytes []byte) {
	msg := string(msgBytes)
	logrus.Infof("消费kafka消息:%v", msg)
}
