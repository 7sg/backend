package data

import (
	"backend-GuardRails/internal/biz"
	"backend-GuardRails/internal/conf"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/segmentio/kafka-go"
)

type fileContentKafkaRepo struct {
	producer *kafka.Writer
	consumer *kafka.Reader
	log      *log.Helper
}

func NewFileContentKafkaRepo(c *conf.Data, logger log.Logger) biz.FileContentKafkaRepo {
	return &fileContentKafkaRepo{
		producer: newKafkaFileContentProducer(c),
		consumer: newKafkaFileContentConsumer(c),
		log:      log.NewHelper(logger),
	}
}

func (f *fileContentKafkaRepo) PublishFileContent(ctx context.Context, gitClone *biz.FileContent) error {
	return nil
}

func (f *fileContentKafkaRepo) GetMessage(ctx context.Context) (*kafka.Message, error) {
	return nil, nil
}

func (f *fileContentKafkaRepo) CommitMessage(ctx context.Context, message *kafka.Message) error {
	return nil
}
