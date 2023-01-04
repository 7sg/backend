package service

import (
	"context"
	"fmt"

	v1 "backend-GuardRails/api/repository/v1"
	"backend-GuardRails/internal/biz"
)

type RepositoryService struct {
	v1.UnimplementedRepositoryServer

	uc *biz.RepositoryUsecase
}

// NewRepositoryService new a repository service.
func NewRepositoryService(uc *biz.RepositoryUsecase) *RepositoryService {
	return &RepositoryService{uc: uc}
}

func (r *RepositoryService) CreateRepository(
	ctx context.Context,
	createRepoRequest *v1.CreateRepositoryRequest,
) (*v1.CreateRepositoryResponse, error) {
	if len(createRepoRequest.Name) == 0 || len(createRepoRequest.Link) == 0 {
		return nil, fmt.Errorf("name or link is empty")
	}
	repository, err := r.uc.CreateRepository(ctx, &biz.Repository{
		Name: createRepoRequest.Name,
		Link: createRepoRequest.Link,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create repository, try again later")
	}

	return &v1.CreateRepositoryResponse{
		Id:      repository.ID,
		Message: "Repository successfully created",
	}, nil
}

func (r *RepositoryService) GetRepository(
	ctx context.Context,
	getRepositoryRequest *v1.GetRepositoryRequest,
) (*v1.GetRepositoryResponse, error) {
	repository, err := r.uc.GetRepository(ctx, getRepositoryRequest.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to get repository, try again later")
	}

	return &v1.GetRepositoryResponse{
		Id:   repository.ID,
		Name: repository.Name,
		Link: repository.Link,
	}, nil
}
