package main

import (
	"context"
	"flag"
	"log"
	"os"
	"strings"

	"github.com/Shopify/sarama"
)

// Sarama configuration options
var (
	brokers  = ""
	version  = ""
	group    = ""
	topics   = ""
	assignor = ""
	oldest   = true
	verbose  = false
	tls      = true
	sasl     = true
	user     = ""
	password = ""
)

// Consumer type
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

func init() {
	flag.StringVar(&brokers, "brokers", "127.0.0.1:9092", "Kafka bootstrap brokers to connect to, as a comma separated list")
	flag.StringVar(&group, "group", "", "Kafka consumer group definition")
	flag.StringVar(&version, "version", "2.1.1", "Kafka cluster version")
	flag.StringVar(&topics, "topics", "", "Kafka topics to be consumed, as a comma seperated list")
	flag.StringVar(&assignor, "assignor", "range", "Consumer group partition assignment strategy (range, roundrobin, sticky)")
	flag.BoolVar(&oldest, "oldest", true, "Kafka consumer consume initial offset from oldest")
	flag.BoolVar(&verbose, "verbose", false, "Sarama logging")
	flag.BoolVar(&sasl, "sasl", false, "SASL enabled")
	flag.StringVar(&user, "user", "", "SASL user")
	flag.StringVar(&password, "password", "", "SASL password")
	flag.BoolVar(&tls, "tls", false, "TLS enabled")

	flag.Parse()

	if len(brokers) == 0 {
		panic("no Kafka bootstrap brokers defined, please set the -brokers flag")
	}

	if len(topics) == 0 {
		panic("no topics given to be consumed, please set the -topics flag")
	}

	if len(group) == 0 {
		panic("no Kafka consumer group defined, please set the -group flag")
	}
	if len(user) == 0 && sasl == true {
		panic("No user provided while using SASL authentication")
	}
	if len(password) == 0 && sasl == true {
		panic("No password provided while using SASL authentication")
	}
}

func main() {
	log.Println("Starting a new Sarama consumer")

	if verbose {
		sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)
	}

	version, erreur := sarama.ParseKafkaVersion(version)
	if erreur != nil {
		log.Panicf("Error parsing Kafka version: %v", erreur)
	}

	/**
	 * Construct a new Sarama configuration.
	 * The Kafka cluster version has to be defined before the consumer/producer is initialized.
	 */
	config := sarama.NewConfig()
	config.Version = version

	// SASL
	if sasl == true {
		config.Net.SASL.Enable = sasl
		config.Net.SASL.User = user
		config.Net.SASL.Password = password
	}
	if tls == true {
		config.Net.TLS.Enable = tls
	}
	switch assignor {
	case "sticky":
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategySticky
	case "roundrobin":
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	case "range":
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	default:
		log.Panicf("Unrecognized consumer group partition assignor: %s", assignor)
	}

	if oldest {
		config.Consumer.Offsets.Initial = sarama.OffsetOldest
	}

	consumer := Consumer{
		ready: make(chan bool),
	}

	saramaClient, erreur := sarama.NewClient(strings.Split(brokers, ","), config)
	if erreur != nil {
		log.Panicf("Fail to create a new Sarama client: %v", erreur)
	}
	defer func() {
		_ = saramaClient.Close()
	}()

	// Connection to a ConsumerGroup
	saramaConsumerGroup, erreur := sarama.NewConsumerGroupFromClient(group, saramaClient)
	if erreur != nil {
		log.Panicf("Fail to create new Sarama consumer group: %v", erreur)
	}
	defer func() {
		_ = saramaConsumerGroup.Close()
	}()

	// Start consuming messages
	saramaContext := context.Background()
	for {
		erreur := saramaConsumerGroup.Consume(saramaContext, strings.Split(topics, ","), &consumer)
		if erreur != nil {
			log.Panicf("Fail to consume: %v", erreur)
		}
	}

}
