package biz

import (
	"context"
	"database/sql"
	"github.com/segmentio/kafka-go"
)

type RepositoryRepo interface {
	CreateRepository(ctx context.Context, repo *Repository) (*Repository, error)
}

type ScanRepo interface {
	CreateScan(ctx context.Context, repoId uint64) (*Scan, *sql.Tx, error)
}

type GitCloneKafkaRepo interface {
	PublishGitClone(ctx context.Context, gitClone *GitCloneEvent) error
	GetMessage(ctx context.Context) (*kafka.Message, error)
	CommitMessage(ctx context.Context, message *kafka.Message) error
}

type FileContentKafkaRepo interface {
	PublishFileContent(ctx context.Context, gitClone *FileContent) error
	GetMessage(ctx context.Context) (*kafka.Message, error)
	CommitMessage(ctx context.Context, message *kafka.Message) error
}
