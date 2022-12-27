package service

import (
	v1 "backend-GuardRails/api/repository/v1"
	"backend-GuardRails/internal/biz"
	"context"
	"fmt"
)

type RepositoryService struct {
	v1.UnimplementedRepositoryServer

	uc *biz.RepositoryUsecase
}

// NewRepositoryService new a repository service.
func NewRepositoryService(uc *biz.RepositoryUsecase) *RepositoryService {
	return &RepositoryService{uc: uc}
}

func (r *RepositoryService) CreateRepository(ctx context.Context, createRepoRequest *v1.CreateRepositoryRequest) (*v1.CreateRepositoryResponse, error) {
	if len(createRepoRequest.Name) == 0 || len(createRepoRequest.Link) == 0 {
		return nil, fmt.Errorf("name or link is empty")
	}
	repository, err := r.uc.CreateRepository(ctx, &biz.Repository{
		Name: createRepoRequest.Name,
		Link: createRepoRequest.Link,
	})
	if err != nil {
		return &v1.CreateRepositoryResponse{Message: "failed to create repository"}, nil
	}
	return &v1.CreateRepositoryResponse{Id: repository.Id, Message: "Repository successfully created"}, nil
}
