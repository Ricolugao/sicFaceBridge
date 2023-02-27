package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sicFaceBridge/controllers"
	"sicFaceBridge/env"
	"sicFaceBridge/model"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	env.CarregaVariaveisDeAmbiente()

	fotos := model.BuscaFotosParaCompreFace(100)

	deliveryChan := make(chan kafka.Event)
	producer := NewKafkaProducer()

	for _, foto := range fotos {

		mensagem, err := json.Marshal(foto)
		controllers.TrataErro(err)

		Publish(mensagem, "teste", producer, nil, deliveryChan)
		go DeliveryReport(deliveryChan) // async
	}

	e := <-deliveryChan
	msg := e.(*kafka.Message)
	if msg.TopicPartition.Error != nil {
		fmt.Println("Erro ao enviar")
	} else {
		fmt.Println("Mensagem enviada:", msg.TopicPartition)
	}
}

func NewKafkaProducer() *kafka.Producer {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers":   "kafka:9092",
		"delivery.timeout.ms": "0",
		"acks":                "1",
		"enable.idempotence":  "false",
	}
	p, err := kafka.NewProducer(configMap)
	if err != nil {
		log.Println(err.Error())
	}
	return p
}

func Publish(msg []byte, topic string, producer *kafka.Producer, key []byte, deliveryChan chan kafka.Event) error {
	message := &kafka.Message{
		Value:          msg,
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            key,
	}
	err := producer.Produce(message, deliveryChan)
	if err != nil {
		return err
	}
	return nil
}

func DeliveryReport(deliveryChan chan kafka.Event) {
	for e := range deliveryChan {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Println("Erro ao enviar")
			} else {
				fmt.Println("Mensagem enviada:", ev.TopicPartition)
				// anotar no banco de dados que a mensagem foi processado.
				// ex: confirma que uma transferencia bancaria ocorreu.
			}
		}
	}
}
