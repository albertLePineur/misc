package main

import (
	"flag"
	"log"
	"strings"

	"github.com/Shopify/sarama"
)

// Sarama configuration options
var (
	brokers    = flag.String("brokers", "localhost:9092", "Comma separated list of brokers (host1:port,hostN:port)")
	version    = flag.String("version", "2.4.0", "Kafka Version (2.1.0 for instance)")
	group      = flag.String("group", "group", "group")
	topics     = flag.String("topics", "monTopic", "Comma separated list of topics")
	tls        = flag.Bool("tls", false, "Activate TLS support")
	sasl       = flag.Bool("sasl", false, "Activate SASL support (requires user/password)")
	saslmethod = flag.String("saslmethod", "", "SASL method to use (PLAIN, SCRAM-SHA-256, SCRAM-SHA-512)")
	user       = flag.String("user", "", "SASL User")
	password   = flag.String("password", "", "SASL Password")
)

func main() {

	flag.Parse()

	cfg := sarama.NewConfig()
	var errors error
	cfg.Version, errors = sarama.ParseKafkaVersion(*version)
	if errors != nil {
		log.Panic(errors.Error())
	}

	// TLS parameters
	cfg.Net.TLS.Enable = *tls

	// SASL parameters
	if *sasl == true {
		switch *saslmethod {
		case "PLAIN":
			cfg.Net.SASL.Enable = *sasl
		case "SCRAM-SHA-256":
			cfg.Net.SASL.Enable = *sasl
			cfg.Net.SASL.Handshake = true
			cfg.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient { return &XDGSCRAMClient{HashGeneratorFcn: SHA256} }
			cfg.Net.SASL.Mechanism = sarama.SASLMechanism(sarama.SASLTypeSCRAMSHA256)
		case "SCRAM-SHA-512":
			cfg.Net.SASL.Enable = *sasl
			cfg.Net.SASL.Handshake = true
			cfg.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient { return &XDGSCRAMClient{HashGeneratorFcn: SHA512} }
			cfg.Net.SASL.Mechanism = sarama.SASLMechanism(sarama.SASLTypeSCRAMSHA512)
		default:
			log.Panic("the saslmethod parameter is missing or not instanciated (PLAIN|SCRAM-SHA-256|SCRAM-SHA-512)")
		}
		if *user == "" {
			log.Panic("No user provided (required for SASL)")
		} else {
			cfg.Net.SASL.User = *user
		}
		if *password == "" {
			log.Panic("No password provided (required for SASL)")
		} else {
			cfg.Net.SASL.Password = *password
		}
	}

	saramaClient, errors := sarama.NewClient(strings.Split(*brokers, ","), cfg)
	if errors != nil {
		log.Fatal(errors.Error())
	}

	KafkaAdmin(saramaClient)
	ConsumerGroup(saramaClient, *group, *topics)

	defer func() { _ = saramaClient.Close() }()
}
