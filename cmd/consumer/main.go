package main

import (
	"kafka_demo/internal/common/log"
	"kafka_demo/internal/consumer"
)

func init() {
	log.Init()
}

func main() {
	new(consumer.Service).Consumer()
}
