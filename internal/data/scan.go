package data

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"backend-GuardRails/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

const (
	Queued     = "Queued"
	InProgress = "InProgress"
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

func (s *scanRepo) GetScan(ctx context.Context, scanID uint64) (*biz.Scan, error) {
	scan := &biz.Scan{}
	var findings string
	var startTime, endTime sql.NullTime
	err := s.data.db.QueryRow("SELECT id, repository_id, enqueued_time, start_time, end_time, status, findings "+
		"FROM scan_repository WHERE id = $1", scanID).
		Scan(&scan.ID, &scan.RepositoryID, &scan.EnqueuedTime, &startTime, &endTime, &scan.Status, &findings)
	if err != nil {
		s.log.Errorf("error while getting scan: %v", err)

		return nil, err
	}
	err = json.Unmarshal([]byte(findings), &scan.Findings)
	if err != nil {
		s.log.Errorf("error while unmarshalling findings: %v", err)

		return nil, err
	}

	if startTime.Valid {
		scan.StartTime = &startTime.Time
	}
	if endTime.Valid {
		scan.EndTime = &endTime.Time
	}

	return scan, nil
}

func (s *scanRepo) CreateScan(ctx context.Context, repoID uint64) (*biz.Scan, *sql.Tx, error) {
	tx, err := s.data.db.BeginTx(ctx, nil)
	if err != nil {
		s.log.Errorf("error while creating transaction: %v", err)

		return nil, nil, err
	}
	scan := &biz.Scan{RepositoryID: repoID}
	err = tx.QueryRow("INSERT INTO scan_repository (repository_id, enqueued_time, status) VALUES ($1, $2, $3) "+
		"RETURNING id, enqueued_time, status", repoID, time.Now(), Queued).
		Scan(&scan.ID, &scan.EnqueuedTime, &scan.Status)
	if err != nil {
		s.log.Errorf("error while inserting scan: %v", err)
		_ = tx.Rollback()

		return nil, nil, err
	}

	return scan, tx, nil
}

func (s *scanRepo) UpdateScanStartTimeAndTotalFiles(
	ctx context.Context,
	scanID uint64,
	startTime time.Time,
	totalFiles uint32,
) error {
	_, err := s.data.db.ExecContext(
		ctx,
		"UPDATE scan_repository SET start_time = $1, status = $2, total_files = $3 WHERE id = $4",
		startTime,
		InProgress,
		totalFiles,
		scanID,
	)
	if err != nil {
		s.log.Errorf("error while updating scan start time: %v", err)

		return err
	}

	return nil
}

func (s *scanRepo) UpdateScanStartAndEndTimeWithSuccess(
	ctx context.Context,
	scanID uint64,
	startTime time.Time,
	endTime time.Time,
) error {
	_, err := s.data.db.ExecContext(
		ctx,
		"UPDATE scan_repository SET start_time = $1, end_time = $2, status = $3 WHERE id = $4",
		startTime,
		endTime,
		Success,
		scanID,
	)
	if err != nil {
		s.log.Errorf("error while updating scan start and end time: %v", err)
	}

	return err
}

func (s *scanRepo) UpdateScanStartAndEndTimeWithFailure(
	ctx context.Context,
	scanID uint64,
	startTime time.Time,
	endTime time.Time,
) error {
	_, err := s.data.db.ExecContext(
		ctx,
		"UPDATE scan_repository SET start_time = $1, end_time = $2, status = $3 WHERE id = $4",
		startTime,
		endTime,
		Failure,
		scanID,
	)
	if err != nil {
		s.log.Errorf("error while updating scan start and end time: %v", err)
	}

	return err
}

func (s *scanRepo) UpdateScanFindings(
	ctx context.Context,
	scanID uint64,
	findings string,
) (*biz.ScanUpdate, error) {
	scanUpdate := &biz.ScanUpdate{}
	err := s.data.db.QueryRow("UPDATE scan_repository SET findings = findings || $1::jsonb,  scanned_files = scanned_files+1 "+
		"WHERE id = $2 RETURNING total_files, scanned_files", findings, scanID).
		Scan(&scanUpdate.TotalFiles, &scanUpdate.ScannedFiles)
	if err != nil {
		s.log.Errorf("error while updating scan findings: %v", err)

		return nil, err
	}

	return scanUpdate, nil
}

func (s *scanRepo) UpdateScanWithSuccess(
	ctx context.Context,
	scanID uint64,
	endTime time.Time,
) error {
	_, err := s.data.db.ExecContext(
		ctx,
		"UPDATE scan_repository SET end_time = $1, status = $2 WHERE id = $3",
		endTime,
		Success,
		scanID,
	)
	if err != nil {
		s.log.Errorf("error while updating scan with success: %v", err)
	}

	return err
}
