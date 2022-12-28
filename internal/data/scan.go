package data

import (
	"backend-GuardRails/internal/biz"
	"context"
	"database/sql"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

const (
	Queued     = "Queued"
	InProgress = "In Progress"
	Success    = "Success"
	Failure    = "Failure"
)

type scanRepo struct {
	data *Data
	log  *log.Helper
}

// NewScanRepo .
func NewScanRepo(data *Data, logger log.Logger) biz.ScanRepo {
	return &scanRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (s *scanRepo) CreateScan(ctx context.Context, repoId uint64) (*biz.Scan, *sql.Tx, error) {
	tx, err := s.data.db.BeginTx(ctx, nil)
	if err != nil {
		s.log.Errorf("error while creating transaction: %v", err)
		return nil, nil, err
	}
	scan := &biz.Scan{RepositoryId: repoId}
	err = tx.QueryRow("INSERT INTO scan_repository (repository_id, enqueued_time, status) VALUES ($1, $2, $3) RETURNING id, enqueued_time, status", repoId, time.Now(), Queued).
		Scan(&scan.Id, &scan.EnqueuedTime, &scan.Status)
	if err != nil {
		s.log.Errorf("error while inserting scan: %v", err)
		tx.Rollback()
		return nil, nil, err
	}

	return scan, tx, nil
}
