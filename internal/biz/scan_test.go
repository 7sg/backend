package biz_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"backend/internal/biz"
	"backend/mocks"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestScanRepository_Success(t *testing.T) {
	scanMocks := mocks.ScanRepo{}
	db, smock, err := sqlmock.New()
	smock.ExpectBegin()
	smock.ExpectCommit()
	tx, err := db.Begin()
	gitCloneKafkaMocks := mocks.GitCloneKafkaRepo{}
	repositoryID := uint64(1)
	expected := &biz.Scan{
		ID:           1,
		RepositoryID: repositoryID,
		Status:       "Queued",
	}
	scanMocks.On("CreateScan", mock.Anything, repositoryID).Return(expected, tx, nil)
	gitCloneKafkaMocks.On("PublishGitClone", mock.Anything, &biz.GitCloneEvent{
		ScanID:       expected.ID,
		RepositoryID: repositoryID,
	}).Return(nil)
	actual, err := biz.NewScanUsecase(&scanMocks, &gitCloneKafkaMocks, log.NewStdLogger(os.Stdout)).
		ScanRepository(nil, repositoryID)
	if err != nil {
		t.Errorf("expected err %v, actual err %v", nil, err)
	}
	assert.Equal(t, expected, actual)
	scanMocks.AssertExpectations(t)
	smock.ExpectationsWereMet()
}

func TestScanRepository_Fail(t *testing.T) {
	scanMocks := mocks.ScanRepo{}
	gitCloneKafkaMocks := mocks.GitCloneKafkaRepo{}
	repositoryID := uint64(1)
	createScanErr := fmt.Errorf("error while creating scan")
	scanMocks.On("CreateScan", mock.Anything, repositoryID).Return(nil, nil, createScanErr)

	_, err := biz.NewScanUsecase(&scanMocks, &gitCloneKafkaMocks, log.NewStdLogger(os.Stdout)).
		ScanRepository(nil, repositoryID)
	assert.ErrorIs(t, err, createScanErr)
	scanMocks.AssertExpectations(t)
}

func TestScanRepository_PublishKafkaFail(t *testing.T) {
	scanMocks := mocks.ScanRepo{}
	db, smock, _ := sqlmock.New()
	smock.ExpectBegin()
	smock.ExpectRollback()
	tx, _ := db.Begin()
	gitCloneKafkaMocks := mocks.GitCloneKafkaRepo{}
	repositoryID := uint64(1)
	expected := &biz.Scan{
		ID:           1,
		RepositoryID: repositoryID,
	}
	scanMocks.On("CreateScan", mock.Anything, repositoryID).Return(expected, tx, nil)
	publishKafkaErr := fmt.Errorf("error while publishing git clone")
	gitCloneKafkaMocks.On("PublishGitClone", mock.Anything, &biz.GitCloneEvent{
		ScanID:       expected.ID,
		RepositoryID: repositoryID,
	}).Return(publishKafkaErr)
	_, err := biz.NewScanUsecase(&scanMocks, &gitCloneKafkaMocks, log.NewStdLogger(os.Stdout)).
		ScanRepository(nil, repositoryID)
	assert.ErrorIs(t, err, publishKafkaErr)
	scanMocks.AssertExpectations(t)
	smock.ExpectationsWereMet()
}

func TestScanRepository_TransactionCommitFail(t *testing.T) {
	scanMocks := mocks.ScanRepo{}
	db, smock, _ := sqlmock.New()
	smock.ExpectBegin()
	txnCommitErr := fmt.Errorf("error while commiting transaction")
	smock.ExpectCommit().WillReturnError(txnCommitErr)
	tx, _ := db.Begin()
	gitCloneKafkaMocks := mocks.GitCloneKafkaRepo{}
	repositoryID := uint64(1)
	expected := &biz.Scan{
		ID:           1,
		RepositoryID: repositoryID,
	}
	scanMocks.On("CreateScan", mock.Anything, repositoryID).Return(expected, tx, nil)
	gitCloneKafkaMocks.On("PublishGitClone", mock.Anything, &biz.GitCloneEvent{
		ScanID:       expected.ID,
		RepositoryID: repositoryID,
	}).Return(nil)
	_, err := biz.NewScanUsecase(&scanMocks, &gitCloneKafkaMocks, log.NewStdLogger(os.Stdout)).
		ScanRepository(nil, repositoryID)
	assert.ErrorIs(t, err, txnCommitErr)
	scanMocks.AssertExpectations(t)
	smock.ExpectationsWereMet()
}

func TestGetScan_Success(t *testing.T) {
	currentTime := time.Now()
	startTime := currentTime.Add(1 * time.Minute)
	endTime := currentTime.Add(2 * time.Minute)
	scanMocks := mocks.ScanRepo{}
	scanID := uint64(1)
	expected := &biz.Scan{
		ID:           1,
		RepositoryID: 1,
		Status:       "Success",
		EnqueuedTime: &currentTime,
		StartTime:    &startTime,
		EndTime:      &endTime,
		Findings: []*biz.Finding{
			{
				Type:   "type",
				RuleID: "ruleID",
				Location: biz.Location{
					Path: "path",
					Positions: biz.Positions{
						Begin: biz.Begin{
							Line: 1,
						},
					},
				},
				Metadata: biz.Metadata{
					Severity:    "severity",
					Description: "description",
				},
			},
		},
	}
	scanMocks.On("GetScan", mock.Anything, scanID).Return(expected, nil)
	actual, err := biz.NewScanUsecase(&scanMocks, nil, log.NewStdLogger(os.Stdout)).
		GetScan(nil, scanID)
	if err != nil {
		t.Errorf("expected err %v, actual err %v", nil, err)
	}
	assert.Equal(t, expected, actual)
	scanMocks.AssertExpectations(t)
}

func TestGetScan_Fail(t *testing.T) {
	scanMocks := mocks.ScanRepo{}
	scanID := uint64(1)
	getScanErr := fmt.Errorf("error while getting scan")
	scanMocks.On("GetScan", mock.Anything, scanID).Return(nil, getScanErr)
	_, err := biz.NewScanUsecase(&scanMocks, nil, log.NewStdLogger(os.Stdout)).GetScan(nil, scanID)
	assert.ErrorIs(t, err, getScanErr)
	scanMocks.AssertExpectations(t)
}
