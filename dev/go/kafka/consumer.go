package main

import (
	"context"
	"log"
	"strings"

	"github.com/Shopify/sarama"
)

// Consumer structure
type Consumer struct {
	ready chan bool
}

// Setup method
func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(consumer.ready)
	return nil
}

// Cleanup method
func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim method
func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		log.Printf("CONSUMER : Message claimed: topic=%s, timestamp=%v,	offset=%b, key=%s, value = %s", message.Topic, message.Timestamp, message.Offset, string(message.Key), string(message.Value))
		session.MarkMessage(message, "")
	}
	return nil
}

// ConsumerGroup func
func ConsumerGroup(cli sarama.Client, group string, topics string) {
	saramaConsumerGroup, erreur := sarama.NewConsumerGroupFromClient(group, cli)
	if erreur != nil {
		log.Panicf("Fail to create new Sarama consumer group: %v", erreur)
	}
	defer func() {
		_ = saramaConsumerGroup.Close()
	}()

	// Start consuming messages
	consumer := Consumer{
		ready: make(chan bool),
	}
	saramaContext := context.Background()
	for {
		erreur := saramaConsumerGroup.Consume(saramaContext, strings.Split(topics, ","), &consumer)
		if erreur != nil {
			log.Panicf("Fail to consume: %v", erreur)
		}
	}
}
