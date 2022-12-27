package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Repository struct {
	Id   string
	Name string
	Link string
}

type RepositoryRepo interface {
	CreateRepository(ctx context.Context, repo *Repository) (*Repository, error)
}

type RepositoryUsecase struct {
	repo RepositoryRepo
	log  *log.Helper
}

// NewRepositoryUsecase new a Repository usecase.
func NewRepositoryUsecase(repo RepositoryRepo, logger log.Logger) *RepositoryUsecase {
	return &RepositoryUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (r *RepositoryUsecase) CreateRepository(ctx context.Context, repo *Repository) (*Repository, error) {
	repository, err := r.repo.CreateRepository(ctx, repo)
	if err != nil {
		r.log.Errorf("error while CreateRepository: %v", err)
		return nil, err
	}
	return repository, nil
}
