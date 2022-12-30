package data

import (
	"backend-GuardRails/internal/conf"
	"github.com/segmentio/kafka-go"
)

func newKafkaGitCloneConsumer(c *conf.Data) *kafka.Reader {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{c.Kafka.Brokers},
		GroupID:     c.Kafka.GitCloneConsumerGroupId,
		Topic:       c.Kafka.GitCloneTopic,
		StartOffset: kafka.FirstOffset,
	})
	return r
}

func newKafkaFileContentConsumer(c *conf.Data) *kafka.Reader {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{c.Kafka.Brokers},
		GroupID:     c.Kafka.FileContentConsumerGroupId,
		Topic:       c.Kafka.FileContentTopic,
		StartOffset: kafka.FirstOffset,
	})
	return r
}

func newKafkaGitCloneProducer(c *conf.Data) *kafka.Writer {
	w := &kafka.Writer{
		Addr:                   kafka.TCP(c.Kafka.Brokers),
		Topic:                  c.Kafka.GitCloneTopic,
		AllowAutoTopicCreation: true,
	}
	return w
}

func newKafkaFileContentProducer(c *conf.Data) *kafka.Writer {
	w := &kafka.Writer{
		Addr:                   kafka.TCP(c.Kafka.Brokers),
		Topic:                  c.Kafka.FileContentTopic,
		AllowAutoTopicCreation: true,
	}
	return w
}
