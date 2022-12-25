package service

import (
	v1 "backend-GuardRails/api/repository/v1"
	"backend-GuardRails/internal/biz"
	"context"
)

type RepositoryService struct {
	v1.UnimplementedRepositoryServer

	uc *biz.RepositoryUsecase
}

// NewRepositoryService new a repository service.
func NewRepositoryService(uc *biz.RepositoryUsecase) *RepositoryService {
	return &RepositoryService{uc: uc}
}

func (r *RepositoryService) CreateRepository(context.Context, *v1.CreateRepositoryRequest) (*v1.CreateRepositoryResponse, error) {
	return nil, nil
}
