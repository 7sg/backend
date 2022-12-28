package biz

import (
	"context"
	"database/sql"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type Scan struct {
	Id           uint64
	RepositoryId uint64
	Status       string
	EnqueuedTime time.Time
}

type ScanRepo interface {
	CreateScan(ctx context.Context, repoId uint64) (*Scan, *sql.Tx, error)
}

type ScanUsecase struct {
	repo ScanRepo
	log  *log.Helper
}

// NewScanUsecase new a Scan usecase.
func NewScanUsecase(repo ScanRepo, logger log.Logger) *ScanUsecase {
	return &ScanUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (s *ScanUsecase) ScanRepository(ctx context.Context, repoId uint64) (*Scan, error) {
	scan, tx, err := s.repo.CreateScan(ctx, repoId)
	if err != nil {
		s.log.Errorf("error while creating scan: %v", err)
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		s.log.Errorf("error while committing transaction: %v", err)
		return nil, err
	}
	return scan, nil
}
