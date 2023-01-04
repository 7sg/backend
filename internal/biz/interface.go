package biz

import (
	"context"
	"database/sql"
	"time"

	"github.com/segmentio/kafka-go"
)

type RepositoryRepo interface {
	CreateRepository(ctx context.Context, repo *Repository) (*Repository, error)
	GetRepository(ctx context.Context, id uint64) (*Repository, error)
}

type ScanRepo interface {
	CreateScan(ctx context.Context, repoID uint64) (*Scan, *sql.Tx, error)
	GetScan(ctx context.Context, scanID uint64) (*Scan, error)
	UpdateScanStartTimeAndTotalFiles(
		ctx context.Context,
		scanID uint64,
		startTime time.Time,
		totalFiles uint32,
	) error
	UpdateScanStartAndEndTimeWithSuccess(
		ctx context.Context,
		scanID uint64,
		startTime time.Time,
		endTime time.Time,
	) error
	UpdateScanStartAndEndTimeWithFailure(
		ctx context.Context,
		scanID uint64,
		startTime time.Time,
		endTime time.Time,
	) error
	UpdateScanFindings(ctx context.Context, scanID uint64, findings string) (*ScanUpdate, error)
	UpdateScanWithSuccess(ctx context.Context, scanID uint64, endTime time.Time) error
}

type GitCloneKafkaRepo interface {
	PublishGitClone(ctx context.Context, gitClone *GitCloneEvent) error
	GetMessage(ctx context.Context) (*kafka.Message, error)
	CommitMessage(ctx context.Context, message *kafka.Message) error
}

type FileContentKafkaRepo interface {
	PublishFileContent(ctx context.Context, gitClone []*FileContent) error
	GetMessage(ctx context.Context) (*kafka.Message, error)
	CommitMessage(ctx context.Context, message *kafka.Message) error
}
