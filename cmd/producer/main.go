package main

import (
	"kafka_demo/internal/common/log"
	"kafka_demo/internal/producer"
	"time"

	"github.com/sirupsen/logrus"
)

func init() {
	log.Init()
}

const TimeLayout = "2006-01-02 15:04:05"

func main() {
	nowTime := time.Now().Format(TimeLayout)
	msg := "{\"data\":\"this is a message\",\"time\":\"" + nowTime + "\"}"
	if err := new(producer.Service).ProduceMsg(msg); err != nil {
		logrus.Errorf("发布消息失败，错误：%v", err.Error())
	}
}
