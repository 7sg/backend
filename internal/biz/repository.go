package biz

import "github.com/go-kratos/kratos/v2/log"

type RepositoryRepo interface {
}

type RepositoryUsecase struct {
	repo RepositoryRepo
	log  *log.Helper
}

// NewRepositoryUsecase new a Repository usecase.
func NewRepositoryUsecase(repo RepositoryRepo, logger log.Logger) *RepositoryUsecase {
	return &RepositoryUsecase{repo: repo, log: log.NewHelper(logger)}
}
