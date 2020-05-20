package main

import (
	"fmt"
	"kafka_demo/internal/producer"
)

func main() {
	msg := `{"website_id":33,"data":{"website_id":33,"store_id":150,"increment_id":"20200518192100001","three_part_order_number":"","carrier_code":"ucs","track_number":"9975660274846","created_at":"2020-04-23 08:36:31"}}`
	if err := new(producer.Service).ProduceMsg(msg); err != nil {
		fmt.Println(err)
	}
}
