package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Repository struct {
	ID   uint64
	Name string
	Link string
}

type RepositoryUsecase struct {
	repo RepositoryRepo
	log  *log.Helper
}

// NewRepositoryUsecase new a Repository usecase.
func NewRepositoryUsecase(repo RepositoryRepo, logger log.Logger) *RepositoryUsecase {
	return &RepositoryUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (r *RepositoryUsecase) CreateRepository(
	ctx context.Context,
	repo *Repository,
) (*Repository, error) {
	repository, err := r.repo.CreateRepository(ctx, repo)
	if err != nil {
		r.log.Errorf("error while CreateRepository: %v", err)

		return nil, err
	}

	return repository, nil
}

func (r *RepositoryUsecase) GetRepository(ctx context.Context, id uint64) (*Repository, error) {
	repository, err := r.repo.GetRepository(ctx, id)
	if err != nil {
		r.log.Errorf("error while GetRepository: %v", err)

		return nil, err
	}

	return repository, nil
}
