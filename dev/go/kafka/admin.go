package main

import (
	"log"

	"github.com/Shopify/sarama"
)

// KafkaAdmin function
func KafkaAdmin(cli sarama.Client) {
	clusterAdmin, errors := sarama.NewClusterAdminFromClient(cli)
	if errors != nil {
		log.Fatal(errors.Error())
	}

	listTopics, errors := clusterAdmin.ListTopics()
	if errors != nil {
		log.Fatal(errors.Error())
	}
	log.Print(listTopics)
	if val, ok := listTopics[*topics]; ok {
		log.Print("Le topic suivant est bien present : ", val)
	} else {
		log.Print("Le topic suivant n'existe pas et va etre cree : ", val)

		errors = clusterAdmin.CreateTopic(*topics, &sarama.TopicDetail{NumPartitions: 5, ReplicationFactor: 1}, false)
		if errors != nil {
			log.Fatal("MAYDAY MAYDAY : ", errors)
		}
	}
}
