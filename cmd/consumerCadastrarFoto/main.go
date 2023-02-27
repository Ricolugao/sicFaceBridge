package main

import (
	"fmt"
	"sicFaceBridge/controllers"
	"sicFaceBridge/env"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	env.CarregaVariaveisDeAmbiente()
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"client.id":         "sicFaceBridge",
		"group.id":          "cadastrarFoto",
		// "auto.offset.reset": "earliest",
	}
	c, err := kafka.NewConsumer(configMap)
	if err != nil {
		fmt.Println("error consumer", err.Error())
	}
	topics := []string{"teste"}
	c.SubscribeTopics(topics, nil)
	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {

			controllers.CadastraFotoCompreFace(msg.Value)
			// fmt.Println(mensagem[0], msg.TopicPartition)
		}
	}
}
