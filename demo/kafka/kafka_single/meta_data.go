package main

import (
	"fmt"
	"github.com/IBM/sarama"
)

// kafka高版本server能兼容低版本client

// MetadataTest
// 简单说, kafka集群的metadata包括:
// 所有broker的信息: ip和port;
// 所有topic的信息: topic name, partition数量, 每个partition的leader, isr, replica集合等
// kafka集群的每一台broker都缓存了整个集群的metadata, 当broker或某一个topic的metadata信息发生变化时,
// 集群的Controller 都会感知到作相应的状态转换, 同时把发生变化的新的metadata信息广播到所有的broker;
func MetadataTest() {
	fmt.Printf("metadata test\n")

	config := sarama.NewConfig()
	config.Version = sarama.V0_11_0_2

	client, err := sarama.NewClient([]string{"localhost:9092"}, config)
	if err != nil {
		fmt.Printf("metadata_test try create client err :%s\n", err.Error())
		return
	}

	defer client.Close()

	// get topic set
	topics, err := client.Topics()
	if err != nil {
		fmt.Printf("try get topics err %s\n", err.Error())
		return
	}

	fmt.Printf("topics(%d):\n", len(topics))

	for _, topic := range topics {
		fmt.Println(topic)
	}

	// get broker set
	brokers := client.Brokers()
	fmt.Printf("broker set(%d):\n", len(brokers))
	for _, broker := range brokers {
		fmt.Printf("%s\n", broker.Addr())
	}
}
