package data

import (
	"context"
	"encoding/json"

	"backend-GuardRails/internal/biz"
	"backend-GuardRails/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/segmentio/kafka-go"
)

type GitCloneKafkaRepo struct {
	producer *kafka.Writer
	consumer *kafka.Reader
	log      *log.Helper
}

func NewGitCloneKafkaRepo(c *conf.Data, logger log.Logger) biz.GitCloneKafkaRepo {
	return &GitCloneKafkaRepo{
		producer: newKafkaGitCloneProducer(c),
		consumer: newKafkaGitCloneConsumer(c),
		log:      log.NewHelper(logger),
	}
}

func (g *GitCloneKafkaRepo) PublishGitClone(
	ctx context.Context,
	gitClone *biz.GitCloneEvent,
) error {
	bytes, err := json.Marshal(gitClone)
	if err != nil {
		g.log.Errorf("error while marshalling git clone event: %v", err)

		return err
	}
	err = g.producer.WriteMessages(ctx, kafka.Message{
		Value: bytes,
	})
	if err != nil {
		g.log.Errorf("error while publishing git clone: %v", err)

		return err
	}

	return nil
}

func (g *GitCloneKafkaRepo) GetMessage(ctx context.Context) (*kafka.Message, error) {
	message, err := g.consumer.FetchMessage(ctx)
	if err != nil {
		g.log.Errorf("error while fetching git clone message: %v", err)

		return nil, err
	}

	return &message, nil
}

func (g *GitCloneKafkaRepo) CommitMessage(ctx context.Context, message *kafka.Message) error {
	err := g.consumer.CommitMessages(ctx, *message)
	if err != nil {
		g.log.Errorf("error while committing git clone message: %v", err)

		return err
	}

	return nil
}
