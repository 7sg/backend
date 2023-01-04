package data

import (
	"context"

	"backend-GuardRails/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type repositoryRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewRepositoryRepo(data *Data, logger log.Logger) biz.RepositoryRepo {
	return &repositoryRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r repositoryRepo) CreateRepository(
	ctx context.Context,
	repo *biz.Repository,
) (*biz.Repository, error) {
	err := r.data.db.QueryRow("INSERT INTO repository (name, link) VALUES ($1, $2) RETURNING id", repo.Name, repo.Link).
		Scan(&repo.ID)
	if err != nil {
		r.log.Errorf("error while inserting repository: %v", err)

		return nil, err
	}

	return repo, nil
}

func (r repositoryRepo) GetRepository(ctx context.Context, id uint64) (*biz.Repository, error) {
	repo := &biz.Repository{}
	err := r.data.db.QueryRow("SELECT id, name, link FROM repository WHERE id = $1", id).
		Scan(&repo.ID, &repo.Name, &repo.Link)
	if err != nil {
		r.log.Errorf("error while getting repository: %v", err)

		return nil, err
	}

	return repo, nil
}
