package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/Shopify/sarama"
)

// SaramaClient structure
type SaramaClient struct {
	SaramaConfig sarama.Config
	Brokers      []string
}

// Admin method
func (cli *SaramaClient) Admin() sarama.ClusterAdmin {
	client, err := sarama.NewClient(cli.Brokers, &cli.SaramaConfig)
	if err != nil {
		log.Fatal(err.Error())
	}
	admin, err := sarama.NewClusterAdminFromClient(client)
	if err != nil {
		log.Fatal(err.Error())
	}
	return admin
}

// ListTopics method
func (cli *SaramaClient) ListTopics(admin sarama.ClusterAdmin) (gin.Context, error) {
	listTopics, errors := admin.ListTopics()
	if errors != nil {
		return nil, errors
	}
	log.Print(listTopics)
	return listTopics, nil
}

// DescribeTopics method
func (cli *SaramaClient) DescribeTopics(admin sarama.ClusterAdmin, topicName []string) error {
	describeTopic, errors := admin.DescribeTopics(topicName)
	if errors != nil {
		return errors
	}
	fmt.Println(describeTopic)
	return nil
}

// CreateTopics method
func (cli *SaramaClient) CreateTopics(admin sarama.ClusterAdmin, topicName string, topicDetail sarama.TopicDetail, valid bool) error {
	err := admin.CreateTopic(topicName, &topicDetail, valid)
	if err != nil {
		return err
	}
	return nil
}
