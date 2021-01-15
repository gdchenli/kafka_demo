package kafka

import (
	cluster "github.com/bsm/sarama-cluster"
)

type ConsumerConfig struct {
	Brokers []string
	GroupId string
	Topic   []string
}

type kafkaConsumerInterface interface {
	HandleMsg([]byte)
}

func Consume(consumer *ConsumerConfig, kafkaConsumer kafkaConsumerInterface) {

	config := cluster.NewConfig()
	config.Group.Mode = cluster.ConsumerModePartitions
	c, err := cluster.NewConsumer(consumer.Brokers, consumer.GroupId, consumer.Topic, config)
	if err != nil {
		return
	}

	defer c.Close()

	for {
		select {
		case part, ok := <-c.Partitions():
			if !ok {
				return
			}
			// start a separate goroutine to consume messages
			go func(pc cluster.PartitionConsumer) {
				for msg := range pc.Messages() {
					//fmt.Fprintf(os.Stdout, "Topic:%s \n Partition: %d  \n Offset: %d \n Value:%s\n ", msg.Topic, msg.Partition, msg.Offset, string(msg.Value))
					kafkaConsumer.HandleMsg(msg.Value)
					c.MarkOffset(msg, "") // mark message as processed
				}
			}(part)
		}
	}
}
