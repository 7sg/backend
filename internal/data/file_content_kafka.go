package data

import (
	"context"
	"encoding/json"

	"backend-GuardRails/internal/biz"
	"backend-GuardRails/internal/conf"
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

func (f *fileContentKafkaRepo) PublishFileContent(
	ctx context.Context,
	fileContents []*biz.FileContent,
) error {
	f.log.Infof("publishing file content, content size: %d", len(fileContents))
	messages := make([]kafka.Message, 0)
	for _, fileContent := range fileContents {
		bytes, err := json.Marshal(fileContent)
		if err != nil {
			f.log.Errorf("error while marshalling file content: %v", err)

			return err
		}
		messages = append(messages, kafka.Message{Value: bytes})
	}

	return f.producer.WriteMessages(ctx, messages...)
}

func (f *fileContentKafkaRepo) GetMessage(ctx context.Context) (*kafka.Message, error) {
	message, err := f.consumer.FetchMessage(ctx)
	if err != nil {
		f.log.Errorf("error while fetching file content message: %v", err)

		return nil, err
	}

	return &message, nil
}

func (f *fileContentKafkaRepo) CommitMessage(ctx context.Context, message *kafka.Message) error {
	err := f.consumer.CommitMessages(ctx, *message)
	if err != nil {
		f.log.Errorf("error while committing file content message: %v", err)

		return err
	}

	return nil
}
