package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type Scan struct {
	ID           uint64
	RepositoryID uint64
	Status       string
	Findings     []*Finding
	EnqueuedTime *time.Time
	StartTime    *time.Time
	EndTime      *time.Time
}

type Finding struct {
	Type     string
	RuleID   string
	Location Location
	Metadata Metadata
}

type Location struct {
	Path      string
	Positions Positions
}

type Positions struct {
	Begin Begin
}

type Begin struct {
	Line uint32
}

type Metadata struct {
	Description string
	Severity    string
}

type ScanUpdate struct {
	TotalFiles   uint32
	ScannedFiles uint32
}

type GitCloneEvent struct {
	ScanID       uint64
	RepositoryID uint64
}

type ScanUsecase struct {
	repo      ScanRepo
	KafkaRepo GitCloneKafkaRepo
	log       *log.Helper
}

// NewScanUsecase new a Scan usecase.
func NewScanUsecase(repo ScanRepo, kafkaRepo GitCloneKafkaRepo, logger log.Logger) *ScanUsecase {
	return &ScanUsecase{repo: repo, KafkaRepo: kafkaRepo, log: log.NewHelper(logger)}
}

func (s *ScanUsecase) ScanRepository(ctx context.Context, repoID uint64) (*Scan, error) {
	scan, tx, err := s.repo.CreateScan(ctx, repoID)
	if err != nil {
		s.log.Errorf("error while creating scan: %v", err)

		return nil, err
	}
	err = s.KafkaRepo.PublishGitClone(ctx, &GitCloneEvent{ScanID: scan.ID, RepositoryID: repoID})
	if err != nil {
		_ = tx.Rollback()
		s.log.Errorf("error while publishing git clone: %v", err)

		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		s.log.Errorf("error while committing transaction: %v", err)

		return nil, err
	}

	return scan, nil
}

func (s *ScanUsecase) GetScan(ctx context.Context, scanID uint64) (*Scan, error) {
	scan, err := s.repo.GetScan(ctx, scanID)
	if err != nil {
		s.log.Errorf("error while getting scan: %v", err)

		return nil, err
	}
	s.log.Infof("scan: %+v", scan)

	return scan, nil
}
