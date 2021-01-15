package main

import (
	"github.com/sirupsen/logrus"
	"kafka_demo/internal/common/log"
	"kafka_demo/internal/producer"
	"time"
)

func init() {
	log.Init()
}

const TimeLayout = "2006-01-02 15:04:05"

func main() {

	for{
		nowTime := time.Now().Format(TimeLayout)
		if err := new(producer.Service).ProduceMsg(nowTime); err != nil {
			logrus.Errorf("发布消息失败，错误：%v", err.Error())
		}

		time.Sleep(time.Second)
	}

}
