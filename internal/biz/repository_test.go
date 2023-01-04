package biz_test

import (
	"fmt"
	"os"
	"testing"

	"backend/internal/biz"
	"backend/mocks"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// write unit test for internal/biz/repository.go
func TestCreateRepository_Success(t *testing.T) {
	mocks := mocks.RepositoryRepo{}
	expected := &biz.Repository{
		ID:   1,
		Name: "my repo",
		Link: "github.com/my/repo",
	}
	mocks.On("CreateRepository", mock.Anything, mock.Anything).Return(expected, nil)
	actual, err := biz.NewRepositoryUsecase(&mocks, log.NewStdLogger(os.Stdout)).
		CreateRepository(nil, nil)
	if err != nil {
		t.Errorf("expected err %v, actual err %v", nil, err)
	}
	assert.Equal(t, expected, actual)
	mocks.AssertExpectations(t)
}

func TestCreateRepository_Failure(t *testing.T) {
	mocks := mocks.RepositoryRepo{}
	expectedErr := fmt.Errorf("error while CreateRepository")
	mocks.On("CreateRepository", mock.Anything, mock.Anything).Return(nil, expectedErr)
	actual, err := biz.NewRepositoryUsecase(&mocks, log.NewStdLogger(os.Stdout)).
		CreateRepository(nil, nil)
	assert.EqualError(t, err, expectedErr.Error())
	assert.Nil(t, actual)
	mocks.AssertExpectations(t)
}

func TestGetRepository_Success(t *testing.T) {
	mocks := mocks.RepositoryRepo{}
	ID := uint64(1)
	expected := &biz.Repository{
		ID:   ID,
		Name: "my repo",
		Link: "github.com/my/repo",
	}
	mocks.On("GetRepository", mock.Anything, ID).Return(expected, nil)
	actual, err := biz.NewRepositoryUsecase(&mocks, log.NewStdLogger(os.Stdout)).
		GetRepository(nil, ID)
	if err != nil {
		t.Errorf("expected err %v, actual err %v", nil, err)
	}
	assert.Equal(t, expected, actual)
	mocks.AssertExpectations(t)
}

func TestGetRepository_Failure(t *testing.T) {
	mocks := mocks.RepositoryRepo{}
	expectedErr := fmt.Errorf("error while GetRepository")
	mocks.On("GetRepository", mock.Anything, mock.Anything).Return(nil, expectedErr)
	actual, err := biz.NewRepositoryUsecase(&mocks, log.NewStdLogger(os.Stdout)).
		GetRepository(nil, 1)
	assert.EqualError(t, err, expectedErr.Error())
	assert.Nil(t, actual)
	mocks.AssertExpectations(t)
}
