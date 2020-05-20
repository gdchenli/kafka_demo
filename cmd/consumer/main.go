package main

import "kafka_demo/internal/consumer"

func main() {
	new(consumer.Service).Consumer()
}
